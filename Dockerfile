FROM golang:1.12 AS build

WORKDIR /src/
COPY main.go /src/
COPY go.mod go.sum /src/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
COPY index.html /bin/
COPY envfile /bin/
WORKDIR /bin/
ENTRYPOINT ["/bin/demo"]