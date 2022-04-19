<template>
  <v-text-field v-model="roomName" label="房间名" required :rules="roomRules" type="text"></v-text-field>
  <v-expand-transition>
    <v-alert v-show="showError" color="red">
      房间名不可为空
    </v-alert>
  </v-expand-transition>
  <v-btn @click="join()">加入房间</v-btn>
</template>

<script>
export default {
  data() {
    return {
      roomName: "",
      showError: false,
      roomRules: [
        (v) => !!v || "房间名不可为空",
      ],
    }
  },
  methods: {
    mounted() {
      this.roomName = this.$route.params.roomName;
    },
    join: function() {
      if (this.roomName!="") {
          this.$router.push({
            path: `/room/${this.roomName}`,
        })
      }else{
        this.showError=true;
      }
    }
  }
};
</script>

<style>
</style>