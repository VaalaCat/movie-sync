import { Player } from '@/components/player';
import { socket } from '@/components/socket';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { UserList } from '@/components/user-list';
import { ClientMessage } from '@/lib/types/message';
import { $playerState, $userInfo } from '@/store/player';
import { useStore } from '@nanostores/react';
import { useRouter } from 'next/router'
import { useState } from 'react';

export default function Page() {
    const router = useRouter()
    const userInfo = useStore($userInfo);
    const playerState = useStore($playerState);
    const [url, setUrl] = useState<string | undefined>();

    const roomName = router.query.room as string
    return (
        <div className='flex flex-col lg:flex-row m-2'>
            {roomName && <div className='lg:w-3/4 lg:mr-2'>
                <Player roomName={roomName} />
            </div>}
            <div className='lg:w-1/4 mb-1 border rounded'>
                {roomName && <div className='m-2 flex flex-row gap-2'>
                    <Input onChange={(e) => {
                        $userInfo.set({ username: e.target.value })
                    }} />
                    <Button onClick={() => {
                        socket.connect();
                        socket.emit('join', JSON.stringify({
                            username: userInfo?.username,
                            room: roomName,
                        } as ClientMessage));
                    }}>加入</Button>
                </div>}
                {userInfo?.username && <div className='m-2 flex flex-row gap-2'>
                    <Input key={playerState?.url} defaultValue={playerState?.url} onChange={(e) => {
                        setUrl(e.target.value)
                    }} />
                    <Button onClick={() => {
                        $playerState.set({ ...playerState, url })
                        socket.emit('setUrl', JSON.stringify({
                            room: roomName,
                            username: userInfo?.username,
                            url: url
                        } as ClientMessage))
                    }} >设置</Button>
                </div>}
                {roomName && <UserList roomName={roomName} />}
            </div>
        </div>)
}