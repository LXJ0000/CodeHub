FROM golang:1.21.4-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o bluebell_app .

#########################
# 接下来创建一个基础镜像
#########################
FROM scratch

COPY ./conf/config.yaml /conf

# 从builder镜像中把/build/main 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]