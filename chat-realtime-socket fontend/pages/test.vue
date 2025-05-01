<template>
    <div>
    <pre>{{ org }}</pre>
    <pre>asfasfsafa</pre>
    <pre>{{ loading }}</pre>
    </div>
</template>

<script setup lang="ts">
import type { List } from '~/models/org.models'
import service from '~/service'

const loading = ref(false)

const org = ref<List[]>([])

const listorg = async () => {
    loading.value = true
    await service.org
        .list()
        .then((resp: any) => {
                console.log(resp)
                org.value = resp
     
        })
        .catch((err: any) => {
            console.log(err)
        })
        .finally(() => {
            loading.value = false
        })
}

onMounted(async () => {
    await listorg()
})

definePageMeta({
    layout: 'call',
});
</script>