<template>
  <div class="video">
    <v-text-field label="Regular" v-model="url"></v-text-field>
    <v-btn @click="changeSrc()">加载</v-btn>
    <v-btn @click="getTime()">获取时间</v-btn>
    <v-btn @click="stopVideo()">暂停</v-btn>
    <v-btn @click="playVideo()">播放</v-btn>
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
const socket = io.connect("/socket.io");

/* eslint-disable */
export default {
  name: "Videojs",
  prop: ["videosrc"],
  data() {
    return {
      start: false,
      url: "",
    };
  },
  mounted() {
    this.initVideo();
    socket.on("online", (name) => {
      console.log(name);
    });
    setTimeout(() => {
      this.getTime();
    }, 4000);
    socket.on("message", (data) => {
      console.log(data);
    });
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
    },
    getTime() {
      console.log("videoTimeis:", myPlayer.currentTime());
      socket.send("me", myPlayer.currentTime());
      return myPlayer.currentTime();
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