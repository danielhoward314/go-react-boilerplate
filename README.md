# Go-React Boilerplate

## Golang API and Web Server for SPA React App

### Goals
<ol>
<li>App can serve a React SPA on webpack-dev-server and proxy API requests to the Golang server's API. App can expect hot reloading with updates to the React code.</li>
<li>App can bundle the React app with Webpack, serve the resulting static assets with the Golang server's webserver and respond to the browser XHR requests with the Golang server's API. App configures CORS within the server to whitelist domains.</li>
<li>App can create a Dockerized production build of (2).</li>
<li>App can run automated deployment of (3).</li>
<li>App can use a .env file within the root directory for Go environment variables.</li>
<li>App can run a Dockerized version of (1) and (2).</li>
<li>App configures https and http.</li>
<li>App configures cache control headers.</li>
<li>App configures Webpack for dev and prod, with html-from-ejs-template generation, loaders for .scss files and other asset file types.
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
cd client
npm run build-client
```

...and go to `http://localhost:8080` to see the same results as (1), except these are all coming from the Go server rather than from webpack-dev-server and Go server separately.
</li>
<li>For Goal (3), run the following...

```
docker build . -t <tag>
docker run -d -p 8080:8080 <image>
```

...and go to `http://localhost:8080/api/v1/health` to test the api.
</li>
