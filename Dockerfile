FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY ./app.linux .
COPY .env .
RUN ["chmod", "+x", "./app.linux"]

CMD ["/app/app.linux"]