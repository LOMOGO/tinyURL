FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o tinyURL .

##由于我们只需要构建好后的二进制文件即可，因此这里再构建一个空镜像，用以减少镜像体积
#FROM scratch

#由于需要在docker-compose中的tinyurl中执行命令，因此不能再使用scratch
FROM alpine

COPY ./config /config
COPY ./log /log
COPY ./docs /docs

COPY --from=builder /build/tinyURL /

#因为要在mysql服务启动后再启动web服务，因此将该命令注释，改为在docker-compose.yml文件中执行
#ENTRYPOINT ["/tinyURL"]