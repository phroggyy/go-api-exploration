import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import VueAsyncData from 'vue-async-data'

window.$ = window.jQuery = require('jquery');

Vue.use(VueResource)
Vue.use(VueRouter)
Vue.use(VueAsyncData)
Vue.config.debug = true

const router = new VueRouter({
    hashbang: false
});

import App from './App.vue'
import Index from './views/Index.vue'

router.map({
    '/': { component: Index }
})

router.redirect({
    '*': '/'
})

router.start(App, '#app')