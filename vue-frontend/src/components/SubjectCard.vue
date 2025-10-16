<template>
  <div
    class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-200 hover:border-blue-300 transition-all duration-300 group cursor-pointer"
    @click="navigateToLab"
    :class="{'opacity-50 cursor-not-allowed': requiresUniversity && !universityId}"
  >
    <!-- Card Image -->
    <div class="h-48 bg-gradient-to-br from-blue-500 to-purple-600 relative overflow-hidden">
      <img
        v-if="icon"
        :src="icon"
        :alt="title"
        class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
      />
      <div v-else class="w-full h-full flex items-center justify-center">
        <svg class="w-16 h-16 text-white opacity-80" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path>
        </svg>
      </div>
      <div class="absolute inset-0 bg-black bg-opacity-20 group-hover:bg-opacity-10 transition-all duration-300"></div>

      <!-- Overlay for disabled state -->
      <div
        v-if="requiresUniversity && !universityId"
        class="absolute inset-0 bg-gray-500 bg-opacity-70 flex items-center justify-center"
      >
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
        </svg>
      </div>
    </div>

    <!-- Card Content -->
    <div class="p-6">
      <h3 class="text-xl font-bold text-gray-800 mb-2 group-hover:text-blue-600 transition-colors">
        {{ title }}
      </h3>
      <p class="text-gray-600 text-sm leading-relaxed mb-4 line-clamp-3">
        {{ description }}
      </p>

      <!-- Action Button -->
      <button
        @click.stop="navigateToLab"
        :disabled="requiresUniversity && !universityId"
        class="w-full bg-blue-500 text-white py-3 px-4 rounded-lg font-semibold hover:bg-blue-600 active:bg-blue-700 transition-colors duration-200 flex items-center justify-center group/btn disabled:bg-gray-400 disabled:cursor-not-allowed"
        :class="{'bg-blue-500 hover:bg-blue-600': !requiresUniversity || universityId, 'bg-gray-400': requiresUniversity && !universityId}"
      >
        <template v-if="requiresUniversity && !universityId">
          Select University
        </template>
        <template v-else>
          Enter Lab
          <svg class="w-4 h-4 ml-2 transform group-hover/btn:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path>
          </svg>
        </template>
      </button>

      <!-- Warning message for missing university -->
      <div v-if="requiresUniversity && !universityId" class="mt-2 text-xs text-red-500 text-center">
        Please select a university first
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useLabStore } from '@/stores/useLabStore'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  description: {
    type: String,
    required: true
  },
  icon: {
    type: String,
    default: ''
  },
  path: {
    type: String,
    required: true
  },
  universityId: {
    type: [String, Number],
    default: null
  },
  requiresUniversity: {
    type: Boolean,
    default: false
  }
})

const router = useRouter()
const labStore = useLabStore()

const navigateToLab = () => {
  // If university is required but not provided, don't navigate
  if (props.requiresUniversity && !props.universityId) {
    return
  }

  // For existing routes that don't need university context
  if (props.path.startsWith('/') && !props.path.startsWith('/lab/')) {
    router.push(props.path)
    return
  }

  // For new lab routes that need university context
  if (props.universityId) {
    const labSlug = props.path.replace('/lab/', '')
    const lab = labStore.getLabs.find(l => l.slug === labSlug)
    if (lab) {
      labStore.setCurrentLab(lab)
    }

    router.push({
      path: props.path,
      query: { university_id: props.universityId }
    })
  } else {
    // Fallback for routes without university context
    router.push(props.path)
  }
}
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
