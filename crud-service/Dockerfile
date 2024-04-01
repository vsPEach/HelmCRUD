FROM golang as build

WORKDIR src

COPY go.* .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/app .

FROM alpine:3.18.0

COPY --from=build go/src/bin/app .

EXPOSE 8000

ENTRYPOINT ["./app"]