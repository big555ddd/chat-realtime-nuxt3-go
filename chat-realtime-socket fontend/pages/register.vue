<template>
  <div class="min-h-screen flex items-center justify-center bg-background">
    <div class="bg-cut-background p-8 rounded-lg shadow-lg w-full max-w-md">
      <h2 class="text-3xl font-extrabold mb-6 text-center text-light">Register</h2>
      <form @submit.prevent="handleRegister">
        <!-- Username -->
        <div class="mb-4">
          <label for="username" class="block text-sm font-medium text-light">Username</label>
          <div class="relative mt-1">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </span>
            <input type="text" id="username" v-model="username"
              class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
              placeholder="Enter your username" required />
          </div>
        </div>

        <!-- Email -->
        <div class="mb-4">
          <label for="email" class="block text-sm font-medium text-light">Email</label>
          <div class="relative mt-1">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7 7 7-7M4 4h16v16H4V4z" />
              </svg>
            </span>
            <input type="email" id="email" v-model="email"
              class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
              placeholder="Enter your email" required />
          </div>
        </div>

        <!-- Password -->
        <div class="mb-4">
          <label for="password" class="block text-sm font-medium text-light">Password</label>
          <div class="relative mt-1">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M12 3C8.13 3 5 6.13 5 9v6c0 2.87 3.13 6 7 6s7-3.13 7-6V9c0-2.87-3.13-6-7-6zm-1 8H8v2h3v3h2v-3h3V9h-3V7h-2v2z" />
              </svg>
            </span>
            <input type="password" id="password" v-model="password"
              class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
              placeholder="Enter your password" required />
          </div>
        </div>

        <!-- First Name -->
        <div class="mb-4">
          <label for="first_name" class="block text-sm font-medium text-light">First Name</label>
          <div class="relative mt-1">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M12 2C6.48 2 2 6.48 2 12c0 5.52 4.48 10 10 10s10-4.48 10-10c0-5.52-4.48-10-10-10zm0 14c-2.21 0-4-1.79-4-4s1.79-4 4-4 4 1.79 4 4-1.79 4-4 4z" />
              </svg>
            </span>
            <input type="text" id="first_name" v-model="firstName"
              class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
              placeholder="Enter your first name" required />
          </div>
        </div>

        <!-- Last Name -->
        <div class="mb-4">
          <label for="last_name" class="block text-sm font-medium text-light">Last Name</label>
          <div class="relative mt-1">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6L9 18l-5-5" />
              </svg>
            </span>
            <input type="text" id="last_name" v-model="lastName"
              class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
              placeholder="Enter your last name" required />
          </div>
        </div>

        <!-- Display Name -->
        <div class="mb-6">
          <label for="display_name" class="block text-sm font-medium text-light">Display Name</label>
          <div class="relative mt-1">
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12l-8-8-8 8" />
              </svg>
            </span>
            <input type="text" id="display_name" v-model="displayName"
              class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
              placeholder="Enter your display name" required />
          </div>
        </div>

        <!-- Submit Button -->
        <button type="submit"
          :class="loading ? 'opacity-50 cursor-not-allowed' : 'bg-red  hover:bg-rose-500 text-white py-2 rounded-lg w-full'"
          :disabled="loading" class="transition duration-300 flex justify-center items-center font-semibold">
          <span v-if="!loading">Register</span>
          <span v-else>
            <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none"
              viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 000 8H4z"></path>
            </svg>
          </span>
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

import service from "~/service";


const username = ref('')
const email = ref('')
const password = ref('')
const firstName = ref('')
const lastName = ref('')
const displayName = ref('')
const router = useRouter()
const route = useRoute()

onMounted(() => {
  // รับค่า email จาก query parameter ถ้ามีอยู่
  if (route.query.email && typeof route.query.email === 'string') {
    email.value = route.query.email
  }
})
const loading = ref(false)

const handleRegister = async () => {
  loading.value = true
  await service.auth
    .register({
      username: username.value,
      email: email.value,
      password: password.value,
      first_name: firstName.value,
      last_name: lastName.value,
      display_name: displayName.value,
    })
    .then(async (resp: any) => {
      if (resp.status.code == '200') {
        router.push('/login')
      }
    })
    .catch((err: any) => {
      errorResp(err.response)
    })
    .finally(() => {
      loading.value = false
    })
}

definePageMeta({
    layout: 'call',
});
</script>

<style>
/* Add any additional custom styles here */
</style>