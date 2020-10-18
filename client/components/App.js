import React from 'react'

class App extends React.Component {

  componentDidMount() {
    console.log('before fetch')
    fetch('/api/v1/health')
      .then(response => response.json())
      .then(data => console.log(data));
    console.log('after fetch')
  }

  render () {
    return (
      <div>
        <h3>Proof of concept(stage two): </h3>
        <ul>
          <li>created a React SPA app;</li>
          <li>app runs on webpack-dev-server;</li>
          <li>api calls get proxied to go server's endpoints;</li>
          <li>go server configured to enable CORS from SPA origin;</li>
        </ul>
      </div>
    )
  }
}

export default App
