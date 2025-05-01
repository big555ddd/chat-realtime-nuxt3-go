<template>
    <div class="w-full h-10 bg-background"></div> <!-- Red line -->
    <div class="flex h-screen bg-background">
        <!-- Left Panel: Contact List -->
        <div class="w-1/4 bg-gray-800 border-r ml-10">
            <div class="p-4 text-lg font-semibold bg-red text-white">Contacts</div>
            <div class="space-y-4 p-4">
                <div v-for="contact in contacts" :key="contact.id" @click="selectContact(contact)"
                    class="flex items-center space-x-3 p-3 cursor-pointer hover:bg-gray-700 rounded-lg">
                    <img v-if="contact.avatar != ''" :src="contact.avatar" alt="Contact Avatar"
                        class="w-10 h-10 rounded-full" />
                    <svg v-else class="h-10 w-10 text-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <div>
                        <p class="text-sm font-medium text-white">{{ contact.name }}</p>
                        <p class="text-xs text-gray-400">{{ contact.status }}</p>
                    </div>
                </div>
            </div>
        </div>

        <!-- Right Panel: Chat Window -->
        <div class="flex-1 flex flex-col mr-10">
            <!-- Chat Header -->
            <div class="flex items-center justify-between p-4 bg-red text-white ">
                <div class="flex items-center space-x-2">
                    <!-- Conditionally render contact information -->
                    <template v-if="selectedContact">
                        <svg v-if="selectedContact.avatar === ''" class="h-7 w-7 text-light" fill="none"
                            viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <img v-else :src="selectedContact.avatar" alt="User Avatar" class="rounded-full w-7 h-7" />
                        <span class="font-semibold text-lg">{{ selectedContact.name }}</span>
                    </template>

                    <span v-else class="text-lg font-semibold">Select a contact</span>
                </div>

                <button @click="handleCloseChat" class="text-white hover:text-gray-300">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>

            <!-- Chat Messages -->
            <div class="flex-1 p-4 bg-gray-800">
                <div class="space-y-4">
                    <div v-for="(message, index) in selectedContact?.messages" :key="index" class="flex space-x-3">
                        <!-- Conditionally render message details -->
                        <template v-if="message.sent_by === 'contact'">
                            <svg v-if="selectedContact.avatar === ''" class="h-7 w-7 text-light" fill="none"
                                viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                            </svg>
                            <img v-else :src="selectedContact.avatar" alt="User Avatar" class="rounded-full w-7 h-7" />
                            <div class="bg-gray-700 p-3 rounded-lg shadow-md max-w-xs">
                                <p class="text-sm text-white">{{ message.text }}</p>
                                <span class="text-xs text-gray-400">{{ message.time }}</span>
                            </div>
                        </template>
                        <template v-else>
                            <div class="flex justify-end w-full">
                                <div class="bg-red p-3 rounded-lg shadow-md max-w-xs text-right">
                                    <p class="text-sm text-white">{{ message.text }}</p>
                                    <span class="text-xs text-gray-400">{{ message.time }}</span>
                                </div>
                            </div>
                        </template>
                    </div>
                </div>
            </div>


            <!-- Chat Input -->
            <div class="flex items-center p-4 bg-gray-800 border-t">
                <input v-model="messageInput" type="text" @keyup.enter="sendMessage"
                    class="flex-1 p-3 border rounded-lg text-sm text-white bg-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-600"
                    placeholder="Type a message..." />
                <button @click="sendMessage" class="ml-3 p-3 bg-red hover:bg-rose-500 text-white rounded-lg">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                </button>
            </div>
        </div>
    </div>
    <div class="w-full h-10 bg-background"></div> <!-- Red line -->
</template>

<script setup>

const contacts = ref([
    {
        id: 1,
        name: 'John Doe',
        avatar: 'https://via.placeholder.com/40',
        status: 'Online',
        messages: [
            { text: 'Hey, how are you?', time: '10:15', sent_by: 'contact' }
        ],
    },
    {
        id: 2,
        name: 'Jane Smith',
        avatar: '',
        status: 'Offline',
        messages: [
            { text: 'See you soon!', time: '9:00', sent_by: 'contact' }
        ],
    }
]);

const selectedContact = ref(null);  // Initially, no contact is selected
const messageInput = ref('');

function selectContact(contact) {
    selectedContact.value = contact;
}

function handleCloseChat() {
    selectedContact.value = null; // Close chat
}

function sendMessage() {
    if (messageInput.value.trim() && selectedContact.value) {
        const newMessage = {
            text: messageInput.value,
            time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
            sent_by: 'me'
        };

        selectedContact.value.messages.push(newMessage);  // Add the new message to the selected contact's messages

        messageInput.value = ''; // Clear message input after sending
    }
}
</script>

<style scoped>
/* You can customize additional styles here */
</style>
