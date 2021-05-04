package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"ts3exporter/api"
	"ts3exporter/collector"
)

func main() {
	url, ok := os.LookupEnv("TS3_API_URL")
	if !ok {
		url = "http://127.0.0.1:10080"
	}

	apiKey, ok := os.LookupEnv("TS3_API_KEY")
	if !ok {
		log.Fatal("TS3_API_KEY missing")
	}

	client := api.NewClient(url, apiKey)

	prometheus.MustRegister(collector.NewServerInfo(client))

	http.Handle("/metrics", promhttp.Handler())
	log.Print("Start ts3exporter")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
