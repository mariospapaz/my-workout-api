FROM golang:1.18.3-alpine3.16

WORKDIR /.

COPY .mongo/data.json ./

COPY ./app/* ./

RUN go get "go.mongodb.org/mongo-driver/mongo"

RUN go get "github.com/gin-gonic/gin"

ENV PORT=25585

EXPOSE 25585

CMD ["go", "run", "."]