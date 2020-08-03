package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/jvrplmlmn/FooBarQuix/internal/service"
)

type Config struct {
	Host string
	Port string `default:"80"`

	ShutdownTimeout time.Duration `default:"5s"`
}

func main() {
	log.Println("Starting 'foobarqix' service ...")

	// Load the service configuration from environment variables
	var c Config
	if err := envconfig.Process("foobarqix", &c); err != nil {
		log.Fatalf("Failed to process config from environment variables: %s", err)
	}

	// Configure the HTTP multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/ready", ReadinessHandler)
	mux.HandleFunc("/", NewNumberProcessor(service.NewFooBarQix()).Handler)

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(c.Host, c.Port),
		Handler: mux,
	}

	// Start listening from connections and serve traffic
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Error shutting down server: %s", err)
		}
	}()

	// Capture the system signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive it
	<-signalChan
	log.Println("Shutdown signal received, exiting...")

	// Configure a shutdown timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.ShutdownTimeout)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown the server: %s", err)
	}
}

func ReadinessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

type NumberProcessor struct {
	processor service.Processor
}

func NewNumberProcessor(processor service.Processor) *NumberProcessor {
	return &NumberProcessor{processor: processor}
}

type Response struct {
	Result string `json:"result"`
}

func (h *NumberProcessor) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	number, err := strconv.Atoi(r.URL.EscapedPath()[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transformedNumber := h.processor.CalculateForNumber(number)
	resp, err := json.Marshal(Response{Result: transformedNumber})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
