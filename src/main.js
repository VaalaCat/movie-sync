import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'

require('video.js/dist/video-js.css')

loadFonts()

createApp(App)
  .use(router)
  .use(vuetify)
  .mount('#app')
