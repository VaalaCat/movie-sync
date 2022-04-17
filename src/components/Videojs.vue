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
    <div v-show="method=='create' && url!='' && user!=''">
    <v-btn @click="createRoom()" >创建房间</v-btn>
    </div>
    </v-expand-transition>
    <v-expand-transition>
    <div  v-show="method=='join' && url!='' && user!=''">
    <v-btn @click="joinRoom()">加入房间</v-btn>
    </div>
    </v-expand-transition>
    <v-expand-transition>
    <div v-show="method!='x' && url!='' && user!=''">
    <v-btn @click="syncVideo()" >手动同步</v-btn>
    </div>
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
    <div v-show="start"
      ><video id="myVideo" class="video-js" >
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
      currentTime: 0,
    };
  },
  mounted() {
    this.room = this.$route.params.roomName;
    this.timer = setInterval(() => {
      this.syncVideo();
    }, 10000);
    socket.on("allUsers",(data)=>{
      this.userList = data.split(",").slice(0, -1);
      if(this.method == "x"){
        this.method="join";
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
      myPlayer.currentTime(parseFloat(data));
    })
    socket.on("join", (data)=>{
      console.log("joined",data)
      this.getTime();
      this.getUsers();
    })
    socket.on("sync", (data) => {
      let myTime= JSON.stringify(myPlayer.currentTime())
      console.log("data:",data)
      console.log("myTime:",myTime)
      if (parseFloat(myTime)<parseFloat(data.split(":")[1]) && Math.abs(parseFloat(myTime)-parseFloat(data.split(":")[1]))>5) {
        console.log("change")
        myPlayer.currentTime(parseFloat(data.split(":")[1]));
      }
    });
    setTimeout(()=>{
      this.getUsers();
      this.getUrl();
    }, 3000)
  },
  methods: {
    initVideo() {
      myPlayer = Video(myVideo, {
        controls: true,
        preload: true,
        width: "640px",
        // height: "200px",
      });
      console.log("stop");
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
        if(tmpTime - this.currentTime > 2 || tmpTime - this.currentTime < -2){
          console.log("xxxxxxxxxxxxxxxx")
          socket.emit('setTime',`${this.room}:${this.user}:`+JSON.stringify(myPlayer.currentTime()))
        }
        this.currentTime =myPlayer.currentTime()
      })
      socket.on("getTime", (data) => {
        console.log("getTime: " + data)
        socket.emit("time",`${this.room}:${this.user}:`+JSON.stringify(myPlayer.currentTime()));
      })
      this.getTime();
    },
    join() {
      socket.emit("join",`${this.room}:${this.user}`);
    },
    syncVideo() {
      console.log("syncVideo");
      socket.emit("getTime", `${this.room}:${this.user}`);
    },
    getTime() {
      console.log("videoTimeis:", myPlayer.currentTime());
      socket.emit("time",`${this.room}:${this.user}:`+JSON.stringify(myPlayer.currentTime()));
    },
    getUsers() {
      socket.emit("getUsers",`${this.room}:${this.user}`)
    },
    getUrl() {
      console.log("getUrl");
      socket.emit("getUrl",`${this.room}`)
    },
    setUrl() {
      socket.emit("setUrl",`${this.room}:::${this.url}`)
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