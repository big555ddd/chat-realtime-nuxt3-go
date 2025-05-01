<template>
  <div v-if="show" class="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50">
    <div class="bg-gray-900 p-6 rounded-lg shadow-lg max-w-md w-full text-white relative">
      <button @click="closeModal" class="absolute top-2 right-2 text-gray-400 hover:text-gray-200">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
      <h2 class="text-2xl font-semibold mb-4">Add Friend</h2>
      <div class="flex items-center space-x-2 mb-4">
        <input type="text" v-model="search" @keyup.enter="fetchUserList" placeholder="Search users..."
          class="flex-grow p-2 rounded-lg bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <button @click="fetchUserList"
          class="bg-gray-700 text-white p-2 rounded-lg hover:bg-gray-600 transition duration-300">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M21 21l-4.35-4.35m1.35-5.65a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
        </button>
      </div>
      <div v-if="loading" class="text-center text-gray-500">Loading...</div>
      <div v-else>
        <div v-if="contacts && contacts.length">
          <ul>
            <li v-for="contact in contacts" :key="contact.id"
              class="flex items-center justify-between p-2 border-b border-gray-700">
              <div>
                <p class="text-sm font-medium">{{ contact.display_name }}</p>
                <p class="text-xs text-gray-400">{{ contact.status }}</p>
              </div>
              <div v-if="contact.id === id"
                class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition duration-300">
                You
              </div>
              <button v-else-if="!contactsFriend?.find((item) => item.id === contact.id)" @click="addFriend(contact)"
                class="bg-green-500 text-white px-4 py-2 rounded-lg hover:bg-green-600 transition duration-300">
                Add
              </button>
              <button v-else
                class="bg-gray-500 text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition duration-300">
                Added
              </button>
            </li>
          </ul>
          <!-- Pagination Controls -->
          <Pagination :currentPage="page" :totalPages="paginate.total_pages" @prevPage="prevPage"
            @nextPage="nextPage" />
        </div>
        <div v-else class="text-center text-gray-500">No users found</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Paginate, Param } from '~/models/http.models';
import service from '~/service';
import { useCookies } from '@vueuse/integrations/useCookies';
import type { AddFriend, Contact } from '~/models/chat.models';



const cookies = useCookies();
const id = cookies.get('id') || '';
const search = ref('');
const loading = ref(false);
const contacts = ref<Contact[] | null>(null);
const contactsFriend = ref<Contact[] | null>(null);
const page = ref(1);
const limit = ref(5);

const paginate = ref<Paginate>({
  current_page: 0,
  per_page: 0,
  total_pages: 0,
  total: 0
});

const props = defineProps({
  show: Boolean,
});

const emit = defineEmits(['close']);

const fetchUserList = () => {
  const param: Param = {
    limit: limit.value,
    page: page.value,
    search: search.value,
  };
  loading.value = true;
  service.user
    .listUser(param)
    .then((resp: any) => {
      if (resp.status.code == '200') {
        contacts.value = resp.data;
        paginate.value = resp.pagination;
      }
    })
    .catch((err: any) => {
      console.log(err);
    })
    .finally(() => {
      loading.value = false;
    });
};

const fetchUserListFriend = () => {
  const param: Param = {
    limit: 9999,
    page: 1,
    search: '',
  };
  loading.value = true;
  service.user
    .listfriend(id, param)
    .then((resp: any) => {
      if (resp.status.code == '200') {
        contactsFriend.value = resp.data;
      }
    })
    .catch((err: any) => {
      console.log(err);
    })
    .finally(() => {
      loading.value = false;
    });
};

const addFriend = (contact: Contact) => {
  const body: AddFriend = {
    user_id: contact.id,
  };
  service.chat
    .AddFriend(body)
    .then((resp: any) => {
      if (resp.status.code == 200) {
        console.log('Friend added successfully:', resp);
        fetchUserListFriend();
        fetchUserList();
      }
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const closeModal = () => {
  emit('close');
};

const nextPage = () => {
  if (page.value < paginate.value.total_pages) {
    page.value++;
    fetchUserList();
  }
};

const prevPage = () => {
  if (page.value > 1) {
    page.value--;
    fetchUserList();
  }
};

onMounted(async () => {
  await fetchUserListFriend();
  await fetchUserList();
});
</script>

<style scoped>
/* Add any additional styling here */
</style>