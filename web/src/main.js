import { createApp } from 'vue'

import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from '@/App.vue'
import router from '@/router'

import event from '@/js/pubsub'
import utils from '@/js/utils'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import '@/css/layout.less'
import '@/css/icon.less'
import '@/css/moment.less'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate);

app.config.globalProperties = {
    $event: event,
    $utils: utils
}

app.provide('event', event);
app.provide('utils', utils);

app.use(pinia).
use(router).
use(ElementPlus).
mount('#app')

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
  