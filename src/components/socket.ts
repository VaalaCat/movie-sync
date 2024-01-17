import { io } from "socket.io-client";

export const socket = io("ws://localhost:9999", {
	transports: ["websocket"],
	autoConnect: true,
})