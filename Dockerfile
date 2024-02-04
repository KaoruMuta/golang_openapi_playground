FROM golang:1.21

WORKDIR /app
COPY . /app

CMD [ "make", "run" ]