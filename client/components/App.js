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
      <div>A blank canvas</div>
    )
  }
}

export default App
