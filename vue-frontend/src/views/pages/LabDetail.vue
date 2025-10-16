<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="container mx-auto px-4">
        <div class="flex items-center justify-between py-4">
          <div class="flex items-center space-x-4">
            <router-link
              to="/mainLab"
              class="flex items-center text-gray-600 hover:text-blue-600 transition-colors"
            >
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              Back to Labs
            </router-link>
            <div class="h-6 w-px bg-gray-300"></div>
            <router-link
              to="/askAI"
              class="text-gray-600 hover:text-blue-600 transition-colors"
            >
              Ask AI
            </router-link>
          </div>

          <div class="text-sm text-gray-500">
            {{ universityStore.getCurrentUniversity?.name }}
          </div>
        </div>
      </div>
    </div>

    <div class="container mx-auto px-4 py-8">
      <!-- Lab Header -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8">
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <h1 class="text-3xl font-bold text-gray-800 mb-2">{{ labStore.getCurrentLab?.name }}</h1>
            <p class="text-gray-600 text-lg">{{ labStore.getCurrentLab?.description }}</p>

            <!-- Lab Metadata -->
            <div class="flex items-center space-x-6 mt-4 text-sm text-gray-500">
              <div class="flex items-center">
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                {{ labStore.getCurrentLab?.config?.duration || 45 }} minutes
              </div>
              <div class="flex items-center">
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
                </svg>
                {{ labStore.getCurrentLab?.config?.apparatus_count || 5 }} apparatus
              </div>
            </div>
          </div>

          <!-- Quick Actions -->
          <div class="flex space-x-3 ml-6">
            <router-link
              to="/askAI"
              class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors flex items-center"
            >
              <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"></path>
              </svg>
              AI Help
            </router-link>
          </div>
        </div>
      </div>

      <!-- Experiments Section -->
      <div class="mb-8">
        <h2 class="text-2xl font-semibold text-gray-800 mb-6">Available Experiments</h2>

        <div v-if="experiments.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="experiment in experiments"
            :key="experiment.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all cursor-pointer group"
            @click="navigateToExperiment(experiment)"
          >
            <div class="flex items-center justify-between mb-3">
              <h3 class="text-lg font-semibold text-gray-800 group-hover:text-blue-600 transition-colors">
                {{ experiment.title }}
              </h3>
              <span
                class="px-2 py-1 text-xs rounded-full capitalize"
                :class="getDifficultyClass(experiment.difficulty)"
              >
                {{ experiment.difficulty }}
              </span>
            </div>

            <p class="text-gray-600 text-sm mb-4 line-clamp-2">
              {{ experiment.summary }}
            </p>

            <div class="flex items-center justify-between text-sm text-gray-500">
              <span>{{ experiment.config?.apparatus_count || 3 }} apparatus</span>
              <span class="flex items-center">
                <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                {{ experiment.config?.duration || 30 }}min
              </span>
            </div>
          </div>
        </div>

        <div v-else class="text-center py-12 bg-white rounded-lg border border-gray-200">
          <div class="mx-auto h-16 w-16 text-gray-400 mb-4">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-2">Experiments Coming Soon</h3>
          <p class="text-gray-500 mb-6">
            Laboratory experiments are currently being developed.
          </p>
          <router-link
            to="/physics"
            class="inline-flex items-center px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
          >
            Explore Existing Practicals
          </router-link>
        </div>
      </div>

      <!-- Integration with Existing Practicals -->
      <div class="border-t border-gray-200 pt-8">
        <h2 class="text-xl font-semibold text-gray-800 mb-6">Related Practical Experiments</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <router-link
            v-if="labStore.getCurrentLab?.slug === 'physics'"
            to="/pendulum"
            class="bg-white p-4 rounded-lg border border-gray-200 hover:border-blue-300 hover:shadow-md transition-all group"
          >
            <h3 class="font-semibold text-gray-800 group-hover:text-blue-600 transition-colors">
              Simple Pendulum Experiment
            </h3>
            <p class="text-sm text-gray-600 mt-1">Study gravitational acceleration using pendulum motion</p>
            <div class="flex items-center text-sm text-blue-600 mt-2">
              <span>View Experiment</span>
              <svg class="w-4 h-4 ml-1 transform group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path>
              </svg>
            </div>
          </router-link>

          <router-link
            v-if="labStore.getCurrentLab?.slug === 'physics'"
            to="/projectile"
            class="bg-white p-4 rounded-lg border border-gray-200 hover:border-blue-300 hover:shadow-md transition-all group"
          >
            <h3 class="font-semibold text-gray-800 group-hover:text-blue-600 transition-colors">
              Projectile Motion Analysis
            </h3>
            <p class="text-sm text-gray-600 mt-1">Analyze projectile trajectory and range</p>
            <div class="flex items-center text-sm text-blue-600 mt-2">
              <span>View Experiment</span>
              <svg class="w-4 h-4 ml-1 transform group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path>
              </svg>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useLabStore } from '@/stores/useLabStore'
import { useUniversityStore } from '@/stores/useUniversityStore'
import { labService } from '@/services/labService'

const route = useRoute()
const router = useRouter()
const labStore = useLabStore()
const universityStore = useUniversityStore()

const experiments = ref([])
const loading = ref(false)

const labSlug = computed(() => route.params.labSlug)

onMounted(async () => {
  const universityId = route.query.university_id
  const slug = labSlug.value

  if (universityId && slug) {
    try {
      loading.value = true
      await labStore.fetchLabBySlug(universityId, slug)

      // Fetch experiments for this lab
      if (labStore.getCurrentLab?.id) {
        const response = await labService.getExperiments(labStore.getCurrentLab.id)
        experiments.value = response.data || []
      }
    } catch (error) {
      console.error('Failed to load lab details:', error)
    } finally {
      loading.value = false
    }
  }
})

function getDifficultyClass(difficulty) {
  const classes = {
    beginner: 'bg-green-100 text-green-800',
    intermediate: 'bg-yellow-100 text-yellow-800',
    advanced: 'bg-red-100 text-red-800'
  }
  return classes[difficulty] || 'bg-gray-100 text-gray-800'
}

function navigateToExperiment(experiment) {
  // For now, redirect to existing practicals based on experiment type
  if (experiment.slug.includes('pendulum')) {
    router.push('/pendulum')
  } else if (experiment.slug.includes('projectile')) {
    router.push('/projectile')
  } else {
    // Default to physics practicals
    router.push('/physics')
  }
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
