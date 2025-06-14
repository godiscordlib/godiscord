// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from '@tailwindcss/vite'
export default defineNuxtConfig({
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  app: {
    head: {
      title: "godiscord",
      meta: [
        {
          name: "google-site-verification",
          content: "w6I7NxoiHb1ooEJ-dbj-mrfLbjSdBkcHQuQvtX5XDME",
        }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon' ,href: '/favicon.ico' }
      ]
    }
  },
  site: {
    url: 'https://godiscord.foo.ng',
    name: 'godiscord',
    description: 'godiscord is a Go library to make Discord bots.',
    defaultLocale: 'en',
  },
  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/eslint',
    '@nuxt/content',
    '@nuxt/scripts',
    '@nuxtjs/seo',
    '@vueuse/nuxt',
  ],
})