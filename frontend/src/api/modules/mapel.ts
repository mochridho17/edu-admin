import apiClient from '../client'
import type { ApiResponse, MataPelajaran, MataPelajaranForm } from '../../types'

export const mapelApi = {
  getAll() {
    return apiClient.get<ApiResponse<MataPelajaran[]>>('/mapel')
  },
  create(data: MataPelajaranForm) {
    return apiClient.post<ApiResponse<MataPelajaran>>('/mapel', data)
  },
  update(id: number, data: Partial<MataPelajaranForm>) {
    return apiClient.put<ApiResponse<MataPelajaran>>(`/mapel/${id}`, data)
  },
  delete(id: number) {
    return apiClient.delete<ApiResponse<null>>(`/mapel/${id}`)
  },
}
