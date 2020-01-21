FROM alpine:3
RUN mkdir /app
WORKDIR /app
COPY ./bin/app .
COPY ./config.yml .
EXPOSE 8080
CMD ./app
