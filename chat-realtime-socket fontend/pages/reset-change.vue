<template>
      <div class="min-h-screen bg-background flex items-center justify-center">
        <div class="max-w-4xl mx-auto text-white p-8">
            <h1 class="text-5xl font-bold mb-6">
                Reset Password
            </h1>
            <p class="text-lg mb-8">
                Enter your new password.
            </p>
            <form @submit.prevent="handleChange" class="flex space-x-4 mb-6">
                <input
                    type="password"
                    placeholder="password"
                    class="px-4 py-3 rounded-lg text-cut-background w-full"
                    v-model="body.password"
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
import type { ChangePasswordResetRequest } from '~/models/auth.models';
import service from '~/service';


const route = useRoute()

const token = route.query.token 

const body = ref<ChangePasswordResetRequest>({
  password: '',
    token: token as string
})

definePageMeta({
    layout: 'call',
});

const handleChange = async () => {
    await service.auth.changepassword(body.value)
        .then((resp: any) => {
            if (resp.status.code == 200) {
                console.log('Change password success')
                swalToast({
                    icon: 'success',
                    title: 'Change password success'
                })
            }
        })
        .catch((err: any) => {
            console.log(err)
        })
}

</script>