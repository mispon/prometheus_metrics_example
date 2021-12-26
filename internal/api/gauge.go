package api

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/mispon/prometheus_metrics_example/pkg/metrics/github.com/mispon/prometheus_metrics_example"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	gaugeCounter = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "sandbox",
			Name:      "rand_current",
		},
	)
)

func init() {
	_ = prometheus.Register(gaugeCounter)
	rand.Seed(time.Now().Unix())
}

func (s server) Gauge(context.Context, *emptypb.Empty) (*pb.BaseResponse, error) {
	num := rand.Float64() * 100
	gaugeCounter.Set(num)
	return &pb.BaseResponse{Status: "ok"}, nil
}
