FROM golang:1.11-alpine AS build

WORKDIR /src/
COPY main.go /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
COPY index.html /bin/
COPY envfile /bin/
WORKDIR /bin/
ENTRYPOINT ["/bin/demo"]