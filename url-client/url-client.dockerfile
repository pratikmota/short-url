#Build a tiny docker image
FROM alpine:latest
RUN mkdir /app
COPY url-client /app

CMD ["/app/url-client"]