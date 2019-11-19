FROM golang:alpine

RUN apk update && apk add --no-cache git

ENV PORT 8080

WORKDIR /go/src/github.com/ritcrap/my-website/server

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/my-website-server

CMD [ "/go/bin/my-website-server" ]

