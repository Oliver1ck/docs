import { createApp } from 'vue'

import './assets/styles/null.css'
import './assets/styles/fonts.css'
import './assets/styles/variables.css'
import App from './App.vue'
import router from './router'

createApp(App).use(router).mount('#app')
