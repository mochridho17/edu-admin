import apiClient from '../client'
import type { ApiResponse, JurnalGuru, JurnalGuruForm } from '../../types'

export const jurnalApi = {
  getAll(params?: {
    nama_guru?: string
    mapel_id?: number
    kelas_id?: number
    semester?: string
    tahun_ajaran_id?: number
  }) {
    return apiClient.get<ApiResponse<JurnalGuru[]>>('/jurnal-guru', { params })
  },
  create(data: JurnalGuruForm) {
    return apiClient.post<ApiResponse<JurnalGuru>>('/jurnal-guru', data)
  },
  update(id: number, data: Partial<JurnalGuruForm>) {
    return apiClient.put<ApiResponse<JurnalGuru>>(`/jurnal-guru/${id}`, data)
  },
  delete(id: number) {
    return apiClient.delete<ApiResponse<null>>(`/jurnal-guru/${id}`)
  },
}
