# Go-React Boilerplate

## Golang API and Web Server for SPA React App

### Goals
<ol><h3>This app can: </h3>
<li>locally serve a React SPA on webpack-dev-server and proxy API requests to the Golang server's API. expect hot reloading with updates to the React code.</li>
<li>bundle the React app with Webpack, locally serve the resulting static assets with the Golang server's webserver and respond to the browser XHR requests with the Golang server's API, and configures CORS within the server to whitelist domains.</li>
<li>create a Dockerized production build of (2) and run that locally.</li>
<li>run automated deployment of (3).</li>
<li>configure https and http.</li>
<li>configure cache control headers.</li>
<li>configure Webpack for dev, docker and local prod mimic, with html-from-ejs-template generation, loaders for .scss files and other asset file types.
</ol>

### Methods of running locally

Prerequisites for non-Dockerized version:

<ul>
<li>Go installed with +/= version in go.mod</li>
<li>Local system has correct configuration of GOPATH and its child directories: bin, pkg, src; this repo is cloned into src</li>
<li>Node installed (10+)</li>
<li>npm install (6+)</li>
</ul>


<ol>
<li>For Goal (1), run the following commands from the root directory...

```
make init
make run separate
```

...and go to `http://localhost:3000` to see the React app, and open the console to see a succesful XHR call proxied to the Go API.
</li>
<li>For Goal (2), run the following (assumes already ran `make init`):

```
make test prod
// in another terminal tab
cd client
npm run build-client
```

...and go to `http://localhost:8080` to see the same results as (1), except these are all coming from the Go server rather than from webpack-dev-server and Go server separately.
</li>
<li>For a Dockerized version of goal (2)...

```
docker build . -t <tag>
docker run -d -p 8080:8080 <image>

OR

docker-compose up --build (optional -d to run in background)
```

...and go to `http://localhost:3000`.
</li>
<li>To run the React app and Golang API separately in Docker containers, run the following...

```
docker-compose -f docker-compose.dev.yml -f docker-compose.dev.yml up --build
```

...and go to `http://localhost:3000`.
</li>


