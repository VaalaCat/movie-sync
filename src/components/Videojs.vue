<template>
  <div class="video">
    <v-expand-transition>
    <v-alert v-show="method=='create' && userList.length <= 1">这个房间没有人哦，您可以把这个房间名「{{room}}」分享给朋友</v-alert>
    </v-expand-transition>
    <v-expand-transition>
    <v-alert v-show="method=='create' && userList.length > 1">房间「{{room}}」中有 {{userList}} 正在一起观影</v-alert>
    </v-expand-transition>
    <v-expand-transition>
    <div v-show="method=='create'">
    <v-text-field label="视频地址" v-model="url" ></v-text-field>
    </div>
    </v-expand-transition>
    <v-expand-x-transition>
    <v-alert v-show="method=='join'">房间「{{room}}」中有 {{userList}}，<br />大家正在快乐的观看:<br /> {{url}}</v-alert>
    </v-expand-x-transition>
    <!-- <v-text-field label="房间号" v-model="room"></v-text-field> -->
    <v-expand-transition>
    <div v-show="url!=''">
    <v-text-field label="用户名" v-model="user"></v-text-field>
    </div>
    </v-expand-transition>
    <!-- <v-btn @click="changeSrc()">加载</v-btn>
    <v-btn @click="getTime()">获取时间</v-btn>
    <v-btn @click="stopVideo()">暂停</v-btn>
    <v-btn @click="playVideo()">播放</v-btn>
    
    <v-btn @click="getUsers()">显示所有用户</v-btn>
    <v-btn @click="getUrl()">获取视频地址</v-btn>
    <v-btn @click="setUrl()">设置视频地址</v-btn>
    <v-btn @click="join()">加入</v-btn> -->
    <v-expand-transition>
    <v-row class="d-flex justify-center">
      <v-col>
        <div v-show="method=='create' && url!='' && user!=''">
        <v-btn @click="createRoom()" >设置房间</v-btn>
        </div>
        <div v-show="method=='join' && url!='' && user!=''">
        <v-btn @click="joinRoom()" >加入房间</v-btn>
        </div>
      </v-col>
      <v-col>
        <div v-show="method!='x' && url!='' && user!=''">
        <v-btn @click="syncVideo()" >手动同步</v-btn>
        </div>
      </v-col>
    </v-row>
    </v-expand-transition>
    <v-expand-transition>
    <div v-show="method=='x'">
    <v-progress-circular
      indeterminate
      color="primary">
    <br/>
    <br/>
    <br/>
    Loading...
    </v-progress-circular>
    </div>
    </v-expand-transition>
    <v-expand-transition>
    <div v-show="start" style="margin: 30px;">
      <video id="myVideo" class="video-js vjs-theme-forest">
        <source :src="url" type="video/mp4" /></video
    ></div>
    </v-expand-transition>
  </div>
</template>

<script>
import Video from "video.js";
import io from "socket.io-client";
let myPlayer;
const socket = io.connect();


/* eslint-disable */
export default {
  name: "Videojs",
  prop: ["videosrc"],
  data() {
    return {
      start: false,
      url: "",
      time: 0,
      user: "",
      userList: [],
      room: "room",
      method: "x",
      timer: null,
      sendTimer: null,
      createTimer: null,
      currentTime: 0,
    };
  },
  mounted() {
    this.room = this.$route.params.roomName;

    socket.on("allUsers",(data)=>{
      this.userList = data.split(",").slice(0, -1);
      if(this.userList.indexOf(this.user)!=-1 && this.userList.length==1 && this.method=="join"){
        this.method = "create";
      }
      if(this.userList.indexOf(this.user)==-1 && this.userList.length>=1 && this.method!="join"){
        this.method = "join";
      }
      if(this.userList.length==0){
        this.method = "create";
      }
    })
    socket.on("Null",(data)=>{
      this.method = "create";
    })
    socket.on("setUrl", (data) => {
      this.url = data;
      console.log(data);
    });
    socket.on("setTime", (data) => {
      let showName = data.split(":::")[0];
      let newTime = data.split(":::")[1];
      if (showName == this.user){
        return
      }
      if (Math.abs(parseFloat(newTime) - parseFloat(myPlayer.currentTime()))>2) {
        console.log("setTime");
        myPlayer.currentTime(parseFloat(newTime));
      }
    })
    socket.on("join", (data)=>{
      console.log("joined",data.split(":::"))
      this.getTime();
      this.getUsers();
    })
    socket.on("sync", (data) => {
      let myTime= JSON.stringify(myPlayer.currentTime())
      console.log("data:",data)
      console.log("myTime:",myTime)
      let showName = data.split(":::")[2];
      console.log(showName)
      if (parseFloat(myTime)<parseFloat(data.split(":::")[1]) && Math.abs(parseFloat(myTime)-parseFloat(data.split(":::")[1]))>2) {
        console.log("change")
        myPlayer.currentTime(parseFloat(data.split(":::")[1]));
      }
    });
    socket.on("getTime", (data) => {
      console.log("getTime: " + data)
      socket.emit("time",`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()));
    })
    socket.on("leaveRoom", (data)=>{
      this.getUsers();
    })
    this.timer = setInterval(() => {
      if (this.start==true){
        this.syncVideo();
      }
    }, 10000);
    this.sendTimer = setInterval(() => {
      if (this.start==true){
        this.sendTime();
      }
    },1000);
    this.createTimer = setInterval(() => {
      this.getUsers();
      this.getUrl();
    },3000);
  },
  methods: {
    initVideo() {
      myPlayer = Video(myVideo, {
        controls: true,
        preload: true,
        // width: "640px",
        // fill: true,
        // responsive: true,
        fluid: true,
        // height: "200px",
      });
    },
    changeSrc() {
      this.start = true;
      this.initVideo();
      console.log("changeSrc");
      myPlayer.src({
        src: this.url,
        type: "video/mp4",
      });
      myPlayer.on('timeupdate',  () => {
        let tmpTime=myPlayer.currentTime()
        if(tmpTime==0) {
          return;
        }
        if(tmpTime - this.currentTime > 2 || tmpTime - this.currentTime < -2){
          console.log("xxxxxxxxxxxxxxxx")
          socket.emit('setTime',`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()))
        }
        this.currentTime =myPlayer.currentTime()
      })
      this.getTime();
    },
    join() {
      socket.emit("join",`${this.room}:::${this.user}`);
    },
    syncVideo() {
      console.log("syncVideo");
      socket.emit("sync", `${this.room}:::${this.user}`);
    },
    getTime() {
      console.log("videoTimeis:", myPlayer.currentTime());
      socket.emit("time",`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()));
    },
    sendTime() {
      this.getTime();
    },
    getUsers() {
      socket.emit("getUsers",`${this.room}:::${this.user}`)
    },
    getUrl() {
      console.log("getUrl");
      socket.emit("getUrl",`${this.room}`)
    },
    setUrl() {
      socket.emit("setUrl",`${this.room}:::${this.user}:::${this.url}`)
    },
    createRoom(){
      this.join();
      this.setUrl();
      this.getUrl();
      this.changeSrc();
    },
    joinRoom(){
      this.join();
      this.getUrl();
      this.changeSrc();
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