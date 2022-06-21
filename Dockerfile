FROM golang:alpine3.16

WORKDIR /.

COPY ../data.json ./

COPY ./app/* ./

COPY ./go.mod ./

COPY ./go.sum ./

RUN go get ./

EXPOSE 25585

CMD ["go", "run", "."]