FROM golang:alpine3.16 as builder
WORKDIR /build
COPY ./app/*.go .
COPY ./go.mod .
COPY ./go.sum .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o my-workout .

#RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o my-workout .


FROM alpine:3.16.0
WORKDIR /.

COPY --from=builder /build/my-workout .
COPY ./data.json ./

EXPOSE 8080

# executable
ENTRYPOINT [ "./my-workout" ]

# 3Mb, 300 milliseconds between allocation
CMD [ "3", "300" ]