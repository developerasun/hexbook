FROM golang:1.25-alpine as runtime
RUN apk add --no-cache gcc musl-dev sqlite-dev
WORKDIR /hexbook

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN ls -al
RUN cd cmd && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./server.run ./main.go

# @dev multi-stage bulid for less image size
FROM alpine:3.22 as runner
WORKDIR /hexbook
RUN mkdir -p cmd logs
COPY --from=runtime /hexbook/cmd/server.run ./cmd
EXPOSE 3010

CMD ["./cmd/server.run"]
