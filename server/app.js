const express = require('express');
const cors = require('cors')
const app = express();
// const path = require('path');
const http = require('http');
const server = http.createServer(app);
const io = require("socket.io")({
  allowEIO3: true
}).listen(server);
const config = require('./config')

app.use(cors())
// app.use((req, res, next) => {
//   console.log('Time: ', Date.now());
//   next();
// });

// app.use(express.static(path.join(__dirname, '../dist')));

io.on('connection', (socket) => {
  console.log('a user connected');
});

server.listen(config.PORT, () => {
  console.log('listening on *:', config.PORT);
});