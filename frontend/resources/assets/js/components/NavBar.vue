<template>
<div>
    <div class="top-bar col-xs-12" id="main-top">
        <h1 class="text-center text-white">{{ title }}</h1>
    </div>
    <img src="img/profile.jpg" alt="Profile picture" class="img-circle" id="profile-pic">
</div>
</template>

<script>
    export default {
        props: ['title']
    }
</script>