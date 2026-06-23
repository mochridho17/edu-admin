import apiClient from '../client'
import type { ApiResponse, Siswa, SiswaForm } from '../../types'

export const siswaApi = {
  getAll(params?: { kelas_id?: number; search?: string }) {
    return apiClient.get<ApiResponse<Siswa[]>>('/siswa', { params })
  },
  getById(id: number) {
    return apiClient.get<ApiResponse<Siswa>>(`/siswa/${id}`)
  },
  create(data: SiswaForm) {
    return apiClient.post<ApiResponse<Siswa>>('/siswa', data)
  },
  update(id: number, data: Partial<SiswaForm>) {
    return apiClient.put<ApiResponse<Siswa>>(`/siswa/${id}`, data)
  },
  delete(id: number) {
    return apiClient.delete<ApiResponse<null>>(`/siswa/${id}`)
  },
}
