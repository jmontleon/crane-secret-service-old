FROM registry.access.redhat.com/ubi8/go-toolset:latest AS builder
WORKDIR $APP_ROOT/src/github.com/konveyor/crane-secret-service
COPY . .
RUN go build -o $APP_ROOT/crane-secret-service main.go

FROM registry.access.redhat.com/ubi8-minimal
WORKDIR /
COPY --from=builder /opt/app-root/crane-secret-service .
ENTRYPOINT ["/crane-secret-service"]
