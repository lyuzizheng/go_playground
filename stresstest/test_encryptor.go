package stresstest

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

//var pool = &syncx.Pool{
//	New: func() interface{} {
//		return &struct{}
//	},
//	NoGC: true,
//}

// 进行zlib压缩
func DoZlibCompressBufferPool(src []byte) []byte {

	// Get a buffer from the pool
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)

	// Reset the buffer before writing to it
	buf.Reset()

	w := zlib.NewWriter(buf)
	_, error := w.Write(src)
	if error != nil {
		return nil
	}
	error = w.Close()
	if error != nil {
		return nil
	}

	// Create a new slice with the compressed data
	compressedData, _ := ioutil.ReadAll(buf)
	return compressedData
}

// 进行zlib压缩
func DoZlibCompressOld(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, error := w.Write(src)
	if error != nil {
		return nil
	}
	error = w.Close()
	if error != nil {
		return nil
	}

	return in.Bytes()
}

var writerPool = sync.Pool{
	New: func() interface{} {
		return zlib.NewWriter(nil)
	},
}

func DoZlibCompressWritterPool(data []byte) []byte {
	var buf bytes.Buffer
	writer := writerPool.Get().(*zlib.Writer)
	writer.Reset(&buf)
	_, err := writer.Write(data)
	if err != nil {
		return nil
	}
	err = writer.Close()
	if err != nil {
		return nil
	}
	writer.Reset(nil) // Reset the writer state before returning it to the pool
	writerPool.Put(writer)
	return buf.Bytes()
}

var (
	writerPool2 = sync.Pool{
		New: func() interface{} {
			return zlib.NewWriter(nil)
		},
	}
	bufferPool2 = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func DoZlibCompress2Pool(data []byte) ([]byte, error) {
	buf := bufferPool2.Get().(*bytes.Buffer)
	buf.Reset()
	writer := writerPool2.Get().(*zlib.Writer)
	writer.Reset(buf)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	writer.Reset(nil) // Reset the writer state before returning it to the pool
	writerPool.Put(writer)
	compressedData := buf.Bytes()
	bufferPool2.Put(buf)
	return compressedData, nil
}
