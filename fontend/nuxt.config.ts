// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: ['@pinia/nuxt', '@nuxtjs/tailwindcss'],
  plugins: ['~/plugins/websocket.ts'],
  runtimeConfig: {
    public: {
      NUXT_PUBLIC_BASE_URL: process.env.NUXT_PUBLIC_BASE_URL || 'http://localhost:8080',
      NUXT_PUBLIC_SOCKET_URL: process.env.NUXT_PUBLIC_SOCKET_URL || 'localhost:8080',
      NUXT_PUBLIC_CALLBACK_URL: process.env.NUXT_PUBLIC_CALLBACK_URL || 'http://localhost:8080/callback',
      NUXT_PUBLIC_FIREBASE_API_KEY: process.env.NUXT_PUBLIC_FIREBASE_API_KEY || '',
      NUXT_PUBLIC_FIREBASE_AUTH_DOMAIN: process.env.NUXT_PUBLIC_FIREBASE_AUTH_DOMAIN || '',
      NUXT_PUBLIC_FIREBASE_DATABASE_URL: process.env.NUXT_PUBLIC_FIREBASE_DATABASE_URL || '',
      NUXT_PUBLIC_FIREBASE_PROJECT_ID: process.env.NUXT_PUBLIC_FIREBASE_PROJECT_ID || '',
      NUXT_PUBLIC_FIREBASE_STORAGE_BUCKET: process.env.NUXT_PUBLIC_FIREBASE_STORAGE_BUCKET || '',
      NUXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID: process.env.NUXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID || '',
      NUXT_PUBLIC_FIREBASE_APP_ID: process.env.NUXT_PUBLIC_FIREBASE_APP_ID || '',
      NUXT_PUBLIC_FIREBASE_MEASUREMENT_ID: process.env.NUXT_PUBLIC_FIREBASE_MEASUREMENT_ID || '',
    },
  },
})