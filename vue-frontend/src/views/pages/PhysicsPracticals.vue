<template>
  <div class="physics-branches-page bg-gray-50 min-h-screen py-10 px-4 sm:px-6 md:px-12 lg:px-20">
    <!-- Page Header -->
    <div class="w-full text-center mb-12">
      <h1 class="text-3xl sm:text-4xl md:text-5xl font-bold mb-4 text-gray-900">
        Explore Physics Branches
      </h1>
      <p class="text-gray-600 text-base sm:text-lg md:text-xl max-w-3xl mx-auto leading-relaxed">
        Select a branch of physics to explore sub-branches and hands-on experiments.
      </p>
    </div>

    <!-- Branches Grid -->
    <div v-if="!selectedBranch" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="branch in branches"
        :key="branch.name"
        @click="selectBranch(branch)"
        class="branch-card bg-white rounded-3xl shadow-md hover:shadow-xl transition-transform transform hover:-translate-y-1 hover:scale-105 p-8 flex flex-col items-center justify-center cursor-pointer group"
      >
        <div
          class="w-20 h-20 flex items-center justify-center rounded-full bg-blue-100 mb-4 transition group-hover:bg-blue-200"
        >
          <component :is="branch.icon" class="w-10 h-10 text-blue-600" />
        </div>
        <h2 class="text-lg sm:text-xl font-semibold text-gray-800 text-center">
          {{ branch.label }}
        </h2>
        <p class="mt-2 text-gray-500 text-sm text-center">
          {{ branch.description }}
        </p>
      </div>
    </div>

    <!-- Sub-branches Grid -->
    <div v-else-if="selectedBranch && !selectedSubBranch" class="w-full">
      <button
        @click="selectedBranch = null"
        class="mb-6 text-blue-600 font-medium hover:underline flex items-center gap-1"
      >
        <ArrowLeftIcon class="w-5 h-5"/> Back to Branches
      </button>
      <h2 class="text-2xl sm:text-3xl font-bold mb-6">{{ selectedBranch.label }} Sub-branches</h2>

      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="sub in selectedBranch.subBranches"
          :key="sub.name"
          @click="selectSubBranch(sub)"
          class="sub-branch-card bg-white rounded-2xl shadow-md hover:shadow-xl transition-transform transform hover:-translate-y-1 hover:scale-105 flex flex-col items-center justify-center p-6 cursor-pointer group"
        >
          <div class="w-16 h-16 flex items-center justify-center rounded-full bg-indigo-100 mb-3 transition group-hover:bg-indigo-200">
            <component :is="sub.icon" class="w-8 h-8 text-indigo-600" />
          </div>
          <h3 class="text-lg font-semibold text-gray-800 text-center">{{ sub.label }}</h3>
          <p class="mt-1 text-gray-500 text-sm text-center">{{ sub.description }}</p>
        </div>
      </div>
    </div>

    <!-- Experiments Grid -->
    <div v-else-if="selectedSubBranch" class="w-full">
      <button
        @click="selectedSubBranch = null"
        class="mb-6 text-blue-600 font-medium hover:underline flex items-center gap-1"
      >
        <ArrowLeftIcon class="w-5 h-5"/> Back to {{ selectedBranch.label }} Sub-branches
      </button>
      <h2 class="text-2xl sm:text-3xl font-bold mb-6">{{ selectedSubBranch.label }} Experiments</h2>

      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="(exp, index) in selectedSubBranch.experiments"
          :key="index"
          class="experiment-card bg-white rounded-3xl shadow-md hover:shadow-xl transition-transform transform hover:-translate-y-1 hover:scale-105 flex flex-col w-full"
        >
          <div class="flex items-center p-4 sm:p-6 border-b border-gray-200">
            <div class="flex-shrink-0 w-12 h-12 sm:w-14 sm:h-14 rounded-full bg-blue-100 flex items-center justify-center mr-4">
              <component :is="exp.icon" class="w-6 h-6 sm:w-8 sm:h-8 text-blue-600" />
            </div>
            <h2 class="text-base sm:text-lg md:text-xl font-semibold text-gray-800 leading-snug">
              {{ exp.title }}
            </h2>
          </div>

          <div class="flex-1 p-4 sm:p-6">
            <p class="text-gray-600 text-sm sm:text-base leading-relaxed">
              {{ exp.description }}
            </p>
          </div>

          <div class="p-4 sm:p-6 border-t border-gray-200">
            <router-link
              :to="{ name: exp.routeName }"
              class="inline-block w-full text-center bg-blue-600 text-white font-medium py-2 sm:py-3 rounded-xl hover:bg-blue-700 transition"
            >
              View Experiment Methodology →
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  ClockIcon,
  ArrowUpIcon,
  BoltIcon,
  BeakerIcon,
  BookOpenIcon,
  ArrowLeftIcon,
  FireIcon,
  Cog6ToothIcon,
  StarIcon,
  HeartIcon,
  GlobeAltIcon,
  CpuChipIcon,
  WrenchScrewdriverIcon
} from '@heroicons/vue/24/outline'
import { ref } from "vue"

