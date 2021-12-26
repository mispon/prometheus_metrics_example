package api

import (
	"context"

	pb "github.com/mispon/prometheus_metrics_example/pkg/metrics/github.com/mispon/prometheus_metrics_example"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	requestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "sandbox",
			Name:      "requests_total",
		},
	)
)

func init() {
	_ = prometheus.Register(requestsCounter)
}

func (s server) Counter(context.Context, *emptypb.Empty) (*pb.BaseResponse, error) {
	requestsCounter.Inc()
	return &pb.BaseResponse{Status: "ok"}, nil
}
