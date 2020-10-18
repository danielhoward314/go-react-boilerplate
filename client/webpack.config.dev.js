const path = require('path')
const { merge } = require('webpack-merge')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const baseConfig = require('./webpack.config')
const devServerPort = 3000
const proxyHeaders = {
  'x-forwarded-proto': 'http',
  'x-forwarded-port': devServerPort,
  'host': '127.0.0.1'
}

module.exports = merge(baseConfig, {
  mode: 'development',
  output: {
    filename: './bundle.js',
    path: path.resolve(__dirname, 'dist')
  },
  devServer: {
    port: devServerPort,
    publicPath: '/',
    contentBase: './static',
    hot: true,
    quiet: false,
    proxy: {
      '/api/*': { target: 'http://localhost:8080', secure: false, headers: proxyHeaders }
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
