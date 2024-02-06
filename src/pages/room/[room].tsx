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
    const [urlInput, setUrlInput] = useState<string | undefined>();

    const roomName = router.query.room as string
    return (
        <div className='flex flex-col lg:flex-row m-2 justify-center'>
            {roomName && url && <div className='w-full lg:mr-2'>
                <Player roomName={roomName} />
            </div>}
            <div className='w-full lg:w-[450px] mb-1 border rounded'>
                {roomName && <div className='m-2 flex flex-row gap-2'>
                    <Input onChange={(e) => {
                        $userInfo.set({ username: e.target.value })
                    }} placeholder='你的用户名' />
                    <Button onClick={() => {
                        socket.connect();
                        socket.emit('join', JSON.stringify({
                            username: userInfo?.username,
                            room: roomName,
                        } as ClientMessage));
                    }}>加入房间</Button>
                </div>}
                {userInfo?.username && <div className='m-2 flex flex-row gap-2'>
                    <Input key={playerState?.url} defaultValue={playerState?.url} onChange={(e) => {
                        setUrlInput(e.target.value)
                    }} placeholder='视频直链' />
                    <Button onClick={() => {
                        setUrl(urlInput)
                        $playerState.set({ ...playerState, url: urlInput })
                        socket.emit('setUrl', JSON.stringify({
                            room: roomName,
                            username: userInfo?.username,
                            url: urlInput
                        } as ClientMessage))
                    }} >修改链接</Button>
                </div>}
                {roomName && <UserList roomName={roomName} />}
            </div>
        </div>)
}