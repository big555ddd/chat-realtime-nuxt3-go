// plugins/websocket.ts
import { useCookies } from '@vueuse/integrations/useCookies';

export const connectWebSocket = (contactId: string): WebSocket | null => {
  const cookies = useCookies();
  const token = cookies.get('token');

  if (!token) {
    console.error('No token found.');
    return null;
  }

  const config = useRuntimeConfig() as unknown as { public: { NUXT_PUBLIC_SOCKET_URL: string } };

  const baseUrl = config?.public?.NUXT_PUBLIC_SOCKET_URL ;

  // Determine the protocol (wss:// if the page is loaded over https://)
  const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws';

  // WebSocket URL with the contactId as a query parameter
  const socketUrl = `${protocol}://${baseUrl}/api/v1/sockets/ws/${contactId}?Authorization=${token}`;
  // console.log('Connecting to WebSocket:', socketUrl);

  // Creating a new WebSocket connection
  const websocket = new WebSocket(socketUrl);

  console.log('WebSocket connecting...');

  // Event handlers for WebSocket connection events
  websocket.onopen = () => {
    console.log('WebSocket connected');
  };

  websocket.onmessage = (event) => {
  };

  websocket.onerror = (error) => {
    console.error('WebSocket error: ', error);
  };

  websocket.onclose = () => {
    console.log('WebSocket disconnected');
  };

  return websocket;
};

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.provide('connectWebSocket', connectWebSocket);
});
