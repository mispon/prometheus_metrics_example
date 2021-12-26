package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mispon/prometheus_metrics_example/internal/api"
	metrics "github.com/mispon/prometheus_metrics_example/pkg/metrics/github.com/mispon/prometheus_metrics_example"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go runHTTP(ctx)
	go runGRPC(ctx)
	go runMetrics(ctx)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT)

	<-stopCh
	cancel()
}

func runHTTP(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong!\n"))
	})

	httpServer := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(ctx context.Context) {
	grpcServer := grpc.NewServer()
	apiServer := api.NewServer()

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
	}()

	metrics.RegisterMetricsServiceServer(grpcServer, apiServer)
	reflection.Register(grpcServer)

	listen, _ := net.Listen("tcp", ":82")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func runMetrics(ctx context.Context) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	metricsServer := http.Server{
		Addr:    ":84",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := metricsServer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := metricsServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
