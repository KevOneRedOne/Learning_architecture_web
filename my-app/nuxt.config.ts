// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
      head: {
        charset: 'utf-8',
        viewport: 'width=device-width, initial-scale=1',
        title: 'My App',
      }
  },
  $development: {
    devtools: { enabled: true },
  },
  css: ['~/assets/css/style.css'],

})
