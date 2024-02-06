import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@radix-ui/react-label'
import { Inter } from 'next/font/google'
import { useRouter } from 'next/router'
import { useState } from 'react'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const router = useRouter()
  const [roomName, setRoomName] = useState('')

  return (
    <main
      className={`${inter.className}`}
    >
      <div className='flex justify-center items-center h-screen'>
        <div className='flex flex-col gap-3'>
          <Label className='font-bold text-center'>在线电影院</Label>
          <Input placeholder='请输入房间名' onChange={(e) => setRoomName(e.target.value)} />
          <Button onClick={() => {
            router.push(`/room/${roomName}`)
          }}>加入</Button>
        </div>
      </div>
    </main>
  )
}
