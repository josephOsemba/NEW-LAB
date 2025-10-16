import apiClient from './api'

export const universityService = {
  async getUniversities() {
    const response = await apiClient.get('/universities')
    return response.data
  },

  async getUniversity(universityId) {
    const response = await apiClient.get(`/universities/${universityId}`)
    return response.data
  },

  async getUniversityBySlug(slug) {
    const response = await apiClient.get(`/universities/${slug}`)
    return response.data
  }
}
