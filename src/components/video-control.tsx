import captionStyles from '@/styles/captions.module.css';
import styles from '@/styles/video-layout.module.css';

import { Captions, Controls, Gesture } from '@vidstack/react';

import * as Buttons from './video-buttons';
import * as Menus from './video-menus';
import * as Sliders from './video-sliders';
import { TimeGroup } from './video-time-group';
import { Title } from './video-title';

export interface VideoLayoutProps {
    thumbnails?: string;
    roomName: string
}

export function VideoLayout({ thumbnails, roomName }: VideoLayoutProps) {
    return (
        <>
            <Gestures />
            <Captions
                className={`${captionStyles.captions} media-preview:opacity-0 media-controls:bottom-[85px] media-captions:opacity-100 absolute inset-0 bottom-2 z-10 select-none break-words opacity-0 transition-[opacity,bottom] duration-300`}
            />
            <Controls.Root
                className={`${styles.controls} media-controls:opacity-100 absolute inset-0 z-10 flex h-full w-full flex-col bg-gradient-to-t from-black/10 to-transparent opacity-0 transition-opacity`}
            >
                <div className="flex-1" />
                <Controls.Group className="flex w-full items-center px-2">
                    <Sliders.Time thumbnails={thumbnails} />
                </Controls.Group>
                <Controls.Group className="-mt-0.5 flex w-full items-center px-2 pb-2">
                    <Buttons.Play roomName={roomName} tooltipPlacement="top start" />
                    <Buttons.Mute roomName={roomName} tooltipPlacement="top" />
                    <Sliders.Volume />
                    <TimeGroup />
                    <Title />
                    <div className="flex-1" />
                    <Buttons.Caption roomName={roomName}  tooltipPlacement="top" />
                    <Menus.Settings placement="top end" tooltipPlacement="top" />
                    {/* <Buttons.PIP tooltipPlacement="top" /> */}
                    <Buttons.Fullscreen roomName={roomName}  tooltipPlacement="top end" />
                </Controls.Group>
            </Controls.Root>
        </>
    );
}

function Gestures() {
    return (
        <>
            <Gesture
                className="absolute inset-0 z-0 block h-full w-full"
                event="pointerup"
                action="toggle:paused"
            />
            <Gesture
                className="absolute inset-0 z-0 block h-full w-full"
                event="dblpointerup"
                action="toggle:fullscreen"
            />
            <Gesture
                className="absolute left-0 top-0 z-10 block h-full w-1/5"
                event="dblpointerup"
                action="seek:-10"
            />
            <Gesture
                className="absolute right-0 top-0 z-10 block h-full w-1/5"
                event="dblpointerup"
                action="seek:10"
            />
        </>
    );
}
