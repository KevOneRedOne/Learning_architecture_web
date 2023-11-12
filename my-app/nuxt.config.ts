// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      title: 'My App',
    }
  },
  css: ['~/assets/css/style.css'],
  $development: {
    devtools: { enabled: true },
  },
  $production: {
    devtools: { enabled: false },
  },

  runtimeConfig: {

    public: {
      baseAPIGo: process.env.NUXT_API_PATH_GO,
      baseAPIFass: process.env.NUXT_API_PATH_FASS,
    }
    
    
  },
})



