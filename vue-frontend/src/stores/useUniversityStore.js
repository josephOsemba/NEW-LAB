import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { universityService } from '@/services/universityService'

export const useUniversityStore = defineStore('university', () => {
  const universities = ref([])
  const currentUniversity = ref(null)
  const loading = ref(false)
  const error = ref(null)

  const getUniversities = computed(() => universities.value)
  const getCurrentUniversity = computed(() => currentUniversity.value)
  const isLoading = computed(() => loading.value)
  const getError = computed(() => error.value)

  async function fetchUniversities() {
    loading.value = true
    error.value = null
    try {
      const response = await universityService.getUniversities()
      universities.value = response.data
      return response.data
    } catch (err) {
      error.value = err.message || 'Failed to fetch universities'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function setCurrentUniversity(universityId) {
    const university = universities.value.find(u => u.id === universityId)
    if (university) {
      currentUniversity.value = university
      return university
    }

    // If university not in store, fetch it
    try {
      const response = await universityService.getUniversity(universityId)
      currentUniversity.value = response.data
      return response.data
    } catch (err) {
      error.value = 'Failed to fetch university details'
      throw err
    }
  }

  function clearError() {
    error.value = null
  }

  return {
    universities,
    currentUniversity,
    loading,
    error,
    getUniversities,
    getCurrentUniversity,
    isLoading,
    getError,
    fetchUniversities,
    setCurrentUniversity,
    clearError
  }
})
