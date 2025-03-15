package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/go-ping/ping"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	icmpSuccess = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "icmp_ping_success",
		Help: "1 if the last ping was successful, 0 if it failed.",
	})
	icmpResponseTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "icmp_ping_response_time",
		Help: "Response time of the last ICMP ping in milliseconds.",
	})
	httpResponseTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "http_response_time",
		Help: "Response time of the last HTTP GET request in milliseconds.",
	})
)

func init() {
	// Create a custom registry to register ONLY our custom metrics
	registry := prometheus.NewRegistry()

	// Register only our metrics
	registry.MustRegister(icmpSuccess, icmpResponseTime, httpResponseTime)

	// Create a custom HTTP handler for Prometheus metrics
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
}

// Ping function to measure ICMP response time
func pingHost(host string) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Printf("Error creating pinger: %v", err)
		icmpSuccess.Set(0)
		return
	}
	pinger.SetPrivileged(true) // Required for Windows!
	pinger.Count = 1
	pinger.Timeout = 2 * time.Second

	pinger.Run() // Blocking call
	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		icmpSuccess.Set(1)
		roundedPing := math.Round(stats.AvgRtt.Seconds()*1000.0*100) / 100
		icmpResponseTime.Set(roundedPing)
		fmt.Printf("Ping to %s: %.2f ms\n", host, roundedPing)
	} else {
		icmpSuccess.Set(0)
		icmpResponseTime.Set(0)
		fmt.Printf("Ping to %s failed\n", host)
	}
}

// Perform HTTP GET and measure response time
func httpGet() {
	start := time.Now()
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Printf("HTTP GET error: %v", err)
		httpResponseTime.Set(0)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start).Seconds() * 1000.0
	roundedHTTP := math.Round(duration*100) / 100
	httpResponseTime.Set(roundedHTTP)
	fmt.Printf("HTTP GET to google.com took %.2f ms\n", roundedHTTP)
}

func main() {
	// Get target hostname from command-line argument
	host := flag.String("host", "google.com", "Target host for ICMP ping")
	port := flag.String("port", "8080", "Port for Prometheus metrics server")
	flag.Parse()

	// Start Prometheus HTTP server
	go func() {
		log.Printf("Serving metrics at :%s/metrics", *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}()

	// Run Ping & HTTP GET in a loop
	for {
		pingHost(*host)
		httpGet()
		time.Sleep(5 * time.Second)
	}
}
