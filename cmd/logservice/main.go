package main

import (
	"context"
	"fmt"

	"github.com/douhimed/grading-distributed-sys/log"
	"github.com/douhimed/grading-distributed-sys/registry"
	"github.com/douhimed/grading-distributed-sys/service"

	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "8001"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
