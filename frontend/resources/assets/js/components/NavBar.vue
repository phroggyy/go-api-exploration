<template>
<div>
    <div class="top-bar col-xs-12" id="main-top">
        <h1 class="text-center text-white">{{ title }}</h1>
        <img src="img/profile.jpg" alt="Profile picture" class="col-xs-6 col-sm-4 col-md-2 col-xs-offset-3 col-sm-offset-4 col-md-offset-5 img-circle" id="profile-pic">
    </div>
</div>
</template>

<script>
    export default {
        props: ['title'],
        ready() {
            var img = $('#profile-pic')
            img.css("position", "absolute")
            // For some reason, it appears the image isn't properly
            // loaded before the script runs, so we just delay it
            // by 20 millisecs.
            window.setTimeout(() => {
                img.css("top", $('#main-top').height() - img.height() / 2)
            }, 20)
        }
    }
</script>