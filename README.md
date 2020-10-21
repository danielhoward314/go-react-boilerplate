# Go-React Boilerplate

## Golang API and Web Server for React SPA

### Motivation & Design
<p>Many boilerplates exist for serving up a React/Vue/Angular SPA with Express, but I haven't seen many with Go acting as the webserver. With this project I hope to give other developers a good starting point.</p><p>I designed this app with an eye toward use within a microservice architecture. The Go webserver serves up the webpack bundle of the React SPA. Any XHR requests should hit this app's Go API server and get forwarded to an API Gateway, which would perform cross-cutting concerns (load balancing, rate limiting, distributed tracing, APM, etc.) and route the calls to a fleet of microservices. Long-running workloads could get triggered through messaging systems (Kafka, RabbitMQ, SQS, NATS, etc.) and run asynchronously.</p>

### Goals
<ol><h3>This app can: </h3>
<li>locally serve a React SPA on webpack-dev-server and proxy API requests to the Golang server's API and provide hot reloading with updates to the React code;</li>
<li>bundle the React app with Webpack, locally serve the resulting static assets with a Golang webserver, respond to the browser XHR requests with a Golang API, and configure CORS within the server to whitelist domains;</li>
<li>create a Dockerized production build of (2) and run that locally;</li>
<li>run automated deployment of (3);</li>
<li>configure https and http;</li>
<li>configure cache control headers;</li>
<li>configure Webpack for dev, docker and mimic production, with html-from-ejs-template generation, loaders for .scss files and other asset file types.
</ol>

# How to Run This App
## Options
<p>This app offers several options for running: </p>
<ol>
<li>run React app with webpack-dev-server and Go as API-only;</li>
<li>run Go as webserver for Webpack bundle and API handling XHR requests from the bundled SPA;</li>
<li>run a Dockerized version of (1) with separate containers for the React SPA and Go API server;</li>
<li>run Dockerized version of (2), that serves the dual purpose of building the production image and enabling local debugging of that image.</li>
</ol>

## Recommendation for local development
<p>Though this repo provides Dockerfile.dev and docker-compose.dev.yml capabilites, which may be adapted for use with mounting to local data sources, I recommend developing with method (1).</p><p>webpack-dev-server ships with HMR and Go has several popular packages (e.g. gin) for hot reloading. </p>

## Commands for Each Option

Prerequisites for non-Dockerized version:
<ul>
  <li>Go installed with +/= version in go.mod;</li>
  <li>Local system has correct configuration of GOPATH and its child directories: bin, pkg, src; this repo is cloned into src;</li>
  <li>Node installed (12.19+);</li>
  <li>npm installed (6+).</li>
</ul>
<ol>
  <li>For method (1), run the following commands from the root directory...

  ```
  go get
  cd client && npm install
  cd ..
  go run main.go -env=dev -port=8080 & cd client && npm run start-client
  ```

  ...and go to `http://localhost:3000` to see the React app, and open the console to see a succesful XHR call proxied to the Go API.
  </li>
  <li>For method (2), run the following (assumes already ran `go get and npm install` as above):

  ```
  go run main.go -env=prod -port=8080 -html=$(PWD)/client/dist/index.html -webpack=$(PWD)/client/dist/js/

  // in another terminal tab
  cd client
  npm run build-client
  ```

  ...and go to `http://localhost:8080` to see the same results as (1), except these are all coming from the Go server rather than from webpack-dev-server and Go server separately.
  </li>
    <li>For method (3), run the following...

  ```
  docker-compose -f docker-compose.dev.yml up --build
  ```

  ...and go to `http://localhost:3000`.
  </li>
  <li>For a Dockerized version of goal (2)...

  ```
  docker-compose up --build
  ```

  ...and go to `http://localhost:80`.
  </li>
</ol>

# Manual Deployment & CI/CD

### Manual Docker Push

Assumes Docker Hub as registry, though principles are the same for other registries. If you've logged into Docker Hub from the command line, Docker Desktop should have generated a `~/.docker/config.json` file with base64 encoded credentials. If you run into Docker Hub auth issues, [this article can help](https://mesosphere.github.io/marathon/docs/native-docker-private-registry.html).

Edit the IMAGE_TAG var in the Makefile to point to your Docker Hub repo.

```
make build && make push
```

You should have a built and tagged image, which you can manaully deploy to the container service of a cloud provider. Though it is out of scope for this README to outline the manual version of these steps, the following sections outline an automated CI/CD pipeline to deploy to EKS.

### CI with CircleCI

<p>Note that below steps require upgrading CircleCI plan, since the CI process needs access to the Docker daemon, a Performance plan and higher feature. If you don't want to use CircleCI for this reason, delete the `.circleci` directory and set up another CI pipeline of your choice.
<ol>
  <li>Sign up for CircleCI and link your source control repo in its UI.</li>
  <li>Set up a `DOCKERHUB_PASSWORD` environment variable in CircleCI.</li>
  <li>Edit the `.circleci/config.yml` file as instructed by the comments.</li>

