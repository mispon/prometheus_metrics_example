package api

import (
	pb "github.com/mispon/prometheus_metrics_example/pkg/metrics/github.com/mispon/prometheus_metrics_example"
)

type (
	server struct {
		pb.UnimplementedMetricsServiceServer
	}
)

func NewServer() pb.MetricsServiceServer {
	return server{}
}
