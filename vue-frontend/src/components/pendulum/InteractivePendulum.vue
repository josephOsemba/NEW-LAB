<template>
  <div class="interactive-pendulum">
    <div class="pendulum-visualization relative w-full h-64 bg-gradient-to-b from-blue-50 to-gray-100 rounded-lg border border-gray-300 overflow-hidden">
      <!-- Stand -->
      <div class="absolute top-0 left-1/2 transform -translate-x-1/2 w-4 h-8 bg-gray-600 rounded-t-lg"></div>
      <div class="absolute top-8 left-1/2 transform -translate-x-1/2 w-32 h-2 bg-gray-600"></div>

      <!-- Pendulum string -->
      <div
        class="absolute top-10 left-1/2 transform origin-top pendulum-string"
        :style="stringStyle"
      >
        <div class="w-1 h-full bg-gray-700 mx-auto"></div>
      </div>

      <!-- Bob -->
      <div
        class="absolute bob rounded-full shadow-lg cursor-pointer"
        :style="bobStyle"
        @mousedown="startDrag"
        @touchstart="startDrag"
      ></div>

      <!-- Angle indicator -->
      <div
        class="absolute top-10 left-1/2 transform origin-top angle-indicator"
        :style="angleIndicatorStyle"
      >
        <div class="w-24 h-1 bg-red-500 opacity-50"></div>
      </div>

      <!-- Length indicator -->
      <div class="absolute left-1/2 top-10 transform -translate-x-1/2 length-indicator">
        <div class="text-xs text-gray-600 bg-white px-2 py-1 rounded border">
          {{ props.length.toFixed(2) }}m
        </div>
      </div>
    </div>

    <!-- Animation controls -->
    <div class="mt-4 flex justify-center space-x-4">
      <button
        @click="resetPendulum"
        class="px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700"
      >
        Reset
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'

const emit = defineEmits(['data-recorded', 'period-calculated'])

const props = defineProps({
  length: {
    type: Number,
    default: 1.0
  },
  angle: {
    type: Number,
    default: 15
  },
  gravity: {
    type: Number,
    default: 9.81
  },
  isPlaying: {
    type: Boolean,
    default: false
  }
})

// Reactive state
const currentAngle = ref(props.angle)
const isDragging = ref(false)
const animationId = ref(null)
const oscillationStartTime = ref(null)
const periodMeasurements = ref([])
const oscillationCounter = ref(0)

// Computed styles
const stringStyle = computed(() => {
  const angleRad = (currentAngle.value * Math.PI) / 180
  return {
    transform: `rotate(${currentAngle.value}deg)`,
    height: `${props.length * 100}px`
  }
})

const bobStyle = computed(() => {
  const angleRad = (currentAngle.value * Math.PI) / 180
  const x = Math.sin(angleRad) * props.length * 100
  const y = Math.cos(angleRad) * props.length * 100
  return {
    left: `calc(50% + ${x}px)`,
    top: `${10 + y}px`,
    width: '24px',
    height: '24px',
    backgroundColor: '#3b82f6'
  }
})

const angleIndicatorStyle = computed(() => {
  return {
    transform: `rotate(${currentAngle.value}deg)`
  }
})

// Pendulum physics
const calculatePeriod = (length, gravity) => {
  // T = 2π√(l/g)
  return 2 * Math.PI * Math.sqrt(length / gravity)
}

const updatePendulum = (length, angle, gravity) => {
  currentAngle.value = angle
  if (props.isPlaying) {
    startOscillation()
  }
}

const startOscillation = () => {
  if (animationId.value) {
    cancelAnimationFrame(animationId.value)
  }

  oscillationStartTime.value = Date.now()
  oscillationCounter.value = 0
  periodMeasurements.value = []

  animatePendulum()
}

const stopOscillation = () => {
  if (animationId.value) {
    cancelAnimationFrame(animationId.value)
    animationId.value = null
  }
}

const resetPendulum = () => {
  stopOscillation()
  currentAngle.value = props.angle
}

const animatePendulum = (timestamp = 0) => {
  if (!oscillationStartTime.value) {
    oscillationStartTime.value = timestamp
  }

  const elapsed = (timestamp - oscillationStartTime.value) / 1000
  const period = calculatePeriod(props.length, props.gravity)
  const angularFrequency = 2 * Math.PI / period

  // Simple harmonic motion equation
  currentAngle.value = props.angle * Math.cos(angularFrequency * elapsed)

  // Detect complete oscillations
  if (Math.abs(currentAngle.value - props.angle) < 0.1 && elapsed > period / 2) {
    oscillationCounter.value++

    if (oscillationCounter.value === 20) {
      const totalTime = elapsed
      const measuredPeriod = totalTime / 20

      emit('data-recorded', {
        timeFor20Oscillations: totalTime,
        period: measuredPeriod
      })

      emit('period-calculated', measuredPeriod)
      oscillationCounter.value = 0
      oscillationStartTime.value = timestamp
    }
  }

  if (props.isPlaying) {
    animationId.value = requestAnimationFrame(animatePendulum)
  }
}

const startDrag = (e) => {
  isDragging.value = true
  stopOscillation()

  const handleDrag = (moveEvent) => {
    if (!isDragging.value) return

    const clientX = moveEvent.clientX || moveEvent.touches[0].clientX
    const clientY = moveEvent.clientY || moveEvent.touches[0].clientY
    const rect = e.currentTarget.getBoundingClientRect()

    const centerX = rect.left + rect.width / 2
    const centerY = rect.top - (props.length * 100) + 10

    const deltaX = clientX - centerX
    const deltaY = clientY - centerY

    let angle = Math.atan2(deltaX, deltaY) * (180 / Math.PI)
    angle = Math.max(Math.min(angle, 45), -45)

    currentAngle.value = angle
  }

  const stopDrag = () => {
    isDragging.value = false
    document.removeEventListener('mousemove', handleDrag)
    document.removeEventListener('touchmove', handleDrag)
    document.removeEventListener('mouseup', stopDrag)
    document.removeEventListener('touchend', stopDrag)

    if (props.isPlaying) {
      startOscillation()
    }
  }

  document.addEventListener('mousemove', handleDrag)
  document.addEventListener('touchmove', handleDrag)
  document.addEventListener('mouseup', stopDrag)
  document.addEventListener('touchend', stopDrag)
}

// Watchers
watch(() => props.isPlaying, (newVal) => {
  if (newVal) {
    startOscillation()
  } else {
    stopOscillation()
  }
})

watch(() => props.length, (newVal) => {
  updatePendulum(newVal, props.angle, props.gravity)
})

watch(() => props.angle, (newVal) => {
  if (!isDragging.value) {
    currentAngle.value = newVal
  }
})

// Lifecycle
onMounted(() => {
  if (props.isPlaying) {
    startOscillation()
  }
})

onUnmounted(() => {
  stopOscillation()
})

// Expose methods
defineExpose({
  startOscillation,
  stopOscillation,
  updatePendulum,
  resetPendulum
})
</script>

<style scoped>
.pendulum-string {
  transition: height 0.3s ease;
}

.bob {
  transition: all 0.1s ease;
  transform: translate(-50%, -50%);
}

.bob:active {
  transform: translate(-50%, -50%) scale(1.1);
}

.angle-indicator {
  pointer-events: none;
}

.length-indicator {
  pointer-events: none;
}
</style>
