import { defineNuxtPlugin } from '#app'
import { useCookies } from 'vue3-cookies'

export default defineNuxtPlugin((nuxtApp) => {
  // สามารถเข้าถึง cookies โดยตรงจาก `useCookies`
  nuxtApp.provide('cookies', useCookies())
})
