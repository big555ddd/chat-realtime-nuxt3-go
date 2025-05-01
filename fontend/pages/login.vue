<template>
    <div class="min-h-screen flex items-center justify-center bg-background">
        <div class="bg-cut-background p-8 rounded-lg shadow-2xl w-full max-w-md">
            <h2 class="text-3xl font-extrabold mb-6 text-center text-light">Welcome Back</h2>
            <p class="text-sm text-center text-light mb-6">
                Please login to your account to continue
            </p>
            <form @submit.prevent="handleSubmit" class="space-y-6">
                <!-- Username -->
                <div>
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

                <!-- Password -->
                <div>
                    <label for="password" class="block text-sm font-medium text-light">Password</label>
                    <div class="relative mt-1">
                        <span class="absolute inset-y-0 left-0 flex items-center pl-3">
                            <svg class="h-5 w-5 text-gray-400" width="24" height="24" viewBox="0 0 24 24"
                                stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round"
                                stroke-linejoin="round">
                                <path stroke="none" d="M0 0h24v24H0z" />
                                <rect x="5" y="11" width="14" height="10" rx="2" />
                                <circle cx="12" cy="16" r="1" />
                                <path d="M8 11v-5a4 4 0 0 1 8 0" />
                            </svg>
                        </span>
                        <input type="password" id="password" v-model="password"
                            class="w-full pl-10 pr-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-sm"
                            placeholder="Enter your password" required />
                    </div>
                </div>

                <!-- Submit Button -->
                <button type="submit" :class="loading ? 'opacity-50 cursor-not-allowed' : ''"
                    class="w-full bg-red text-white py-2 rounded-lg hover:bg-rose-500 transition duration-300 flex justify-center items-center font-semibold"
                    :disabled="loading">
                    <span v-if="!loading">Login</span>
                    <span v-else>
                        <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none"
                            viewBox="0 0 24 24">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4">
                            </circle>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4a4 4 0 000 8H4z"></path>
                        </svg>
                    </span>
                </button>
            </form>
            <button @click="router.push('/reset')" class="text-white">
                Reset Password ?
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import service from "~/service";

const username = ref("");
const password = ref("");
const router = useRouter();
const loading = ref(false);

const handleSubmit = async () => {
    loading.value = true;
    await service.auth
        .login({
            username: username.value,
            password: password.value,
        })
        .then(async (resp: any) => {
            if (resp.status.code == "200") {
                const refToken = setCookie("token");
                refToken.value = resp?.data || "";
                const refUsername = setCookie("username");
                refUsername.value = username.value || "";
                await service.user
                    .getUser()
                    .then((resp: any) => {
                        if (resp.status.code == '200') {
                            const id = setCookie("id");
                            id.value = resp.data.id
                        }

                    })
                    .catch((err: any) => {
                        console.log(err)
                    })
                router.push("/user/chat");
            }
        })
        .catch((err: any) => {
            console.log(err.response._data);
            
            errorResp(err.response);
        })
        .finally(() => {
            loading.value = false;
        });
};

definePageMeta({
    layout: 'call',
});
</script>

<style scoped>
/* Add any additional custom styles here */
</style>
