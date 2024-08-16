package cmd

import (
	"context"
	"fmt"
	stlog "log"

	grades "github.com/douhimed/grading-distributed-sys/grages"
	"github.com/douhimed/grading-distributed-sys/log"
	"github.com/douhimed/grading-distributed-sys/registry"
	"github.com/douhimed/grading-distributed-sys/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at url %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
