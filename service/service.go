package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/douhimed/grading-distributed-sys/registry"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	return ctx, err
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server

	srv.Addr = ":" + port

	// Starting the server
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// Gracefully shutdown server
	go func() {
		fmt.Printf("%v started. Press any key to stop.  \n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
