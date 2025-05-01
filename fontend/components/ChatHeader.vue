<template>
    <div class="flex items-center justify-between p-4 bg-red text-white">
        <div class="flex items-center space-x-2">
            <template v-if="selectedContact">
                <img
                    v-if="selectedContact.avatar"
                    :src="selectedContact.avatar"
                    alt="User Avatar"
                    class="rounded-full w-7 h-7"
                />
                <svg
                    v-else
                    class="h-8 w-7 text-light"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                </svg>
                <span class="font-semibold text-lg">{{ selectedContact.display_name }}</span>
            </template>
            <span v-else class="text-lg font-semibold">Select a contact</span>
        </div>

        <button @click="$emit('closeChat')" class="text-white hover:text-gray-300">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="w-6 h-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
        </button>
    </div>
</template>

<script setup lang="ts">
import type { Contact } from '~/models/chat.models';

const props = defineProps({
    selectedContact: Object as () => Contact | null,
});
const emit = defineEmits(['closeChat']);
</script>
