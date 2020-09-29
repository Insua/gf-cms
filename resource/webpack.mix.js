let mix = require('laravel-mix')

mix.setPublicPath('../public/')

mix.js('js/app.js', 'js/app.js').sass('scss/app.scss', 'css/app.css')

if (mix.inProduction()) {
  mix.version().disableNotifications()
} else {
  mix.sourceMaps()
}
