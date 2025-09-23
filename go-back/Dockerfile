FROM golang:1.24.5

ARG GITHUB_TOKEN

COPY ./ /go/src/go-back/
WORKDIR /go/src/go-back/

RUN go get ./
RUN go test .\internal\service\ -v --coverprofile cover.txt
RUN go tool cover --func=cover.txt
RUN go build -o main .

CMD ["/go/src/go-back/main"]

EXPOSE 1111