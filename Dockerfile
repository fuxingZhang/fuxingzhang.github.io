FROM alpine:latest

WORKDIR /app
COPY ./dist/server  .
COPY html ./html
COPY static ./static

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
#     && apk update

# RUN apk add tzdata \
#     && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#     && echo "Asia/Shanghai" > /etc/timezone

EXPOSE 80

ENTRYPOINT ["/app/server"]