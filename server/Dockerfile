FROM golang:1.25 as runtime

WORKDIR /fatcat

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY . .
RUN ls -al
# make linux-compatible
RUN cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/server.run ./main.go

# @dev multi-stage bulid for less image size
FROM alpine:3.22 as runner
WORKDIR /fatcat
RUN mkdir -p src logs
COPY --from=runtime /fatcat/bin/server.run ./src
EXPOSE 3010

CMD ["./src/server.run"]
