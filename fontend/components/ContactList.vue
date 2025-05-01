    <template >
      <div v-if="!loading" class="space-y-4 flex flex-col h-full">
        <!-- Contact list -->
        <div class="flex-1">
          <div class="p-4 bg-gray-800 flex items-center space-x-2">
            <input type="text" v-model="search" @keyup.enter="searchContacts" placeholder="Search contacts..."
              class="flex-grow p-2 rounded-lg bg-gray-700 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
            <button @click="searchContacts"
              class="bg-gray-700 text-white p-2 rounded-lg hover:bg-gray-600 transition duration-300">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M21 21l-4.35-4.35m1.35-5.65a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </button>
            <button @click="addFriend()"
              class="bg-red text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition duration-300">
              Add Friend
            </button>
            <AddFriend :show="showAddFriendForm" @close="closeAddFriendModal" />
          </div>
          <div v-for="contact in contacts" :key="contact.id" @click="selectContact(contact)"
            class="flex items-center space-x-3 p-3 cursor-pointer bg-gray-700 hover:bg-gray-400 border-b border-gray-600 justify-between">
            <div class="flex gap-2">
              <img v-if="contact.avatar" :src="contact.avatar" alt="Contact Avatar" class="w-10 h-10 rounded-full" />
              <svg v-else class="h-10 w-10 text-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>

              <div v-if="contacts?.length != 0">
                <p class="text-sm font-medium text-light">{{ contact.display_name }}</p>
                <p v-if="contact.type == 'text'" class="text-xs text-gray-400">{{ contact.detail }} {{
                  formatUnixToDateTime(contact.message_time) }}</p>
                <p v-else class="text-xs text-gray-400">{{ contact.status }}</p>
              </div>
            </div>
            <span v-if="contact.unread > 0" class="bg-red text-white px-2 py-1 rounded-full text-xs">
              {{ contact.unread }}
            </span>
            <button @click.stop="removeFriend(contact)"
              class="bg-red text-white px-4 py-2 rounded-lg hover:bg-red-600 transition duration-300 flex items-center space-x-2">
              <svg class="h-8 w-8 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M13 7a4 4 0 11-8 0 4 4 0 018 0zM9 14a6 6 0 00-6 6v1h12v-1a6 6 0 00-6-6zM21 12h-6" />
              </svg>
            </button>
          </div>
        </div>
        <Pagination class="px-4 text-light" :currentPage="page" :totalPages="paginate.total_pages" @prevPage="prevPage"
          @nextPage="nextPage" />

        <!-- Bottom section -->
        <div v-if="!loading" class="flex justify-between items-center p-4 bg-gray-800 border-t border-gray-600">
          <div class="flex">
            <svg class="h-12 w-12 text-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <div class="ml-2">
              <p class="text-md font-medium text-light">{{ store.first_name + ' ' + store.last_name }}</p>
              <p class="text-sm text-gray-400">{{ store.display_name }}</p>
            </div>
          </div>
          <!-- Three-dot menu icon for dropdown -->
          <div class="relative">
            <svg @click="toggleDropdown" class="h-8 w-8 text-gray-500 cursor-pointer" viewBox="0 0 24 24" fill="none"
              stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="1" />
              <circle cx="12" cy="5" r="1" />
              <circle cx="12" cy="19" r="1" />
            </svg>
            <!-- Dropdown Menu -->
            <div v-if="showDropdown"
              class="absolute right-0 bottom-full mb-2 bg-background text-light border border-gray-600 rounded-lg shadow-lg w-48">
              <ul>
                <li @click="openProfileModal"
                  class="px-4 py-2 border border-gray-600 cursor-pointer hover:bg-cut-background">My
                  Profile</li>
                <li @click="navigateToSettings"
                  class="px-4 py-2 border border-gray-600 cursor-pointer hover:bg-cut-background">Settings
                </li>
                <li @click="logout" class="px-4 py-2 cursor-pointer hover:bg-cut-background">Logout</li>
              </ul>
            </div>
            <ProfileForm :show="showProfileForm" @close="closeProfileModal" @save="saveProfile" />
          </div>
        </div>
      </div>
      <div v-else class="flex justify-center items-center h-screen">
        <svg xmlns="http://www.w3.org/2000/svg" class="animate-spin h-10 w-10 text-blue-500" viewBox="0 0 24 24"
          fill="none" stroke="currentColor" stroke-width="4">
          <circle cx="12" cy="12" r="10" stroke-opacity="0.25" stroke-linecap="round" stroke-linejoin="round"></circle>
          <path d="M22 12a10 10 0 11-10-10" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1"></path>
        </svg>
      </div>
    </template>
