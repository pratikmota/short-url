#Build a tiny docker image
FROM alpine:latest
RUN mkdir /app
COPY url-server /app

CMD ["/app/url-server"]