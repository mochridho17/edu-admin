import apiClient from '../client'
import type { ApiResponse, Kelas, KelasForm } from '../../types'

export const kelasApi = {
  getAll(params?: { tingkat?: string; tahun_ajaran_id?: number }) {
    return apiClient.get<ApiResponse<Kelas[]>>('/kelas', { params })
  },
  getById(id: number) {
    return apiClient.get<ApiResponse<Kelas>>(`/kelas/${id}`)
  },
  create(data: KelasForm) {
    return apiClient.post<ApiResponse<Kelas>>('/kelas', data)
  },
  update(id: number, data: Partial<KelasForm>) {
    return apiClient.put<ApiResponse<Kelas>>(`/kelas/${id}`, data)
  },
  delete(id: number) {
    return apiClient.delete<ApiResponse<null>>(`/kelas/${id}`)
  },
}
