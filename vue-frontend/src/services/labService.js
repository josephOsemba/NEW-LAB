import apiClient from './api'

export const labService = {
  async getLabs(universityId) {
    const response = await apiClient.get('/labs', {
      params: { university_id: universityId }
    })
    return response.data
  },

  async getLabBySlug(universityId, slug) {
    const response = await apiClient.get(`/labs/${slug}`, {
      params: { university_id: universityId }
    })
    return response.data
  },

  async getExperiments(labId) {
    const response = await apiClient.get('/experiments', {
      params: { lab_id: labId }
    })
    return response.data
  },

  async getExperimentBySlug(labId, slug) {
    const response = await apiClient.get(`/experiments/${slug}`, {
      params: { lab_id: labId }
    })
    return response.data
  }
}
