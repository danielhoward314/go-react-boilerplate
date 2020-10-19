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
COPY ["./client/package.json", "./client/package-lock.json", "./"]
RUN npm install
COPY ["./client", "./"]
RUN echo Before Build && ls
RUN npm run build-client-docker
RUN echo After Build && ls

# 3. Start from scratch image and add results of previous stages
FROM scratch
WORKDIR /app
COPY --from=goBuilder /go-react-boilerplate /app/go-react-boilerplate
COPY --from=nodeBuilder /dist /app/dist
EXPOSE 8080
ENTRYPOINT ["/app/go-react-boilerplate", "-html=/app/dist/index.html", "-webpack=/app/dist/js/"]
CMD ["-env=prod"]
