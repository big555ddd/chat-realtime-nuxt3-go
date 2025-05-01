<template>
    <NuxtPage />
</template>

<script setup lang="ts">
import { useIndexStore } from '~/store/main'
import service from '~/service';

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
const store = useIndexStore()

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

const getinfo = () => {
    const data = store.$state

    service.user
        .getUser()
        .then((resp: any) => {
            if (resp.status.code == 200) {
                user.value = resp.data
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
}

onMounted(() => {
    getinfo()
})


</script>