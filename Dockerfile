FROM mongo:5

ENV API_PORT=25585

WORKDIR /.

COPY .mongo/data.json ./

COPY ./app/* ./

COPY ./go.mod ./

COPY ./go.sum ./

RUN apt update && apt upgrade --yes && apt install golang-go --yes && apt clean && rm -r /var/lib/apt/lists

RUN go get ./

EXPOSE ${API_PORT}

EXPOSE 27017

CMD ["go", "run", "."]