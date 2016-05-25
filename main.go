package main

import (
	"log"
	"github.com/adamdecaf/aws-ecs-nginx-proxy/ecs"
	"github.com/adamdecaf/aws-ecs-nginx-proxy/nginx"
	"github.com/adamdecaf/aws-ecs-nginx-proxy/scheduler"
)

func main() {
	log.Println("spawning aws-ecs-nginx-proxy")

	nginx, err := nginx.New()
	if nginx == nil || err != nil {
		log.Printf("error spawning nginx err=%s\n", err)
	}

	log.Println(nginx)

	ecs, err := ecs.NewClient()
	if ecs == nil || err != nil {
		log.Printf("error starting ecs client err=%s\n", err)
	}

	log.Println(ecs)

	scheduler, err := scheduler.New(*nginx, *ecs)
	if scheduler == nil || err != nil {
		log.Printf("error starting scheduler err=%s\n", err)
	}

	scheduler.Run()
}
