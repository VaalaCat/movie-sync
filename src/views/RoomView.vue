<template>
  <div class="room">
    <Videojs/>
  </div>
</template>
<script>
import Videojs from "../components/Videojs.vue";
export default {
  name: "RoomView",
  components: { Videojs },
};
</script>