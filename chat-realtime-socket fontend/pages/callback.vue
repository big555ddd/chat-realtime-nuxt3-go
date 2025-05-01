<template>
  <div>
    <p>Redirecting...</p>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCookies } from '@vueuse/integrations/useCookies'

const router = useRouter()
const cookies = useCookies()

onMounted(() => {
  const urlParams = new URLSearchParams(window.location.search)
  const token = urlParams.get('data')

  if (token) {
    cookies.set('auth_token', token, { path: '/', secure: true, sameSite: 'Strict' })
    router.push('/')
  } else {
    const existingToken = cookies.get('auth_token')
    if (existingToken) {
      router.push('/')
    } else {
      console.error('Token not found in the URL or cookies')
      router.push('/login')
    }
  }
})
</script>

<style>
/* Add any additional custom styles here */
</style>