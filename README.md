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
<li>Node installed (10+)</li>
<li>npm install (6+)</li>
</ul>


<ol>
<li>For Goal(1), run the following commands from the root directory:

```
make init
make run separate
```
</li>
