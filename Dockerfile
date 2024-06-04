FROM golang:1.22 as build
WORKDIR /api
COPY . .
RUN go mod download && go mod verify
RUN GOOS=linux GOARCH=amd64 go build
RUN ls

FROM alpine:latest
EXPOSE 8080
WORKDIR /
COPY --from=build /api/rssProxy  .
COPY --from=build /api/assets ./assets
COPY --from=build /api/views ./views
RUN ls
CMD ["./rssProxy"]