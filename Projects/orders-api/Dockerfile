FROM golang:1.12.0.0-alpine2.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN do build -o main .
CMD ["/app/main"]

