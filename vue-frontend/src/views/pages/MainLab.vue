<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 py-8">
    <!-- Header with Navigation -->
    <div class="container mx-auto px-4 mb-8">
      <nav class="flex items-center justify-between mb-8">
        <router-link
          to="/"
          class="flex items-center text-gray-700 hover:text-blue-600 transition-colors"
        >
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          Back to Home
        </router-link>

        <div class="flex items-center space-x-4">
          <router-link
            to="/askAI"
            class="bg-white px-4 py-2 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-all text-gray-700 hover:text-blue-600"
          >
            Ask AI Assistant
          </router-link>
          <router-link
            to="/physics"
            class="bg-blue-500 px-4 py-2 rounded-lg shadow-sm hover:bg-blue-600 transition-all text-white"
          >
            Quick Physics
          </router-link>
        </div>
      </nav>
    </div>

    <!-- Main Content -->
    <div class="container mx-auto px-4 mb-12">
      <div class="text-center">
        <h1 class="text-4xl md:text-5xl font-bold text-gray-800 mb-4">
          Virtual Laboratory Platform
        </h1>
        <p class="text-xl text-gray-600 max-w-3xl mx-auto">
          Explore interactive laboratory simulations across multiple disciplines.
          Select your university and choose a lab to begin your scientific journey.
        </p>
      </div>

      <!-- University Selection -->
      <div class="mt-12 max-w-2xl mx-auto">
        <label class="block text-sm font-medium text-gray-700 mb-2">
          Select University
        </label>
        <select
          v-model="selectedUniversityId"
          @change="onUniversityChange"
          class="w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
          :disabled="universityStore.isLoading"
        >
          <option value="">Choose a university...</option>
          <option
            v-for="university in universityStore.getUniversities"
            :key="university.id"
            :value="university.id"
          >
            {{ university.name }}
          </option>
        </select>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="labStore.isLoading" class="container mx-auto px-4">
      <div class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        <span class="ml-3 text-lg text-gray-600">Loading labs...</span>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="labStore.getError" class="container mx-auto px-4">
      <div class="bg-red-50 border border-red-200 rounded-lg p-6 max-w-2xl mx-auto">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">
              Error loading labs
            </h3>
            <p class="text-sm text-red-700 mt-1">
              {{ labStore.getError }}
            </p>
          </div>
          <div class="ml-auto pl-3">
            <button
              @click="labStore.clearError()"
              class="text-red-800 hover:text-red-900"
            >
              <span class="sr-only">Dismiss</span>
              <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Labs Grid -->
    <div v-else-if="selectedUniversityId && labStore.getLabs.length > 0" class="container mx-auto px-4">
      <div class="text-center mb-8">
        <h2 class="text-3xl font-semibold text-gray-800">
          Available Laboratories
        </h2>
        <p class="text-gray-600 mt-2">
          Choose a laboratory to explore experiments and simulations
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <SubjectCard
          v-for="lab in labStore.getLabs"
          :key="lab.id"
          :title="lab.name"
          :description="lab.description"
          :icon="lab.thumbnail"
          :path="`/lab/${lab.slug}`"
          :universityId="selectedUniversityId"
          class="transform transition-all duration-300 hover:scale-105 hover:shadow-xl"
        />
      </div>

      <!-- Quick Access to Existing Practicals -->
      <div class="mt-16 border-t border-gray-200 pt-12">
        <div class="text-center mb-8">
          <h2 class="text-2xl font-semibold text-gray-800">
            Quick Access to Practicals
          </h2>
          <p class="text-gray-600 mt-2">
            Direct access to individual experiments
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 max-w-4xl mx-auto">
          <router-link
            to="/pendulum"
            class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-all text-center group"
          >
            <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-3 group-hover:bg-green-200 transition-colors">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="font-semibold text-gray-800">Pendulum</h3>
            <p class="text-sm text-gray-600 mt-1">Gravity Experiments</p>
          </router-link>

          <router-link
            to="/projectile"
            class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-all text-center group"
          >
            <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-3 group-hover:bg-blue-200 transition-colors">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <h3 class="font-semibold text-gray-800">Projectile</h3>
            <p class="text-sm text-gray-600 mt-1">Motion Analysis</p>
          </router-link>

          <router-link
            to="/newtons-law"
            class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-all text-center group"
          >
            <div class="w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-3 group-hover:bg-purple-200 transition-colors">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <h3 class="font-semibold text-gray-800">Newton's Law</h3>
            <p class="text-sm text-gray-600 mt-1">Force & Motion</p>
          </router-link>

          <router-link
            to="/machine-efficiency"
            class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-all text-center group"
          >
            <div class="w-12 h-12 bg-orange-100 rounded-full flex items-center justify-center mx-auto mb-3 group-hover:bg-orange-200 transition-colors">
              <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
            </div>
            <h3 class="font-semibold text-gray-800">Efficiency</h3>
            <p class="text-sm text-gray-600 mt-1">Machine Analysis</p>
          </router-link>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="selectedUniversityId" class="container mx-auto px-4">
      <div class="text-center py-12">
        <div class="mx-auto h-24 w-24 text-gray-400 mb-4">
          <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path>
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">No laboratories available</h3>
        <p class="text-gray-500">
          There are no laboratories configured for the selected university yet.
        </p>

        <!-- Quick fallback to existing practicals -->
        <div class="mt-6">
          <router-link
            to="/physics"
            class="inline-flex items-center px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
            Go to Physics Practicals
          </router-link>
        </div>
      </div>
    </div>

    <!-- Initial State -->
    <div v-else class="container mx-auto px-4">
      <div class="text-center py-12">
        <div class="mx-auto h-24 w-24 text-blue-500 mb-4">
          <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"></path>
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">Select a University</h3>
        <p class="text-gray-500">
          Choose your university from the dropdown above to view available laboratories.
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useUniversityStore } from '@/stores/useUniversityStore'
import { useLabStore } from '@/stores/useLabStore'
import SubjectCard from '@/components/SubjectCard.vue'

const universityStore = useUniversityStore()
const labStore = useLabStore()

const selectedUniversityId = ref('')

// Fetch universities on component mount
onMounted(async () => {
  try {
    await universityStore.fetchUniversities()
  } catch (error) {
    console.error('Failed to load universities:', error)
  }
})

// Handle university selection change
async function onUniversityChange() {
  if (!selectedUniversityId.value) {
    labStore.labs = []
    return
  }

  try {
    await labStore.fetchLabsByUniversity(selectedUniversityId.value)
    await universityStore.setCurrentUniversity(selectedUniversityId.value)
  } catch (error) {
    console.error('Failed to load labs:', error)
  }
}
</script>

<style scoped>
.container {
  max-width: 1200px;
}
</style>
