# prometheus_metrics_example
Example setup for grpc server with prometheus and grafana

## Run
- `git pull git@github.com:mispon/prometheus_metrics_example.git`
- `cd prometheus_metrics_example`
- `docker-compose up -d`

## Endpoints
- http server:   
  `localhost:8080/ping`
- grpc server:   
  `localhost:8082`
- metrics:   
  `localhost:8084/metrics`
- prometheus:   
  `localhost:9000`
- grafana:   
  `localhost:3000`
