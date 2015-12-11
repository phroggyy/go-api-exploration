<template>
<div class="row">
    <div class="top-bar" style="width: 100%">
        <img src="img/topbar.jpg" style="min-width: 100%; background-position: 50% 50%; min-height: 250px">
        <h1 class="text-center text-white" style="margin-top: -20rem;">Leo Sjöberg</h1>
        <div class="col-xs-6 col-sm-4 col-md-2 col-xs-offset-3 col-sm-offset-4 col-md-offset-5">
            <img src="img/profile.jpg" alt="Profile picture" class="img-circle col-xs-12">
        </div>
    </div>
</div>
    
    <!-- <nav class="navbar navbar-default" role="navigation">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-ex1-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">{{ title }}</a>
        </div>
    
        <div class="collapse navbar-collapse navbar-ex1-collapse">
            <ul class="nav navbar-nav">
                <li class="active"><a href="#">Link</a></li>
                <li><a href="#">Link</a></li>
            </ul>
            <form class="navbar-form navbar-left" role="search">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Search">
                </div>
                <button type="submit" class="btn btn-default">Submit</button>
            </form>
            <ul class="nav navbar-nav navbar-right">
                <li><a href="#">Link</a></li>
                <li class="dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">Dropdown <b class="caret"></b></a>
                    <ul class="dropdown-menu">
                        <li><a href="#">Action</a></li>
                        <li><a href="#">Another action</a></li>
                        <li><a href="#">Something else here</a></li>
                        <li><a href="#">Separated link</a></li>
                    </ul>
                </li>
            </ul>
        </div>
    </nav> -->
</template>

<script>
    export default {
        props: ['title']
    }
</script>