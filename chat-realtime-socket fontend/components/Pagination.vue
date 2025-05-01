<template>
  <div v-if="totalPages > 1" class="flex justify-between items-center mt-4">
    <button
      @click="prevPage"
      :disabled="currentPage === 1"
      class="bg-red text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition duration-300"
    >
      Previous
    </button>
    <span>Page {{ currentPage }} of {{ totalPages }}</span>
    <button
      v-if="currentPage < totalPages"
      @click="nextPage"
      class="bg-red text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition duration-300"
    >
      Next
    </button>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  currentPage: {
    type: Number,
    required: true,
  },
  totalPages: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits(['prevPage', 'nextPage']);

const prevPage = () => {
  emit('prevPage');
};

const nextPage = () => {
  emit('nextPage');
};
</script>

<style scoped>
/* Add any additional styling here */
</style>