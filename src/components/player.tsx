import '@vidstack/react/player/styles/base.css';
import {
    MediaPlayer,
    MediaProvider,
    MediaPlayerInstance,
    isHLSProvider,
    MediaProviderAdapter,
    MediaProviderChangeEvent,
} from '@vidstack/react';
import { VideoLayout } from './video-control';
import { useEffect, useRef, useState } from 'react';
import { socket } from './socket';
import { ClientMessage, ServerMessage } from '@/lib/types/message';
import { $playerState, $userInfo, $userStatus } from '@/store/player';
import { useStore as useNanoStore } from '@nanostores/react';
import HLS from 'hls.js';

export const Player = ({ roomName }: { roomName: string }) => {
    const userinfo = useNanoStore($userInfo);
    let player = useRef<MediaPlayerInstance>(null)
    const playerState = useNanoStore($playerState)
    useEffect(() => {
        if (!player.current) {
            return
        }
        // Subscribe to state updates.
        return player.current.subscribe(({ paused, currentTime, seeking }) => {
            // console.log("player state: ", { paused, time: Math.ceil(currentTime) });
            if (seeking) {
                socket.emit("setTime", JSON.stringify({
                    username: userinfo?.username,
                    time: Math.ceil(currentTime),
                    room: roomName
                } as ClientMessage))
                return
            }
        });
    }, [userinfo, roomName, playerState?.url]);

    useEffect(() => {
        const interval = setInterval(() => {
            if (!player.current) {
                return
            }
            socket.emit("updateMyInfo", JSON.stringify({
                username: userinfo?.username,
                time: Math.ceil(player.current.currentTime),
                room: roomName,
                playing: !player.current.paused
            } as ClientMessage))
        }, 2000)
        return () => clearInterval(interval)
    }, [userinfo?.username])

    useEffect(() => {
        function onConnect() {
            console.log('connected');
        }

        function onDisconnect(e: any, d: any) {
            console.log('disconnected, reason', e, d);
            socket.connect();
        }

        function onRootInit(d: any) {
            const msg = JSON.parse(d) as ServerMessage
            console.log('root init', msg);
            $playerState.set({
                url: msg.url,
                inited: true
            })
        }

        function onRoomInfo(d: any) {
            console.log('get room info response', JSON.parse(d) as ServerMessage);
            const msg = JSON.parse(d) as ServerMessage
            if (!msg.userStatus) {
                return
            }
            $userStatus.set([
                ...msg.userStatus
            ])
            const mintime = Math.min(...msg.userStatus.map(user => user?.time ?? Infinity));
            if (player.current && Math.abs(player.current.currentTime - mintime) > 10) {
                player.current.currentTime = mintime
            }
        }

        function onPause() {
            console.log('pause');
            if (player.current) {
                player.current.pause();
            }
        }

        function onPlay() {
            console.log('play');
            if (player.current && player.current?.state.canPlay) {
                player.current.play();
            }
        }

        function onSetTime(d: any) {
            console.log('set time: ', JSON.parse(d) as ServerMessage);
            const msg = JSON.parse(d) as ServerMessage
            const distUser = msg.userStatus?.find((u) => u.username === msg.actionEmitter)
            if (distUser && distUser.username !== userinfo?.username) {
                if (player.current && distUser.time) {
                    player.current.currentTime = distUser.time
                }
            }
        }

        function onSetUrl(d: any) {
            console.log('set url: ', JSON.parse(d) as ServerMessage);
            const msg = JSON.parse(d) as ServerMessage
            $playerState.set({
                url: msg.url
            })
        }
        socket.on('connect', onConnect);
        socket.on('disconnect', onDisconnect);
        socket.on('rootinit', onRootInit);
        socket.on('roomInfo', onRoomInfo)
        socket.on('pause', onPause)
        socket.on('play', onPlay)
        socket.on('setTime', onSetTime)
        socket.on('setUrl', onSetUrl)

        return () => {
            socket.off('connect', onConnect);
            socket.off('disconnect', onDisconnect);
            socket.off('rootinit', onRootInit);
            socket.off('roomInfo', onRoomInfo)
            socket.off('pause', onPause)
            socket.off('play', onPlay)
            socket.off('setTime', onSetTime)
            socket.off('setUrl', onSetUrl)
        }
    }, [userinfo]);

    function onProviderChange(
        provider: MediaProviderAdapter | null,
        nativeEvent: MediaProviderChangeEvent,
    ) {
        if (isHLSProvider(provider)) {
            provider.library = HLS;
        }
    }

    return (
        <div>
            {playerState?.url && <MediaPlayer
                key={playerState.url}
                className="w-full aspect-video bg-slate-900 text-white font-sans overflow-hidden rounded-md ring-media-focus data-[focus]:ring-4"
                title="test"
                src={playerState.url}
                crossorigin
                playsinline
                onProviderChange={onProviderChange}
                ref={player}
            >
                <MediaProvider>
                </MediaProvider>
                <VideoLayout roomName={roomName} />
            </MediaPlayer>}
        </div>
    );
}