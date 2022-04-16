<template>
  <div class="video">
    <v-text-field label="视频地址" v-model="url"></v-text-field>
    <v-text-field label="房间号" v-model="room"></v-text-field>
    <v-text-field label="用户名" v-model="user"></v-text-field>
    <v-btn @click="changeSrc()">加载</v-btn>
    <v-btn @click="getTime()">获取时间</v-btn>
    <v-btn @click="stopVideo()">暂停</v-btn>
    <v-btn @click="playVideo()">播放</v-btn>
    <v-btn @click="syncVideo()">同步</v-btn>
    <v-btn @click="getUsers()">显示所有用户</v-btn>
    <v-btn @click="join()">加入</v-btn>
    <v-div v-show="start"
      ><video id="myVideo" class="video-js">
        <source :src="url" type="video/mp4" /></video
    ></v-div>
    
  </div>
</template>

<script>
import Video from "video.js";
import io from "socket.io-client";
let myPlayer;
const socket = io.connect("127.0.0.1:8000");
socket.on("allUsers",(data)=>{
  console.log(data)
})

/* eslint-disable */
export default {
  name: "Videojs",
  prop: ["videosrc"],
  data() {
    return {
      start: false,
      url: "https://static.chive.vaa.la/html/video/式守同學不只可愛而已S1E1.mp4",
      time: 0,
      user: "vaala",
      room: "room",
    };
  },
  mounted() {
    this.initVideo();
    setTimeout(() => {
      this.getTime();
    }, 4000);
  },
  methods: {
    initVideo() {
      myPlayer = Video(myVideo, {
        controls: true,
        autoplay: "muted",
        preload: false,
        width: "200px",
        height: "200px",
      });
      console.log("stop");
    },
    changeSrc() {
      this.start = true;
      console.log("changeSrc");
      myPlayer.src({
        src: this.url,
        type: "video/mp4",
      });
      socket.on("getTime", (data) => {
        console.log("getTime: " + data)
        socket.emit("time",`${this.room}:${this.user}:`+JSON.stringify(myPlayer.currentTime()));
      })
    },
    join() {
      socket.on("sync", (data) => {
        let myTime= JSON.stringify(myPlayer.currentTime())
        console.log("data:",data)
        console.log("myTime:",myTime)
        if (parseInt(myTime)<parseInt(data.split(":")[1])) {
          myPlayer.currentTime(parseFloat(data.split(":")[1]));
        }
      });
      socket.emit("join",`${this.room}:${this.user}`);
    },
    syncVideo() {
      console.log("syncVideo");
      socket.emit("getTime", `${this.room}:${this.user}`);
    },
    getTime() {
      console.log("videoTimeis:", myPlayer.currentTime());
    },
    getUsers() {
      socket.emit("getUsers",`${this.room}:${this.user}`)
    },
    stopVideo() {
      myPlayer.pause();
    },
    playVideo() {
      myPlayer.play();
    },
  },
};
</script>

    <style scoped>
</style>