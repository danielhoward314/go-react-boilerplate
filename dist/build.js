console.log('path is /dist/build.js. next step is to replace this file with output of webpack build.')
console.log('start sample fetch from api')
fetch('/api/v1/health')
  .then(response => response.json())
  .then(data => console.log(data));
  console.log('end sample fetch from api')
