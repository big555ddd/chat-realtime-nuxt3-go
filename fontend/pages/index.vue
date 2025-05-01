<template>
  <div class="min-h-screen bg-background flex items-center justify-center">
    <div class="max-w-4xl mx-auto text-white p-8">
      <h1 class="text-5xl font-bold mb-6">
        A Beautiful Chat With Dark Theme Ui
      </h1>
      <p class="text-lg mb-8">
        Keep everything in the same place—even Where u wanna talk.
      </p>
      <form @submit.prevent="handleSignUp" class="flex space-x-4 mb-6">
        <input
          type="email"
          placeholder="Email"
          class="px-4 py-3 rounded-lg text-cut-background w-full"
          v-model="email"
          required
        />
        <button
          type="submit"
          class="bg-red hover:bg-rose-500 text-white px-6 py-3 rounded-lg font-semibold"
        >
          Sign up - it’s free!
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useIndexStore } from '~/store/main'


const email = ref('')
const router = useRouter()
const store = useIndexStore()


const handleSignUp = () => {
  // นำพาผู้ใช้ไปที่หน้า /register พร้อมกับส่ง email ผ่าน query parameter
  router.push({ path: '/register', query: { email: email.value } })

}


</script>

<style>
/* Add any additional custom styles here */
</style>
