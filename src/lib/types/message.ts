export interface ClientMessage {
	username?: string;
	room?: string;
	url?: string;
	time?: number;
	playing?: boolean;
	payload?: string;
}

export interface ServerMessage {
	url?: string;
	userStatus?: UserStatus[];
	actionFrom?: string;
	actionEmitter?: string;
}

export interface UserStatus {
	username?: string;
	time?: number;
	playing?: boolean;
	userID?: string;
}

export interface PlayerState {
	paused?: boolean;
	url?: string;
	during?: number;
	inited?: boolean;
}