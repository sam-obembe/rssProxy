FROM golang:1.22-alpine as build
WORKDIR /api
COPY . .
RUN go mod download && go mod verify
RUN GOOS=linux GOARCH=amd64 go build
RUN ls

FROM alpine:latest
EXPOSE 8080
WORKDIR /rssProxy
COPY --from=build /api/rssProxy  .
RUN ls
CMD ["./rssProxy"]