package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		configPath = flag.String("config", "", "path to config file")
		strategy   = flag.String("strategy", RoundRobin, "strategy ogf the load balancer")
	)
	flag.Parse()

	config := MustParseConfig(*configPath)
	port := fmt.Sprintf(":%v", config.Port)

	log.Printf("balancing from port %v", port)
	http.ListenAndServe(port, NewBalancer(*strategy, config.Hosts))

}
