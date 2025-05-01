<template>
  <div class=" p-4 bg-background text-white ">
    <!-- Main Layout -->
    <div class="flex flex-col lg:flex-row space-y-8 lg:space-y-0 lg:space-x-8">
      <!-- Left Section -->
      <div class="lg:w-1/2 flex flex-col space-y-8">
        <!-- Contact Information (Top) -->
        <div
          class="flex flex-col justify-center items-center rounded-lg border border-gray-700 p-6 bg-cut-background shadow-lg">
          <p class="text-3xl font-bold text-center">Chat Beat</p>
          <div class="w-full h-1 bg-red-600 my-4"></div> <!-- Red line -->
          <div class="space-y-6">
            <div class="flex items-center">
              <svg class="h-8 w-8 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              <p v-if="refToken" class="pl-3 text-lg">79 Kongka Road, Tumbon Paknam, Krabi Thailand 81000</p>
              <p v-else class="pl-3 text-lg">88/8 M12, อำเภอเมืองขอนแก่น ขอนแก่น 40000</p>
            </div>

            <div class="flex items-center">
              <svg class="h-8 w-8 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                stroke-linecap="round" stroke-linejoin="round">
                <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z" />
                <polyline points="22,6 12,13 2,6" />
              </svg>
              <p class="pl-3 text-lg">ChatBeat@gmail.com</p>
            </div>

            <div class="flex items-center">
              <svg class="h-8 w-8 text-white" width="24" height="24" viewBox="0 0 24 24" stroke-width="2"
                stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" />
                <path
                  d="M5 4h4l2 5l-2.5 1.5a11 11 0 0 0 5 5l1.5 -2.5l5 2v4a2 2 0 0 1 -2 2a16 16 0 0 1 -15 -15a2 2 0 0 1 2 -2" />
              </svg>
              <p class="pl-3 text-lg">098-984-9255</p>
            </div>
          </div>
        </div>

        <!-- Google Map (Bottom) -->
        <div class="rounded-lg overflow-hidden shadow-lg">
          <iframe
            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3825.737837458974!2d102.8044411749136!3d16.488807184253126!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31228a9920cec07b%3A0x272578cc2ddadf6d!2z4LiB4Lij4Li04LiXIOC4hOC4reC4meC4i-C4seC4peC5geC4l-C4meC4l-C5jA!5e0!3m2!1sth!2sth!4v1732417418144!5m2!1sth!2sth"
            width="100%" height="300" style="border:0;" loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"></iframe>
        </div>
      </div>

      <!-- Right Section: Contact Form -->
      <div class="lg:w-1/2 bg-cut-background p-6 rounded-lg shadow-lg">
        <p class="text-2xl font-semibold mb-4 text-center">Contact Us</p>
        <form class="space-y-6">
          <div>
            <label for="firstName" class="block text-sm font-medium mb-1">First Name</label>
            <input type="text" id="firstName" class="w-full p-3 rounded-lg bg-gray-700 text-white border border-gray-600"
              placeholder="Enter your first name" />
          </div>
          <div>
            <label for="lastName" class="block text-sm font-medium mb-1">Last Name</label>
            <input type="text" id="lastName" class="w-full p-3 rounded-lg bg-gray-700 text-white border border-gray-600"
              placeholder="Enter your last name" />
          </div>
          <div>
            <label for="description" class="block text-sm font-medium mb-1">Description</label>
            <textarea id="description" rows="4"
              class="w-full p-3 rounded-lg bg-gray-700 text-white border border-gray-600"
              placeholder="Write your message"></textarea>
          </div>
          <button type="submit"
           class="w-full bg-red hover:bg-rose-500 text-white font-semibold py-3 rounded-lg shadow-lg">
            Send Message
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const refToken = useCookie("lang");
</script>

<style scoped>
/* Add custom styles here if needed */
</style>
