FROM golang:1.15 AS builder
ADD ./main.go main.go
ADD ./server server
# Fetch dependencies
RUN go get -d -v

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-react-boilerplate -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /go-react-boilerplate /app/go-react-boilerplate
EXPOSE 8080
ENTRYPOINT ["/app/go-react-boilerplate", "-html=$(PWD)/client/dist/index.html", "-webpack=$(PWD)/client/dist/js/"]
CMD ["-env=prod"]
