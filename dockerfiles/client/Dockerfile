FROM golang as builder
WORKDIR /code
COPY go.mod go.sum client.go ./
COPY proto/ proto/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

FROM gcr.io/distroless/base-debian10
COPY --from=build /code/client /
CMD ["/app"]