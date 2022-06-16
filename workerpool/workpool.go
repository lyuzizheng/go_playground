package workerpool

import "sync"

type WorkerPool struct {
	workerFunction func(interface{})
	workerChannelPool sync.Pool
	workersCount int


}
