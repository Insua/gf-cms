let mix = require('laravel-mix')
const tailwindcss = require('tailwindcss')

mix.setPublicPath('../public/')

mix.js('js/app.js', 'js/app.js')

let sassOption = mix.inProduction() ? {
    processCssUrls: false,
    postCss: [ tailwindcss('./tailwind.config.prod.js') ],
  } : {
  processCssUrls: false,
  postCss: [ tailwindcss('./tailwind.config.js') ],
}

mix.sass('scss/app.scss', 'css/app.css').options(sassOption)

if (mix.inProduction()) {
  mix.version().disableNotifications()
} else {
  mix.sourceMaps()
}
