<template v-if="!store.loading">
    <div class="w-full h-24 flex justify-center bg-cut-background">
        <div class="flex justify-between w-full max-w-screen-xl h-24 font-semibold text-2xl">
            <div class="flex space-x-20 items-center text-light">
                <div @click="handleHome" class="cursor-pointer">Home</div>
                <div @click="handleDemo" class="cursor-pointer">Demo</div>
                <!-- <div @click="handleAbout"class="cursor-pointer">About</div> -->
                <div @click="router.push('/contact')" class="cursor-pointer">Contact</div>
            </div>  
            <div class="flex space-x-12 items-center relative text-light">
                <div v-if="!token" @click="handleLogin" class="cursor-pointer">Login</div>
                <div v-if="!token" @click="handleRegister"
                    class="cursor-pointer bg-red hover:bg-rose-500 h-full w-48 text-center flex items-center justify-center rounded-md">
                    Be A Member?
                </div>
                <div v-else class="text-light relative">
                    <div @click="toggleDropdown" class="cursor-pointer">
                        Welcome, {{ store.display_name }}
                    </div>
                    <div v-if="showDropdown" class="absolute right-0 mt-2 w-48 bg-white border rounded-lg shadow-lg">
                        <div @click="handleProfile" class="px-4 py-2 hover:bg-gray-100 cursor-pointer">
                            My Profile
                        </div>
                        <div @click="handleLogout" class="px-4 py-2 hover:bg-gray-100 cursor-pointer">
                            Logout
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <NuxtPage class=" bg-background" />



<footer class="w-full h-16 flex justify-center items-center bg-cut-background text-light text-center">
    <p class="text-sm font-medium">
        CHATBEAT © 2024 All rights reserved.
    </p>
</footer>
</template>


<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCookies } from '@vueuse/integrations/useCookies'
import service from '~/service';
import { useIndexStore } from '~/store/main'


const router = useRouter()
const cookies = useCookies()
const showDropdown = ref(false)
const token = cookies.get('token')
const store = useIndexStore()

const handleHome = () => {
    router.push('/')
}

interface User {
    id : string,
    username : string,
    email : string,
    password : string,
    first_name : string,
    last_name : string,
    display_name : string,
    role_id : number,
    status : string,
    created_at : number,
    updated_at : number,
}

const user = ref<User >(
    {
        id : '',
        username : '',
        email : '',
        password : '',
        first_name : '',
        last_name : '',
        display_name : '',
        role_id : 0,
        status : '',
        created_at : 0,
        updated_at : 0,
    }
)

const handleLogin = () => {
    router.push('/login')
}
const handleDemo = () => {
    router.push('/demo')
}

const handleRegister = () => {
    router.push('/register')
}

const toggleDropdown = () => {
    showDropdown.value = !showDropdown.value
}

const handleProfile = () => {
    // เปลี่ยนเส้นทางไปที่หน้าโปรไฟล์ผู้ใช้
    router.push('user/profile')
}

const handleLogout = () => {
    // ลบคุกกี้ token และ redirect_url
    cookies.remove('token')

    // เปลี่ยนเส้นทางไปที่ /
    router.push('/')
    //ถ้าอยู่/อยู่แล้วให้refreshหน้า
    location.reload()
}



onMounted(async () => {
    const data = store.$state
    if (data.id != "") {
        data.loading = true
        await service.user
        .getUser()
        .then((resp: any) => {
            if (resp.status.code == '200') {
                user.value = resp.data
                console.log(user.value);
                data.id = user.value.id
                data.username = user.value.username
                data.email = user.value.email
                data.first_name = user.value.first_name
                data.last_name = user.value.last_name
                data.display_name = user.value.display_name
                data.role_id = user.value.role_id
                data.status = user.value.status
                data.created_at = user.value.created_at
                data.updated_at = user.value.updated_at
            }

        })
        .catch((err: any) => {
            console.log(err)
        })
        .finally(() => {
            data.loading = false
        })
    }
})
</script>

<style>
/* Add any additional custom styles here */
</style>