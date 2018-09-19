import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import { log } from 'util'

Vue.config.productionTip = false

new Vue({
  render: (h) => h(App),
}).$mount('#app')
