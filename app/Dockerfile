# 멀티스테이지 빌드: golang:1.25-alpine에서 정적 바이너리를 만들고 scratch로 복사한다.
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /notiflex-api .

FROM scratch
COPY --from=build /notiflex-api /notiflex-api
EXPOSE 8080
ENTRYPOINT ["/notiflex-api"]
