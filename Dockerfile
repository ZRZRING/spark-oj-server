FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai
COPY server ./server
COPY config/config.prod.yaml ./config.yaml
RUN chmod +x ./server
EXPOSE 8000
CMD ["./server"]