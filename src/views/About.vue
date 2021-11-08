<template>
  <div class="about">
    <Videojs
      :videosrc="'https://kol-fans.fp.ps.netease.com/file/616c4cc683c18c62adf6dc7eMDuQSBZl03?from=ysjdm.com|ysjdm.net'"
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