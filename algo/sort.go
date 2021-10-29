package algo

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"plugin"
	"sort"
	"strconv"
	"time"
)

func sortF(input []float64) []float64 {

	ConcurrentQuickSort(&input, 0, len(input)-1)
	return input
}


var (
	pluginPath = flag.String("in", "../plugins/dummy/dummy.so", "plugin path")
	//dataSetSize = flag.Int("size", 10000000, "data set size")
	dataSetSize = 10000000
)

const (
	SortFuncName      = "Sort"
	randomFileNameLen = 3
	precision         = 20
)

type SortFunc = func([]float64) []float64

func LoadPlugin() SortFunc {
	p, err := plugin.Open(*pluginPath)
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup(SortFuncName)
	if err != nil {
		panic(err)
	}
	sortF, ok := f.(SortFunc)
	if !ok {
		fmt.Printf("sort func type: %T\n", f)
		panic("sort function signature wrong")
	}

	return sortF
}



func RunGame() {

	for i := 0; i < 10; i++ {
		data := GenDataSet()
		input := ReadDataSet(data)
		start := time.Now()
		expected := libsort(input)
		fmt.Printf("libsort time used: %v\n", time.Since(start))
		start = time.Now()
		res := sortF(input)
		fmt.Printf("time used: %v\n", time.Since(start))

		ValidateResult(res, expected)
	}

}

func GenDataSet() []string {
	rand.Seed(time.Now().UnixNano())
	if dataSetSize < 0 {
		panic("data set size should be greater than 0")
	}
	data := []string{}
	for i := 0; i < dataSetSize; i++ {
		data = append(data, strconv.FormatFloat(rand.Float64(), 'g', precision, 64))
	}
	return data
}

func ReadDataSet(strs []string) []float64 {
	size := len(strs)
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i], _ = strconv.ParseFloat(strs[i], 64)
	}
	return data
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func GetRandomFileName() string {
	b := make([]rune, randomFileNameLen)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ValidateResult(output, expected []float64) {
	if len(output) != len(expected) {
		fmt.Printf("result length wrong\n")
		return
	}
	for i := 0; i < len(output); i++ {
		if output[i] != expected[i] {
			fmt.Printf("wrong result at index: %v\n", i)
			return
		}
	}
	fmt.Printf("validation passed\n")
}

func libsort(input []float64) []float64 {
	output := make([]float64, len(input))
	copy(output, input)
	sort.Float64s(output)
	return output
}

// ReadFile reads file given full absolute path
func ReadFile(filepath string) string {
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("error reading file %v : %v\n", filepath, err)
		return ""
	}
	return string(bs)
}

// WriteFile writes to file given full absolute path
func WriteFile(filepath, content string) {
	err := ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("error writing file %v: %v\n", filepath, err)
	}
}
