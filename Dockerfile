# builder 源镜像
FROM golang:1.12.4-alpine as builder

#作者
MAINTAINER Pharbers "zyqi@pharbers.com"

# 安装系统级依赖
RUN apk add --no-cache git gcc musl-dev mercurial bash gcc g++ make pkgconfig openssl-dev
ENV PKG_CONFIG_PATH /usr/lib/pkgconfig

# 下载 librdkafka
RUN git clone https://github.com/edenhill/librdkafka.git /app/librdkafka

WORKDIR /app/librdkafka
RUN ./configure --prefix /usr  && \
make && \
make install

# 以LABEL行的变动(version的变动)来划分(变动以上)使用cache和(变动以下)不使用cache
LABEL JobReg.version="1.0.0" maintainer="ClockQ"

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# prod 源镜像
FROM alpine:latest as prod

#作者
MAINTAINER Pharbers "zyqi@pharbers.com"

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=0 /app/ipaas-job-reg .

EXPOSE 9213
ENTRYPOINT ["/app/ipaas-job-reg"]
