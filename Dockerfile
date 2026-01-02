FROM golang:1.25-alpine AS runtime
WORKDIR /hexbook

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN ls -al
RUN cd cmd && GOOS=linux GOARCH=amd64 go build -o ./server.run ./main.go

# @dev multi-stage bulid for less image size
FROM alpine:3.22 AS runner
WORKDIR /hexbook
RUN mkdir -p cmd templates assets logs
COPY --from=runtime /hexbook/cmd/server.run ./cmd
COPY --from=runtime /hexbook/templates/ ./templates
COPY --from=runtime /hexbook/assets/ ./assets
EXPOSE 3016

CMD ["./cmd/server.run"]
