package http

import (
	"context"
	"github.com/cloudwego/netpoll"
	"time"
)

func NetPollPlay() {
	//set poller number
	netpoll.SetNumLoops(5)

	// Create listener
	listener, err := netpoll.CreateListener("tcp", "127.0.0.1:8888")
	if err != nil {
		panic("create netpoll listener fail")
	}
	// handler assign handler
	var onRequest netpoll.OnRequest = handler

	// options: EventLoop init options
	var opts = []netpoll.Option{
		netpoll.WithReadTimeout(1 * time.Second),
		netpoll.WithIdleTimeout(10 * time.Minute),
		netpoll.WithOnPrepare(nil),
	}

	// Create EventLoop
	eventLoop, err := netpoll.NewEventLoop(onRequest, opts...)
	if err != nil {
		panic("create netpoll event-loop fail")
	}

	// Run Server
	err = eventLoop.Serve(listener)
	if err != nil {
		panic("netpoll server exit")
	}
}

// Handler on connection
func handler(ctx context.Context, connection netpoll.Connection) error {
	//connection.Reader().Next()
	return connection.Writer().Flush()
}
