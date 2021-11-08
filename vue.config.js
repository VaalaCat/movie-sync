module.exports = {
  devServer: {
    proxy: 'http://localhost:8000'
  },
  transpileDependencies: [
    'vuetify'
  ]
}
