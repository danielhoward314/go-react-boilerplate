# Multi-stage build

# 1. Build Go binary.
FROM golang:1.15 AS goBuilder
ADD ./main.go main.go
ADD ./server server
# Fetch Go dependencies
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-react-boilerplate -a -tags netgo -ldflags '-s -w' .

# 2. Build webpack bundle of React SPA
FROM node:alpine as nodeBuilder
WORKDIR /webApp
COPY ["./client/package.json", "./"]
RUN npm install
COPY ["./client", "./"]
RUN npm run build-client-docker

# 3. Start from scratch image and add results of previous stages
FROM scratch
WORKDIR /app
COPY --from=goBuilder /go-react-boilerplate /app/go-react-boilerplate
COPY --from=nodeBuilder /dist /app/dist
EXPOSE 3001
ENTRYPOINT ["/app/go-react-boilerplate", "-html=/app/dist/index.html", "-webpack=/app/dist/js/"]
CMD ["-env=prod", "-port=3001"]
