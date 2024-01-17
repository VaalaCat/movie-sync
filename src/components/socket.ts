import { io } from "socket.io-client";

export const socket = io({
	transports: ["websocket"],
	autoConnect: true,
})