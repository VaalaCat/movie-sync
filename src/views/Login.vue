<template>
  <v-text-field v-model="roomName" label="房间名"></v-text-field>
  <v-btn>加入房间</v-btn>
  {{roomName}}
</template>

<script>
export default {
  data() {
    return {
      roomName: ""
    }
  }
};
</script>

<style>
</style>