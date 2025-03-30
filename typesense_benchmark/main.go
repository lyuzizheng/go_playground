package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"sync"
	"time"
)

type SearchResponse struct {
	SearchTime float64 `json:"SearchTime"`
}

type Result struct {
	Duration   time.Duration
	SearchTime float64
	URL        string
}

func makeRequest(url string, headers map[string]string, isLocal bool) (Result, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{}, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	result := Result{
		Duration: time.Since(start),
		URL:      url,
	}

	if isLocal {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Result{}, err
		}

		var searchResp SearchResponse
		if err := json.Unmarshal(body, &searchResp); err != nil {
			return Result{}, err
		}
		result.SearchTime = searchResp.SearchTime
	}

	return result, nil
}

func worker(jobs <-chan string, results chan<- Result, headers map[string]string, isLocal bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		result, err := makeRequest(url, headers, isLocal)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			continue
		}
		results <- result
	}
}

func calculateStats(durations []time.Duration, searchTimes []float64) (avgDur, p95Dur time.Duration, avgSearch, p95Search float64) {
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	var sumDur time.Duration
	for _, d := range durations {
		sumDur += d
	}
	avgDur = sumDur / time.Duration(len(durations))
	p95Index := int(math.Ceil(float64(len(durations))*0.95)) - 1
	p95Dur = durations[p95Index]

	if len(searchTimes) > 0 {
		sort.Float64s(searchTimes)
		var sumSearch float64
		for _, s := range searchTimes {
			sumSearch += s
		}
		avgSearch = sumSearch / float64(len(searchTimes))
		p95Search = searchTimes[p95Index]
	}

	return
}

func runBenchmark(endpoint string, url string, iterations int, numWorkers int, headers map[string]string, isLocal bool) {
	jobs := make(chan string, iterations)
	results := make(chan Result, iterations)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, headers, isLocal, &wg)
	}

	go func() {
		for i := 0; i < iterations; i++ {
			if isLocal {
				food := foods[i%len(foods)]
				jobs <- fmt.Sprintf("http://192.168.100.10:8000/app/v2.81/typesense?query=%s&pageSize=5&currentPage=1&vertical=All&parentId=0&language=en&hexagon=892c105988fffff&city=Erbil&appVersion=3.16", food)
			} else {
				jobs <- url
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var durations []time.Duration
	var searchTimes []float64
	for result := range results {
		durations = append(durations, result.Duration)
		if result.SearchTime > 0 {
			searchTimes = append(searchTimes, result.SearchTime)
		}
	}

	avgDur, p95Dur, avgSearch, p95Search := calculateStats(durations, searchTimes)
	fmt.Printf("\nResults for %s:\n", endpoint)
	fmt.Printf("Total Response Time:\n")
	fmt.Printf("  Average: %v\n", avgDur)
	fmt.Printf("  P95: %v\n", p95Dur)

	if isLocal {
		fmt.Printf("\nTypesense Search Time:\n")
		fmt.Printf("  Average: %.2fms\n", avgSearch)
		fmt.Printf("  P95: %.2fms\n", p95Search)
	}
}

var foods = []string{
	"ChickenBurger", "BeefBurger", "Pizza", "Pasta", "Sushi",
	"Salad", "Sandwich", "HotDog", "Taco", "Kebab",
	"Falafel", "Shawarma", "Steak", "Rice", "Noodles",
}

func main() {
	iterations := 300
	numWorkers := 3

	fmt.Println("Starting search API benchmark...")
	runBenchmark(
		"Local Search API",
		"",
		iterations,
		numWorkers,
		map[string]string{"Connection": "keep-alive"},
		true,
	)

	fmt.Println("\nStarting metrics API benchmark...")
	runBenchmark(
		"Metrics API",
		"https://1qnohjb47fpzycgup-1.a1.typesense.net/metrics.json",
		iterations,
		numWorkers,
		map[string]string{"X-TYPESENSE-API-KEY": "4ybOZTuQuQOFsVPqsaz6KdGAHPt8B1Ka"},
		false,
	)

	fmt.Println("\nStarting health API benchmark...")
	runBenchmark(
		"Health API",
		"https://1qnohjb47fpzycgup-1.a1.typesense.net/health",
		iterations,
		numWorkers,
		map[string]string{},
		false,
	)
}
