<template>
  <div class="video">
    <v-text-field label="Regular" v-model="url"></v-text-field>
    <v-btn @click="changeSrc()">加载</v-btn>
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
const socket = io.connect("http://192.168.8.106:8000");
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
  },
};
</script>

    <style scoped>
</style>