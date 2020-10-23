const path = require('path')
const { merge } = require('webpack-merge')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const baseConfig = require('./webpack.config')
const devServerPort = 3000
// references to go-server below correspond to links alias in spa service of docker-compose.dev.yml
const goServerContainerHost = 'http://go-server:3001'
const proxyHeaders = {
  'x-forwarded-proto': 'http',
  'x-forwarded-port': devServerPort,
  'host': 'go-server'
}

const configObj = merge(baseConfig, {
  mode: 'development',
  output: {
    filename: './bundle.js',
    path: path.resolve(__dirname, 'dist')
  },
  devServer: {
    port: devServerPort,
    publicPath: '/',
    host: '0.0.0.0',
    contentBase: path.join(__dirname, 'static'),
    hot: true,
    quiet: false,
    proxy: {
      '/api/**': { target: goServerContainerHost, changeOrigin: true, secure: false, headers: proxyHeaders }
    }
  },
  watchOptions: {
    aggregateTimeout: 1000,
    poll: 3000,
    ignored: ['node_modules/**']
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: 'index.html',
      inject: false,
      title: 'Go-React Boilerplate',
      template: path.join(__dirname, 'assets', 'index-template.ejs')
    })
  ]
})
console.log(configObj)
module.exports = configObj
