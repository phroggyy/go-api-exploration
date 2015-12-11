var elixir = require('laravel-elixir');

require('laravel-elixir-vueify');

elixir(function (mix) {
    mix.browserify("app.js")
       .scripts([
            "bootstrap-sass/assets/javascripts/bootstrap.js"
        ],"public/js/vendor.js", "node_modules")
       .sass([
            "app.scss"
        ], "public/css/app.css")
       .copy("resources/assets/img", "public/img")
});