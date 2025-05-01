<template>
  <div v-if="isOpen" class="flex-1 flex flex-col">
    <!-- Chat header -->
    <ChatHeader :selectedContact="selectedContact" @closeChat="handleCloseChat" />
    <!-- Messages list -->
    <div ref="chatContainer"
      class="overflow-y-auto mini-scroll h-[calc(79vh-0rem)] max-sm:h-[calc(68vh-6rem)] 2xl:h-[calc(83vh-1rem)] flex flex-col-reverse">
      <div v-if="!loading" class="space-y-4 my-8 mx-4">
        <div v-for="(message, index) in messages || []" :key="index" class="flex"
          :class="message.sender === id ? 'justify-end' : 'justify-start'">
          <div class="px-4 py-2 rounded-lg"
            :class="message.sender === id ? 'bg-blue-500 text-white' : 'bg-gray-700 text-gray-300'">
            <p class="max-w-lg break-words whitespace-pre-wrap">{{ message.detail }}</p>
            <p class="text-xs mt-1">{{ formatUnixToDateTime(message.created_at) }}</p>
          </div>
        </div>
      </div>
    </div>
    <!-- Chat input -->
    <ChatInput :userID="selectedContact?.id" @sendMessage="sendMessage" />
  </div>
</template>

<script setup lang="ts">
import ChatHeader from './ChatHeader.vue';
import ChatInput from './ChatInput.vue';
import type { Contact, Message } from '~/models/chat.models';
import { useCookies } from '@vueuse/integrations/useCookies';
import { connectWebSocket } from '~/plugins/websocket';
import formatUnixToDateTime from '~/utils/time_unix';


const cookies = useCookies();
const id = cookies.get('id');
const loading = ref(false);
const chatContainerRef = ref<HTMLDivElement | null>(null);
const messages = ref<Message[]>([]);
const websocket = ref<WebSocket | null>(null);

const emit = defineEmits(['sendMessage', 'closeChat', 'textnoti']);
const props = defineProps({
  selectedContact: {
    type: Object as PropType<Contact | null>,
    required: false,
  },
  isOpen: {
    type: Boolean,
    required: false,
    default: false,
  },
});

const scrollToBottom = () => {
  const chatContainer = chatContainerRef.value;
  if (chatContainer != null) {
    chatContainer.scrollTop = chatContainer.scrollHeight;
  }
};


function handleCloseChat() {
  emit('closeChat');
}
function sendMessage(message: string, userID: string) {
  if (websocket.value && websocket.value.readyState === WebSocket.OPEN) {
    const outgoingMessage = {
      sender: id,
      recipient: userID,
      detail: message,
      type: 'text',
    };
    websocket.value.send(JSON.stringify(outgoingMessage));
    emit('textnoti', 0, userID)
    scrollToBottom();
  } else {
    console.error("WebSocket connection is not open.");
  }
}

let currentContactId: string | null = null;

const conwebsocket = () => {
  if (!props.selectedContact?.id) {
    console.error("No selected contact ID available.");
    return;
  }
  if (currentContactId !== props.selectedContact.id) {
    if (websocket.value) {
      websocket.value.close();
      websocket.value = null;
    }
    currentContactId = props.selectedContact.id;
    websocket.value = connectWebSocket(props.selectedContact.id);

    if (websocket.value) {
      websocket.value.addEventListener('message', (event) => {
        const incomingMessage = JSON.parse(event.data);
        // console.log('Received message:', incomingMessage);
        if (incomingMessage.total) {
          emit('textnoti', incomingMessage.total, incomingMessage.sender,incomingMessage.type,incomingMessage.detail, incomingMessage.created_at,)
        }
        if ((incomingMessage.recipient === id || incomingMessage.recipient === currentContactId) &&
          (incomingMessage.sender === id || incomingMessage.sender === currentContactId)) {
          messages.value.push(incomingMessage);
          scrollToBottom();
        }        // if(incomingMessage.recipient == props.selectedContact?.id)
        // {
        //   messages.value.push(incomingMessage);
        //   scrollToBottom();
        // }
      });

      websocket.value.addEventListener('close', () => {
        console.log('WebSocket connection closed');
      });

      websocket.value.addEventListener('error', (error) => {
        console.error('WebSocket error:', error);
      });
    } else {
      console.error("WebSocket connection failed.");
    }
  } else {
    console.log('WebSocket already connected for this contact.');
  }
};

watch(() => props.selectedContact?.id, async (newId, oldId) => {
  if (newId !== oldId) {
    messages.value = [];

    await conwebsocket();
    await scrollToBottom();
  }
});

onMounted(async () => {
  if (props.selectedContact?.id) {
    await conwebsocket();
  }
});


</script>
