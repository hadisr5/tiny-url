FROM golang:1.19-alpine AS build

RUN mkdir /app
ADD . /app/
WORKDIR /app



RUN GOPATH=/usr/go CGO_ENABLED=0 go build -o tinyurl .

FROM alpine:3.15

COPY --from=build /app/tinyurl /app/

RUN apk update && \
    apk add --update bash && \
    apk add --update tzdata && \
    cp --remove-destination /usr/share/zoneinfo/Asia/Tehran /etc/localtime && \
    echo "Asia/Tehran" > /etc/timezone && \
    chmod +x /app/tinyurl

EXPOSE 8080

CMD ["/app/tinyurl" ,"serve"]
