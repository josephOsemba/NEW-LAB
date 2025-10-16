import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { labService } from '@/services/labService'

export const useLabStore = defineStore('lab', () => {
  const labs = ref([])
  const currentLab = ref(null)
  const loading = ref(false)
  const error = ref(null)

  const getLabs = computed(() => labs.value)
  const getCurrentLab = computed(() => currentLab.value)
  const isLoading = computed(() => loading.value)
  const getError = computed(() => error.value)

  async function fetchLabsByUniversity(universityId) {
    loading.value = true
    error.value = null
    try {
      const response = await labService.getLabs(universityId)
      labs.value = response.data
      return response.data
    } catch (err) {
      error.value = err.message || 'Failed to fetch labs'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function fetchLabBySlug(universityId, labSlug) {
    loading.value = true
    error.value = null
    try {
      const response = await labService.getLabBySlug(universityId, labSlug)
      currentLab.value = response.data
      return response.data
    } catch (err) {
      error.value = err.message || 'Failed to fetch lab details'
      throw err
    } finally {
      loading.value = false
    }
  }

  function setCurrentLab(lab) {
    currentLab.value = lab
  }

  function clearError() {
    error.value = null
  }

  return {
    labs,
    currentLab,
    loading,
    error,
    getLabs,
    getCurrentLab,
    isLoading,
    getError,
    fetchLabsByUniversity,
    fetchLabBySlug,
    setCurrentLab,
    clearError
  }
})
