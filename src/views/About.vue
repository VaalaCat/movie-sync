<template>
  <div class="about">
    <Videojs
      :videosrc="'https://static.chive.vaa.la/html/video/式守同學不只可愛而已S1E1.mp4'"
    />
  </div>
</template>
<script>
import Videojs from "../components/Videojs.vue";
export default {
  name: "About",
  components: { Videojs },
};
</script>