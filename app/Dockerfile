# syntax=docker/dockerfile:1

FROM golang:latest AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-backend

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go-backend /go-backend
# TODO just move assets, unecessary to copy project after binary
COPY . .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-backend"]