module.exports = {
	publicPath: "/movie",
	devServer: {
		// proxy: {
		// 	'/api': {
		// 		target: 'http://192.168.123.151:8000/ws',
		// 		changeOrigin: true,
		// 		ws: true,
		// 		pathRewrite: {
		// 			'^/api': ''
		// 		}
		// 	},
		// 	'/socket.io': {
		// 		target: 'http://192.168.123.151:8000/',
		// 		changeOrigin: true,
		// 		ws: true,
		// 	}
		// }
	},
	transpileDependencies: [
		'vuetify'
	]
}
