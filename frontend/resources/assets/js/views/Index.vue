<template>
<div class="container">
    <div class="row">
        <div class="col-xs-12 col-sm-8 col-md-6 col-sm-offset-2 col-md-offset-3">
            <h2 class="text-center">Hi.</h2>
            <p class="text-justify">
                I’m Leo, a developer from Stockholm, Sweden. I am currently studying for my bachelor’s degree in computer sciences at the University of Uppsala, in Uppsala, Sweden. I am primarily a web developer and have fairly extensive experience with PHP, as well as several frontend tools. However, I am always interested in learning new concepts and new programming languages, so I know a few other things too...
            </p>
        </div>
    </div>
    <div class="row">
        <div class="col-xs-12 col-sm-8 col-md-6 col-sm-offset-2 col-md-offset-3">
        <h2 class="text-center">Stuff I know</h2>
            <ul class="nav nav-tabs nav-justified">
                <li role="presentation" class="active"><a href="#">Web stuff</a></li>
                <li role="presentation"><a href="#">Other stuff</a></li>
            </ul>
        </div>
    </div>
</div>
</template>

<script>
    export default {
        data() {
            return {
                users: {},
                loading: true
            }
        },
        asyncData() {
            return this.$http.get('http://localhost:8080/api', (data) => {
                return {users:data}
            })
        }
    }
</script>