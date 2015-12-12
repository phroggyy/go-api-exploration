<template>
<div>
    <nav-bar :title="title"></nav-bar>
    <router-view></router-view>
</div>
</template>

<script>
    var NavBar = require('./components/NavBar.vue');

    export default {
        data() {
            return {
                title: "Leo Sjöberg"
            }
        },
        components: {
            'nav-bar': NavBar
        }
    }
</script>