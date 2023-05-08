package stresstest

import (
	"math/rand"
	"testing"
	"time"
)

var inputBytes [][]byte

func randString() []byte {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5400) + 600
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return []byte(string(b))
}

func init() {
	// Generate 1000 inputs in advance and cache them
	inputBytes = make([][]byte, 10000)
	for i := range inputBytes {
		inputBytes[i] = randString()
	}
}

func getInput() []byte {
	return inputBytes[rand.Intn(len(inputBytes))]
}

//func BenchmarkDoZlibCompress(b *testing.B) {
//	var wg sync.WaitGroup
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			DoZlibCompressBufferPool(getInput())
//		}()
//
//	}
//	wg.Wait()
//}
//
//func BenchmarkDoZlibCompressOld(b *testing.B) {
//	var wg sync.WaitGroup
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			DoZlibCompressOld(getInput())
//		}()
//
//	}
//	wg.Wait()
//}
//
//func BenchmarkDoZlibCompressWriter(b *testing.B) {
//	var wg sync.WaitGroup
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			DoZlibCompressWritterPool(getInput())
//		}()
//
//	}
//	wg.Wait()
//}
//
//func BenchmarkDoZlibCompress2Pool(b *testing.B) {
//	var wg sync.WaitGroup
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			DoZlibCompress2Pool(getInput())
//		}()
//
//	}
//	wg.Wait()
//}

func BenchmarkDoZlibCompressP(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			DoZlibCompressBufferPool(getInput())
		}

	})
}

func BenchmarkDoZlibCompressOldP(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			DoZlibCompressOld(getInput())
		}

	})
}

func BenchmarkDoZlibCompressWriterP(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DoZlibCompressWritterPool(getInput())
		}
	})
}

func BenchmarkDoZlibCompress2PoolP(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DoZlibCompress2Pool(getInput())
		}
	})
}
