FROM golang:1.12.4-alpine3.9 as build-env
WORKDIR /go/src/github.com/saaresto/salo-location-suggester/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o salo-location-suggester


FROM alpine:3.7
WORKDIR /app
COPY --from=build-env /go/src/github.com/saaresto/salo-location-suggester/salo-location-suggester .
EXPOSE 8080
ENTRYPOINT ["./salo-location-suggester"]