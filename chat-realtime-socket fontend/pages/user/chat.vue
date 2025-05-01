<template>
    <div class="flex h-screen bg-background">
        <!-- Sidebar for contact list -->
        <div class="w-1/4 bg-gray-800 border-r border-gray-600">
            <ContactList 
                @selectContact="selectContact" 
                v-bind="contactProps" 
                @fakeclick = "handlefakeclick"
            />
        </div>
        <!-- Chat window for selected contact -->
        <ChatWindow 
            v-if="selectedContact" 
            :selectedContact="selectedContact"
            :isOpen="isOpen"
            @closeChat="handleCloseChat" 
            @textnoti="handleNoti" 
        />
        <!-- Placeholder when no contact is selected -->
        <div v-else class="flex-1 flex items-center justify-center text-white">
            <p>Select a contact to start chatting</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import ContactList from '~/components/ContactList.vue';
import ChatWindow from '~/components/ChatWindow.vue';
import type { Contact } from '~/models/chat.models';
import { useCookies } from '@vueuse/integrations/useCookies';

// Selected contact state
const cookies = useCookies();
const refid = cookies.get('id');
const selectedContact = ref<Contact | null>(null);
const total = ref<number>(0);
const id = ref<string>("");
const type = ref<string>("");
const detail = ref<string>("");
const time = ref<number>(0);
const isOpen = ref<boolean>(false);


// Computed props for ContactList
const contactProps = computed(() => {
    const props: { total?: number; idcon?: string, type?: string, detail?:string,time?:number } = {};
    
    if (total.value !== 0) props.total = total.value;
    if (id.value !== "") props.idcon = id.value;
    if (type.value!== "") props.type = type.value;
    if (detail.value!== "") props.detail = detail.value;
    if (time.value!== 0) props.time = time.value;
    return props;
});

function handlefakeclick(contact: Contact){
    selectedContact.value = contact;
    isOpen.value = false;    
}

// Select a contact
function selectContact(contact: Contact) {
    selectedContact.value = contact;
    isOpen.value = true; 
}

// Close chat window
function handleCloseChat() {
    selectedContact.value = null;
}

let currentContactId: string | null = null;
const websocket = ref<WebSocket | null>(null);


function handleNoti(newTotal: number, newId: string,messageType?: string,messageDetails?: string,MessageTime?:number) {
    total.value = newTotal;
    id.value = newId;
    if (messageType && messageDetails && MessageTime){
        type.value = messageType;
        detail.value = messageDetails;
        time.value = MessageTime;
    }
}

definePageMeta({
    layout: 'chat',
});

</script>

<style></style>