<script setup lang="ts">
import { useAuth } from '~/composables/useAuth';
import { useCookies } from '@vueuse/integrations/useCookies';
import service from '~/service';
import type { AddFriend, Contact } from '~/models/chat.models';
import type { UserResponse } from '~/models/user.models';
import type { Paginate, Param } from "~/models/http.models";
import formatUnixToDateTime from '~/utils/time_unix';
import { useIndexStore } from '~/store/main'


const contacts = ref<Contact[] | null>(null);
const store = useIndexStore()
const search = ref('');

const cookies = useCookies();
const id = cookies.get('id') || '';
const loading = ref(false);
const showDropdown = ref(false);
const showProfileForm = ref(false);
const showAddFriendForm = ref(false);
const page = ref(1);

const props = defineProps<{
  total?: number;
  idcon?: string;
  type?: string;
  detail?: string;
  time?: number;
}>();

const emit = defineEmits(['selectContact','fakeclick']);

const newunread = ref(0);

const newctype = ref('');

const newcreated_at = ref(0);

const newcdetail = ref('');


watch(
  () => ({ total: props.total, idcon: props.idcon, type: props.type, detail: props.detail, time: props.time }),
  async ({ total: newtotal, idcon: newId, type: newtype, time: newtime, detail: newdetail }, { total: oldtotal, idcon: oldId,type: oldtype, time: oldtime, detail: olddetail }) => {
    if (newtotal !== oldtotal) {
      newunread.value = newtotal ?? 0;
    }
    if (newtype !== oldtype) {
      newctype.value = newtype ?? '';
    }

    if (newtime!== oldtime) {
      newcreated_at.value = newtime?? 0;
    }

    if (newdetail!== olddetail) {
      newcdetail.value = newdetail?? '';
    }

   
      const contact = contacts.value?.find((c) => c.id === newId);
      if (contact) {
        // Update unread property
        contact.unread = newunread.value;
        contact.type = newctype.value;
        contact.created_at = newcreated_at.value;
        contact.detail = newcdetail.value;

        // Remove the contact from its current position
        const index = contacts.value ? contacts.value.indexOf(contact) : -1;
        if (index > -1) {
          if (contacts.value) {
            contacts.value.splice(index, 1); // Removes the contact from the array
          }
        }

        // Add the contact to the beginning of the array
        if (contacts.value) {
          contacts.value.unshift(contact);
        }
      }
    
  },
  { deep: true } // Enable deep watching for the object
);

const saveProfile = (profileData: { user: UserResponse }) => {
  console.log('Saving user profile:', profileData.user);
  closeProfileModal()
};

const openProfileModal = () => {
  showProfileForm.value = true;
  showDropdown.value = false;
};

const closeProfileModal = () => {
  showProfileForm.value = false;
};


const searchContacts = () => {
  fetchUserList()
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

const addFriend = () => {
  showAddFriendForm.value = !showAddFriendForm.value;
};

const closeAddFriendModal = () => {
  showAddFriendForm.value = false;
  fetchUserList()
};

const paginate = ref<Paginate>({
  current_page: 0,
  per_page: 0,
  total_pages: 0,
  total: 0
});

const removeFriend = async (contact: Contact) => {
  const confirmed = await swalConfirmDialog({
    icon: 'warning',
    title: 'Remove Friend',
    text: `Are you sure you want to remove ${contact.display_name} from your friend list?`,
    confirmButtonText: 'Yes, remove',
    cancelButtonText: 'No, cancel'
  });

  if (confirmed) {
    const body = { user_id: contact.id };
    service.chat
      .RemoveFriend(body)
      .then((resp: any) => {
        if (resp.status.code == 200) {
          fetchUserList();
        }
      })
      .catch((err: any) => {
        console.log(err);
      });
  }
};

const fetchUserList = () => {
  loading.value = true
  const param: Param = {
    limit: 9,
    page: page.value,
    search: search.value
  }
  service.user
    .listfriend(id, param)
    .then((resp: any) => {
      if (resp.status.code == '200') {
        contacts.value = resp.data;
        paginate.value = resp.pagination;

      if (contacts.value && contacts.value.length > 0) {
       
        
       emit('fakeclick', contacts.value[0]);
    }
      }
    })
    .catch((err: any) => {
      console.log(err);
    })
    .finally(() => {
      loading.value = false;
    });
}



onMounted(async () => {
  await fetchUserList()
 
});

const { logout } = useAuth();

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value;
};

const navigateToSettings = () => {
  console.log('Navigating to Settings');
  showDropdown.value = false;

};

function selectContact(contact: Contact) {
  emit('selectContact', contact);
}
</script>

<style scoped>
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}
</style>