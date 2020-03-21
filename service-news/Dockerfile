FROM golang:1.14 AS builder
ENV GOPROXY https://goproxy.cn
COPY . /build
WORKDIR /build/service-news
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]