<template>
    <div class="min-h-screen bg-background flex items-center justify-center">
        <div class="max-w-4xl mx-auto text-white p-8">
            <h1 class="text-5xl font-bold mb-6">
                Reset Password
            </h1>
            <p class="text-lg mb-8">
                Enter your email address and we will send you a link to reset your password.
            </p>
            <form @submit.prevent="handleReset" class="flex space-x-4 mb-6">
                <input
                    type="email"
                    placeholder="Email"
                    class="px-4 py-3 rounded-lg text-cut-background w-full"
                    v-model="body.email"
                    required
                />
                <button
                    type="submit"
                    class="bg-red hover:bg-rose-500 text-white px-6 py-3 rounded-lg font-semibold"
                >
                    Reset Password
                </button>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { ResetRequest } from '~/models/auth.models';
import service from '~/service';

const loading = ref(false)
const body = ref<ResetRequest>(
   {
   email: '',
redirect_url:'http://localhost:3000/reset-change'
   } 
)

const handleReset = async () => {
    loading.value = true
 await service.auth.resetpassword(body.value)
    .then((resp: any) => {
        if (resp.status.code == 200) {
            console.log('Reset password success')
            swalToast({
                icon: 'success',
                title: 'Reset password success'
            })
        }
    })
    .catch((err: any) => {
        console.log(err)
    })
    .finally(() => {
        loading.value = false
    })
}

definePageMeta({
    layout: 'call',
});
</script>