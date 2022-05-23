<template>
  <div class="video">
    <v-expand-transition>
    <v-alert v-show="method=='create' && userList.length <= 1">这个房间没有人哦，您可以把这个房间名「{{room}}」分享给朋友</v-alert>
    </v-expand-transition>
    <v-expand-transition>
    <v-alert v-show="method=='create' && userList.length > 1">房间「{{room}}」中有 {{userList}} 正在一起观影</v-alert>
    </v-expand-transition>
    <v-expand-transition>
    <div v-show="method=='create'||otherUrl">
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
        <v-btn @click="createRoom()" >设置房间</v-btn></div>
        
        <div v-show="method=='join' && url!='' && user!=''">
        <v-btn @click="joinRoom()" >{{joinButText}}</v-btn></div>
        
      </v-col>
      <v-col>
        <div v-show="method!='x' && url!='' && user!=''">
        <v-btn @click="syncVideo()" >手动同步</v-btn>
        </div>
      </v-col>
      <v-col>
        <div v-show="method=='join' && url!='' && user!='' && start">
        <v-switch v-model="otherUrl" color="red" label="独立链接模式"/>
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
        <source :src="url" :type="videoType" /></video
    ></div>
    </v-expand-transition>
  </div>
</template>

<script>
import Video from "video.js";
import io from "socket.io-client";
import url from "url";

let myPlayer;
const socket = io.connect("http://127.0.0.1:8000");


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
      videoType: "video/mp4",
      otherUrl: false,
      joinButText: "加入房间",
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
        this.method = "join"
      }
      if(this.userList.length==0){
        this.method = "create";
      }
      this.getUrl();
    })
    socket.on("play",(data)=>{
      if(myPlayer.readyState()>1 && data!=this.user){
        myPlayer.play()
      }
      this.syncVideo()
    })
    socket.on("pause",(data)=>{
      if(data!=this.user){
        myPlayer.pause()
      }
      this.syncVideo()
    })
    socket.on("Null",(data)=>{
      this.method = "create";
    })
    socket.on("setUrl", (data) => {
      if(this.otherUrl){
        this.joinButText="覆盖URL"
      }else{
        this.joinButText="加入房间"
      }
      if (this.method=="create"){
        return
      }
      if(data!=this.url){
        if (!this.otherUrl){
          this.url = data;
        }
        if (this.user!="" && !this.otherUrl){
          this.changeSrc()
        }
      }
    });
    socket.on("setTime", (data) => {
      let showName = data.split(":::")[0];
      let newTime = data.split(":::")[1];
      myPlayer.off('seeking')
      if (showName == this.user){
        myPlayer.on('seeking', () =>{
          socket.emit('setTime',`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()))
        })
        return
      }
      if (Math.abs(parseFloat(newTime) - parseFloat(myPlayer.currentTime()))>1) {
        myPlayer.one('seeked',()=>{
          myPlayer.play()
        })
        myPlayer.currentTime(parseFloat(newTime));
      }
      myPlayer.on('seeking', () =>{
        socket.emit('setTime',`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()))
      })
    })
    socket.on("join", (data)=>{
      this.getTime();
      this.getUsers();
      this.getUrl();
    })
    socket.on("sync", (data) => {
      let myTime= JSON.stringify(myPlayer.currentTime())
      let showName = data.split(":::")[2];
      if (parseFloat(myTime)<parseFloat(data.split(":::")[1]) && Math.abs(parseFloat(myTime)-parseFloat(data.split(":::")[1]))>1) {
        myPlayer.currentTime(parseFloat(data.split(":::")[1]));
      }
    });
    socket.on("getTime", (data) => {
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
      if(this.method=="create"){
        return
      }
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
      let urlData=url.parse(this.url);
      let tmpVideoType = urlData.pathname.split(".").pop();
      if(tmpVideoType=="mp4"){
        this.videoType = "video/mp4";
      }else if(tmpVideoType=="webm"){
        this.videoType = "video/webm";
      }else if(tmpVideoType=="ogg"){
        this.videoType = "video/ogg";
      }else if(tmpVideoType=="m3u8"){
        this.videoType = "application/x-mpegURL";
      }else{
        this.videoType = "video/mp4";
      }
      myPlayer.src({
        src: this.url,
        type: this.videoType,
      });
      // myPlayer.on('timeupdate',  () => {
      //   let tmpTime=myPlayer.currentTime()
      //   if(tmpTime==0) {
      //     return;
      //   }
      //   if(tmpTime - this.currentTime > 2 || tmpTime - this.currentTime < -2){
      //     socket.emit('setTime',`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()))
      //     myPlayer.pause()
      //   }
      //   this.currentTime = myPlayer.currentTime()
        
      // })
      myPlayer.on('seeking', () =>{
        socket.emit('setTime',`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()))
      })
      myPlayer.on('pause',  () => {
        if(myPlayer.readyState()!=0){
          socket.emit('pause',`${this.room}:::${this.user}`)
        }
      })
      myPlayer.on('play',  () => {
        if(myPlayer.readyState()>1){
          socket.emit('play',`${this.room}:::${this.user}`)
        }
      })
      this.getTime();
    },
    join() {
      socket.emit("join",`${this.room}:::${this.user}`);
    },
    syncVideo() {
      socket.emit("sync", `${this.room}:::${this.user}`);
    },
    getTime() {
      socket.emit("time",`${this.room}:::${this.user}:::`+JSON.stringify(myPlayer.currentTime()));
    },
    sendTime() {
      this.getTime();
    },
    sendTime() {
      this.getTime();
    },
    getUsers() {
      socket.emit("getUsers",`${this.room}:::${this.user}`)
    },
    getUrl() {
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