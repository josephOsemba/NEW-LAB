<template>
  <div class="experiment-page min-h-screen bg-gray-50 py-6 px-4 sm:px-6 md:px-12">

    <div class="mb-4">
      <button
        @click="goBack"
        class="inline-flex items-center px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-xl shadow text-gray-700 font-medium transition"
      >
        ← Back
      </button>
    </div>

    <h1 class="text-3xl sm:text-4xl md:text-5xl font-bold text-center mb-6 text-gray-800">
      Pendulum Gravity
    </h1>

    <div class="hidden sm:flex flex-wrap justify-center gap-4 mb-8">
      <router-link
        v-for="(nav, index) in navItems"
        :key="index"
        :to="{ name: nav.name }"
        class="nav-icon flex flex-col items-center justify-center w-24 py-3 bg-white rounded-2xl shadow hover:shadow-lg transition transform hover:-translate-y-1 hover:scale-105"
      >
        <component :is="nav.icon" class="w-6 h-6 text-blue-600 mb-1" />
        <span class="text-sm font-medium text-gray-700 text-center">{{ nav.label }}</span>
      </router-link>
    </div>

    <div class="sm:hidden relative mb-8">
      <button
        @click="mobileNavOpen = !mobileNavOpen"
        class="w-full flex justify-between items-center px-4 py-3 bg-white rounded-2xl shadow text-gray-700 font-medium hover:bg-gray-100 transition"
      >
        <span>{{ mobileNavLabel }}</span>
        <svg
          :class="{ 'rotate-180': mobileNavOpen }"
          class="w-5 h-5 transition-transform"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M19 9l-7 7-7-7"></path>
        </svg>
      </button>

      <div v-if="mobileNavOpen" class="absolute left-0 right-0 mt-2 bg-white rounded-xl shadow-md overflow-hidden z-10">
        <router-link
          v-for="(nav, index) in navItems"
          :key="index"
          :to="{ name: nav.name }"
          class="block px-4 py-3 hover:bg-gray-100 transition"
          @click="mobileNavOpen = false"
        >
          <div class="flex items-center gap-2">
            <component :is="nav.icon" class="w-5 h-5 text-blue-600" />
            <span class="text-gray-700 font-medium">{{ nav.label }}</span>
          </div>
        </router-link>
      </div>
    </div>

    <div class="experiment-content bg-white p-4 sm:p-6 rounded-2xl shadow-md">
      <router-view />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  BookOpenIcon,
  BeakerIcon,
  ClipboardDocumentListIcon,
  CodeBracketIcon,
  ClipboardDocumentIcon,
  ChatBubbleLeftIcon,
} from '@heroicons/vue/24/outline'

const router = useRouter()
const route = useRoute()
const mobileNavOpen = ref(false)

const navItems = [
  { icon: BookOpenIcon, label: 'Theory', name: 'PendulumTheory' },
  { icon: BeakerIcon, label: 'Procedure', name: 'PendulumProcedure' },
  { icon: ClipboardDocumentListIcon, label: 'Self Evaluation', name: 'PendulumEvaluation' },
  { icon: CodeBracketIcon, label: 'Simulator', name: 'PendulumSimulator' },
  { icon: ClipboardDocumentIcon, label: 'Assignment', name: 'PendulumAssignment' },
  { icon: BookOpenIcon, label: 'Reference', name: 'PendulumReference' },
  { icon: ChatBubbleLeftIcon, label: 'Feedback', name: 'PendulumFeedback' },
]

const goBack = () => {
  router.back()
}

const mobileNavLabel = computed(() => {
  const current = navItems.find(n => n.name === route.name)
  return current ? current.label : 'Navigate'
})
</script>

