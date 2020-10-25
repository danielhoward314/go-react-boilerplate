# Go-React Boilerplate

## Golang API and Web Server for React SPA

### Motivation & Design
<p>Many boilerplates exist for serving up a React/Vue/Angular SPA with Express, but I haven't seen as many with Go acting as the webserver. With this project I hope to give other developers a good starting point.</p><p>I designed this app with an eye toward use within a microservice architecture. The Go webserver serves up the webpack bundle of the React SPA. Any XHR requests should hit this app's Go API server and get forwarded to an API Gateway, which would perform cross-cutting concerns (load balancing, rate limiting, distributed tracing, APM, etc.) and route the calls to a fleet of microservices.</p>

### Goals
<ol><h3>This app can: </h3>
  <li>locally serve a React SPA on webpack-dev-server and proxy API requests to the Golang server's API and provide hot reloading with updates to the React code;</li>
  <li>bundle the React app with Webpack, locally serve the resulting static assets with a Golang webserver, respond to the browser XHR requests with a Golang API, and configure CORS within the server to whitelist domains;</li>
  <li>create a Dockerized production build of (2) and run that locally;</li>
  <li>run automated CI/CD of (3).</li>
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
<p>Though this repo provides Dockerfile.dev and docker-compose.dev.yml capabilites, which may be adapted for use with mounting to local data sources, I recommend developing with method (1).</p><p>webpack-dev-server ships with HMR and Go has several popular packages (e.g. gin) for hot reloading.</p>

## Commands for Each Option

Prerequisites for non-Dockerized version:
<ul>
  <li>go installed with +/= version in go.mod;</li>
  <li>local system has correct configuration of GOPATH and its child directories: bin, pkg, src; this repo is cloned into src;</li>
  <li>node installed (12.19+);</li>
  <li>npm installed (6+).</li>
</ul>

For any of the Dockerized options below, you'll need to install Docker desktop.

<ol>
  <li>For method (1), run the following commands from the root directory...

  ```
  go get
  go run main.go -env=dev -port=3001 & cd client && npm run start-client
  // in another terminal tab
  cd client && npm install
  ```

  ...and go to `http://localhost:3000` to see the React app, and open the console to see a succesful XHR call proxied to the Go API.
  </li>
  <li>For method (2), run the following (assumes already ran `go get and npm install` as above):

  ```
  go run main.go -env=prod -port=3001 -html=$(PWD)/client/dist/index.html -webpack=$(PWD)/client/dist/js/

  // in another terminal tab
  cd client
  npm run build-client
  ```

  ...and go to `http://localhost:3001` to see the same results as (1), except these are all coming from the Go server rather than from webpack-dev-server and Go server separately.
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

  ...and go to `http://localhost:3001`.
  </li>
</ol>

# Manual Deployment & Automated CI/CD

### Manual Docker Push

Assumes Docker Hub as registry, though principles are the same for other registries. If you've logged into Docker Hub from the command line, Docker Desktop should have generated a `~/.docker/config.json` file with base64 encoded credentials. If you run into Docker Hub auth issues, [this article can help](https://mesosphere.github.io/marathon/docs/native-docker-private-registry.html).

```
docker build -t <registry-username>/<repo-name>:latest .
docker push <registry-username>/<repo-name>:latest
```

You should have a built and tagged image, which you can manually deploy to the container service of a cloud provider. The rest of this README covers how to manually deploy to a local Kubernetes cluster and how to automate the same through a CI/CD pipeline in the GitOps style.

### Installations for Kubernetes and CI/CD

Note the following steps outline how to deploy to a local Kubernetes (k8s from hereon) cluster powered by minikube. The steps are similar for deployment to an external k8s cluster, though it requires modifications to some of the configuration files.

<ol>
  <li>install kubectl, the k8s command line tool;</li>
  <li>install minikube, a k8s tool for running a local cluster (assumes you already have a hypervisor, which 2017+ Macs have by default);</li>
  <li>install helm, the k8s package manager (packages are referred to as charts in helm world);</li>
  <li>install circleci CLI tool;</li>
  <li>install Argo CD.</li>
</ol>

### Manual Deployment to minikube

<ol>
  <li>run `minikube start`</li>
  <li>from project root directory, run `kubectl apply -f k8s/` (`k8s/` points to k8s folder where manifest yamls live);</li>
  <li>run `kubectl get deployments` to verify previous step and use the deployment name outputted in the next step;</li>
  <li>

  ```
  expose deployment <deployment-name> --type=NodePort --port=3001
  ```

  What goes into this command comes from configurations specified in the k8s manifest. The deployment name used corresponds to `metadata:name` attached to the Deployment, NodePort comes from `service:spec:type` in the service portion and port needs to correspond to both `containerPort` and `targetPort`.</li>
  <li>run `kubectl get services` to validate previous step and use whatever follows `service/` in the output for the next step;</li>
  <li>run...

  ```
  minikube service <service-name>
  ```

  ...and minikube should open a new browser tab with the app running.
  </li>
</ol>

### CI with CircleCI

<p>Note that the steps below require upgrading your CircleCI plan, since the CI process needs access to the Docker daemon, a Performance plan (or higher) feature. If you don't want to use CircleCI for this reason, delete the `.circleci` directory and set up another CI pipeline of your choice.
<ol>
  <li>Sign up for CircleCI and link your source control repo.</li>
  <li>Set up `DOCKERHUB_USERNAME` and `DOCKERHUB_PASSWORD` environment variables in CircleCI.</li>
  <li>Edit the `.circleci/config.yml` file as instructed by the comments.</li>
  <li>Note the `$CIRCLE_SHA1` in the config file. Since CircleCI is connected to Github, it can expose this environment variable which holds the SHA of the commit that kicked off the CI job. Tying CI/CD jobs (and the images and running deployments they produce, respectively) to commit SHAs is what powers GitOps. Argo CD exposes a similar environment variable for this purpose.</li>
</ol>

### CD with Argo CD

<ol>
  <li>run `kubectl port-forward -n argocd svc/argocd-server 8080:443`</li>
  <li>run `argocd login --insecure :8080` (initial username/password combo is `admin` and the argo-server name as it is generated in the argocd namespace during installation);</li>
  <li>go to `https://localhost:8080` and log in with admin credentials</li>
  <li>you can create an app and link the repo in the UI or if you run into trouble in the UI, run the following instead in the terminal from the root dir of this project:

  ```
  argocd app create <app-name> \
  --repo https://github.com/<username>/<repo-name>.git \
  --path "." \
  --dest-server https://kubernetes.default.svc \
  --dest-namespace default
  ```
  </li>
  <li>if you successfully create an app and connect your repo, you should have the option to sync based on new commits...so go ahead and sync the app;</li>
  <li>explore the dynamic visualization of the deployment;</li>
  <li>view the running app by doing steps 4-6 from the `Manual Deployment to minikube` section.</li>
</ol>
