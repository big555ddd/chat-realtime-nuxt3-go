<template>
  <div v-if="show" class="fixed inset-0 bg-gray-900 bg-opacity-50 flex items-center justify-center z-50"
    @click.self="closeProfile">
    <div class="bg-background p-8 rounded-lg shadow-lg w-full max-w-md">
      <h2 class="text-2xl font-bold mb-6 text-center text-light">User Profile</h2>
      <div v-if="store.first_name && store.last_name && store.display_name">
        <form @submit.prevent="saveProfile">
          <div class="mb-4">
            <label for="first_name" class="block text-light">First Name:</label>
            <input type="text" id="first_name" v-model="form.first_name"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" />
          </div>
          <div class="mb-4">
            <label for="last_name" class="block text-light">Last Name:</label>
            <input type="text" id="last_name" v-model="form.last_name"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" />
          </div>
          <div class="mb-4">
            <label for="display_name" class="block text-light">Display Name:</label>
            <input type="text" id="display_name" v-model="form.display_name"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" />
          </div>
          <button type="submit"
            class="w-full bg-red hover:bg-rose-500 text-white py-2 rounded-lg transition duration-300">
            Save
          </button>
        </form>
      </div>
      <div v-else>
        <p>Loading...</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { UserResponse, UpdateProfile } from '~/models/user.models';
import service from '~/service';
import { useIndexStore } from '~/store/main'


const props = defineProps({
  show: Boolean,

});

const emit = defineEmits(['close', 'save']);
const store = useIndexStore()


// ใช้ตัวแปรภายใน component สำหรับการรับค่าจาก props
const form = ref({
  first_name: store.first_name || '',
  last_name: store.last_name || '',
  display_name: store.display_name || '',
});

// เมื่อ props.user เปลี่ยนแปลง จะทำการอัปเดตค่าภายใน form

const saveProfile = () => {
  if (store.id) {
    const data = store.$state

    // สร้างข้อมูลสำหรับการอัปเดตจาก form
    const userUpdateData: UpdateProfile = {
      first_name: form.value.first_name,
      last_name: form.value.last_name,
      display_name: form.value.display_name,
    };

    service.user
      .updateByID(store.id, userUpdateData)  // ส่งข้อมูลการอัปเดต
      .then((resp: any) => {

        if (resp.status.code == 200) {
          data.first_name = form.value.first_name
        data.last_name = form.value.last_name
        data.display_name = form.value.display_name
        }
        emit('save', { user: form.value });  // ส่งข้อมูลที่อัปเดตไป
      })
      .catch((error: any) => {
        // Handle error
        console.log(error)
        // สามารถแสดงข้อความหรือการแจ้งเตือนได้ที่นี่
      })
      .finally(() => {
        closeProfile();
      });
  } else {
    console.error('User ID is missing');
  }
};

const closeProfile = () => {
  emit('close');
};
</script>

<style scoped>
/* Add any custom styles for the modal if needed */
</style>