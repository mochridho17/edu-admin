import apiClient from '../client'
import type { ApiResponse, TahunAjaran, TahunAjaranForm } from '../../types'

export const tahunAjaranApi = {
  getAll() {
    return apiClient.get<ApiResponse<TahunAjaran[]>>('/tahun-ajaran')
  },
  create(data: TahunAjaranForm) {
    return apiClient.post<ApiResponse<TahunAjaran>>('/tahun-ajaran', data)
  },
  update(id: number, data: Partial<TahunAjaranForm>) {
    return apiClient.put<ApiResponse<TahunAjaran>>(`/tahun-ajaran/${id}`, data)
  },
  delete(id: number) {
    return apiClient.delete<ApiResponse<null>>(`/tahun-ajaran/${id}`)
  },
  activate(id: number) {
    return apiClient.patch<ApiResponse<TahunAjaran>>(`/tahun-ajaran/${id}/activate`)
  },
}
