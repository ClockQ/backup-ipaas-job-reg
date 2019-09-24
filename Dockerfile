# builder 源镜像
FROM golang:1.12.4-alpine as builder

#作者
MAINTAINER Pharbers "zyqi@pharbers.com"

# 安装 主要 依赖
RUN apk add --no-cache bash git

# 安装 rdkafka 依赖
RUN apk add --no-cache gcc g++ make pkgconfig openssl-dev \
&& git clone https://github.com/edenhill/librdkafka.git -b v1.1.0 $GOPATH/librdkafka \
&& cd $GOPATH/librdkafka/ \
&& ./configure --prefix /usr \
&& make \
&& make install

WORKDIR /app

# go mod 依赖下载, 提出来是利用缓存
COPY go.mod go.sum ./
RUN go mod download

# go build 编译项目
COPY . ./
RUN go build


# prod 源镜像
FROM alpine:latest as prod

#作者
MAINTAINER Pharbers "zyqi@pharbers.com"

# 安装 主要 依赖
RUN apk --no-cache add bash git

# 安装 rdkafka 依赖
RUN apk add --no-cache gcc g++ make pkgconfig openssl-dev \
&& git clone https://github.com/edenhill/librdkafka.git -b v1.1.0 /tmp/librdkafka \
&& cd /tmp/librdkafka/ \
&& ./configure --prefix /usr \
&& make \
&& make install \
&& apk del gcc g++ make pkgconfig openssl-dev

# 设置时区
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /app

# 提取执行文件
COPY --from=0 /app/ipaas-job-reg ./

EXPOSE 9213
CMD ["./ipaas-job-reg"]