// Branches structure with icons, descriptions, sub-branches, and experiments
const branches = [
  {
    label: "Classical Physics",
    name: "classical",
    description: "Newtonian mechanics, motion, forces, and classical phenomena",
    icon: ClockIcon,
    subBranches: [
      {
        label: "Mechanics",
        name: "mechanics",
        description: "Study of motion, forces, and energy",
        icon: ArrowUpIcon,
        experiments: [
          {
            icon: ClockIcon,
            title: "Simple Pendulum",
            description: "Measure gravitational acceleration using a pendulum",
            routeName: "PendulumTheory"
          },
          {
            icon: ArrowUpIcon,
            title: "Projectile Motion",
            description: "Analyze horizontal and vertical motion components",
            routeName: "projectile"
          }
        ]
      },
      {
        label: "Electromagnetism",
        name: "electromagnetism",
        description: "Electricity, magnetism, and circuits",
        icon: BoltIcon,
        experiments: [
          {
            icon: BoltIcon,
            title: "Ohm's Law Experiment",
            description: "Study the relationship between voltage, current, and resistance",
            routeName: "newtonsLaw"
          }
        ]
      }
    ]
  },
  {
    label: "Modern Physics",
    name: "modern",
    description: "Quantum mechanics, relativity, and nuclear physics",
    icon: BeakerIcon,
    subBranches: [
      {
        label: "Quantum Mechanics",
        name: "quantum",
        description: "Behavior of particles at atomic scale",
        icon: FireIcon,
        experiments: [
          {
            icon: BookOpenIcon,
            title: "Photoelectric Effect",
            description: "Explore emission of electrons using light",
            routeName: "machineEfficiency"
          }
        ]
      }
    ]
  },
  {
    label: "Applied Physics",
    name: "applied",
    description: "Practical applications of physics in technology",
    icon: Cog6ToothIcon,
    subBranches: [
      {
        label: "Medical Physics",
        name: "medical",
        description: "Imaging (X-rays, MRI), radiation therapy",
        icon: HeartIcon,
        experiments: [
          {
            icon: HeartIcon,
            title: "X-ray Imaging Principles",
            description: "Understand the physics behind medical imaging",
            routeName: "xrayImaging"
          }
        ]
      },
      {
        label: "Engineering Physics",
        name: "engineering",
        description: "Electronics, nanotechnology, instrumentation",
        icon: WrenchScrewdriverIcon,
        experiments: [
          {
            icon: WrenchScrewdriverIcon,
            title: "Semiconductor Properties",
            description: "Explore the behavior of semiconductor materials",
            routeName: "semiconductor"
          }
        ]
      },
      {
        label: "Computational Physics",
        name: "computational",
        description: "Simulations and numerical methods for physical problems",
        icon: CpuChipIcon,
        experiments: [
          {
            icon: CpuChipIcon,
            title: "Monte Carlo Simulation",
            description: "Use random sampling to solve physical problems",
            routeName: "monteCarlo"
          }
        ]
      }
    ]
  },
  {
    label: "Interdisciplinary Fields",
    name: "interdisciplinary",
    description: "Astrophysics, Biophysics, Geophysics, and specialized fields",
    icon: StarIcon,
    subBranches: [
      {
        label: "Astrophysics & Cosmology",
        name: "astrophysics",
        description: "Stars, galaxies, universe evolution",
        icon: StarIcon,
        experiments: [
          {
            icon: StarIcon,
            title: "Stellar Spectroscopy",
            description: "Analyze light from stars to determine composition",
            routeName: "spectroscopy"
          }
        ]
      },
      {
        label: "Biophysics",
        name: "biophysics",
        description: "Physical processes in biology",
        icon: HeartIcon,
        experiments: [
          {
            icon: HeartIcon,
            title: "Protein Folding Dynamics",
            description: "Study the physical principles of protein structure",
            routeName: "proteinFolding"
          }
        ]
      },
      {
        label: "Geophysics",
        name: "geophysics",
        description: "Earth's physical properties and processes",
        icon: GlobeAltIcon,
        experiments: [
          {
            icon: GlobeAltIcon,
            title: "Seismic Wave Analysis",
            description: "Study earthquake waves and earth structure",
            routeName: "seismicWaves"
          }
        ]
      }
    ]
  }
]

const selectedBranch = ref(null)
const selectedSubBranch = ref(null)

function selectBranch(branch) {
  selectedBranch.value = branch
  selectedSubBranch.value = null
}

function selectSubBranch(sub) {
  selectedSubBranch.value = sub
}
</script>

<style scoped>
.branch-card, .sub-branch-card, .experiment-card {
  cursor: pointer;
}
</style>
