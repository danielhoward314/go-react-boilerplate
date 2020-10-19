const path = require('path')
const { merge } = require('webpack-merge')
const HtmlWebpackPlugin = require('html-webpack-plugin')

const baseConfig = require('./webpack.config')
const rootDirPath = path.dirname(__dirname)

module.exports = merge(baseConfig, {
  devtool: 'source-map',
  mode: 'production',
  output: {
    path: path.join(rootDirPath, 'dist'),
    publicPath: '/dist/',
    filename: 'js/[name].bundle.[chunkhash].js'
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
