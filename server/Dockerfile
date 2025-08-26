FROM golang:1.25 as runtime

WORKDIR /fatcat

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN ls -al
# make linux-compatible
RUN cd cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server.run ./main.go

# @dev multi-stage bulid for less image size
FROM alpine:3.22 as runner
WORKDIR /fatcat
RUN mkdir -p cmd logs
COPY --from=runtime /fatcat/cmd/server.run ./cmd
EXPOSE 3010

CMD ["./cmd/server.run"]
