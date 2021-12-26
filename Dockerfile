FROM ubuntu AS builder

RUN apt update -y
RUN apt upgrade -y

RUN apt install -y sudo

RUN useradd -m -G sudo developer
RUN echo 'developer:developer' | chpasswd
USER developer

RUN echo developer | sudo -S DEBIAN_FRONTEND="noninteractive" apt install -y golang
RUN echo developer | sudo -S apt install -y ca-certificates && sudo update-ca-certificates
RUN echo developer | sudo -S apt install -y make protobuf-compiler

ENV GOPATH /home/developer/go
ENV PATH $PATH:/home/developer/go/bin

COPY . /home/developer/go/src/metrics_app
RUN echo developer | sudo -S chown -R developer /home/developer/

WORKDIR /home/developer/go/src/metrics_app

RUN make deps && make build


FROM scratch
COPY --from=builder /home/developer/go/src/metrics_app/bin/app /metrics_app/app
WORKDIR /metrics_app
EXPOSE 80-84
ENTRYPOINT ["./app"]