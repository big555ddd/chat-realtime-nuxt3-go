<template>
    <div class="flex items-center p-4 bg-gray-800 border-t border-gray-600 mt-6">    
        <input
            v-model="message"
            type="text"
            @keyup.enter="sendMessage"
            class="flex-1 p-3 h-12 border rounded-lg text-sm text-white bg-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-600"
            placeholder="Type a message..."
        />
        <button @click="sendMessage" class="ml-3 p-3 bg-red hover:bg-rose-500 text-white rounded-lg">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="w-6 h-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
        </button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps({
    userID:{
      type: String,
    },
});
const message = ref<string>('');
const emit = defineEmits(['sendMessage']);

function sendMessage() {
    if (message.value.trim()) {
        emit('sendMessage', message.value.trim(),props.userID);
        message.value = '';
    }
}
</script>
