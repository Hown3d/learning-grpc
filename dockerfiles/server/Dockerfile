FROM golang as builder
WORKDIR code/
COPY server.go server.go
COPY proto/ proto/

RUN go build

FROM alpine as runner
WORKDIR server/
COPY --from=builder server ./server
CMD ./server
