const express = require('express');
const cors = require('cors')
const app = express();
// const path = require('path');
const http = require('http');
const server = http.createServer(app);
const io = require("socket.io")({
	allowEIO3: true
}).listen(server);
const config = require('./config');

app.use(cors())

io.on('connection', (socket) => {
	console.log(socket.id, "is connected")
	io.on('message', data => {
		console.log(data)
	})
});

server.listen(config.PORT, () => {
	console.log('listening on *:', config.PORT);
});