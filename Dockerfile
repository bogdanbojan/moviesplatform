FROM golang:alpine
WORKDIR moviesplatform
COPY . .
RUN apk add git && cd cmd && go build .

FROM alpine
COPY --from=0 /go/moviesplatform/cmd /go/moviesplatform/cmd
WORKDIR /go/moviesplatform/cmd
CMD ["./cmd"]
