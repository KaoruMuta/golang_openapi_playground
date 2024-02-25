FROM golang:1.22

WORKDIR /app
COPY . /app

CMD [ "make", "run" ]