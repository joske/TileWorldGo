FROM golang:alpine

WORKDIR /go/src/app

RUN apk update && apk add gtk+3.0 gtk+3.0-dev gcc musl-dev
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]
