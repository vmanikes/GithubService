# -------------------------------------------------------------------------------------------------------------------
# First Stage Application Builder
# -------------------------------------------------------------------------------------------------------------------
FROM golang:1.16-alpine as builder

COPY . /go/GithubSearch

WORKDIR /go/GithubSearch

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o main main.go

# -------------------------------------------------------------------------------------------------------------------
# Second Stage Application Container
# -------------------------------------------------------------------------------------------------------------------
FROM alpine as release

EXPOSE 8081
COPY --from=builder /go/GithubSearch/main /

CMD ["./main"]