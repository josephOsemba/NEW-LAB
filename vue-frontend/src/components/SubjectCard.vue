<template>
  <div
    @click="goToLab"
    class="cursor-pointer bg-white shadow-lg rounded-xl hover:shadow-2xl transition-transform transform hover:scale-105 flex flex-col text-center overflow-hidden"
  >
    <!-- Full-width card image -->
    <img
      v-if="icon"
      :src="icon"
      alt="icon"
      class="w-full h-32 md:h-40 lg:h-48 object-cover"
    />

    <div class="p-6 flex flex-col items-center">
      <h3 class="text-xl md:text-2xl font-bold mb-2">{{ title }}</h3>
      <p class="text-gray-600 mb-4 text-sm md:text-base">{{ description }}</p>
      <button
        class="bg-blue-600 text-white px-6 py-2 md:px-8 md:py-3 rounded-lg hover:bg-blue-700 transition"
      >
        Enter Lab
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "SubjectCard",
  props: {
    title: { type: String, required: true },
    description: { type: String, required: true },
    icon: { type: String, default: "" },
    path: { type: String, required: true },
  },
  methods: {
    goToLab() {
      this.$router.push(this.path); // Navigate to subject page
    },
  },
};
</script>
