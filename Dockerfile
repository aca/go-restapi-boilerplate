FROM golang:1.13.7-buster
WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN go build -v -o /bin/api /app/cmd/api

FROM golang:1.13.7-buster
COPY --from=0 /bin/api /bin/api
ENTRYPOINT ["api"]
