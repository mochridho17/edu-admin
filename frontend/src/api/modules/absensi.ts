import apiClient from '../client'
import type {
  ApiResponse,
  Absensi,
  AbsensiBatchForm,
  AbsensiRekapBulanan,
} from '../../types'

export const absensiApi = {
  getByKelasTanggal(kelas_id: number, tanggal: string) {
    return apiClient.get<ApiResponse<Absensi[]>>('/absensi', {
      params: { kelas_id, tanggal },
    })
  },
  submitBatch(data: AbsensiBatchForm) {
    return apiClient.post<ApiResponse<Absensi[]>>('/absensi', data)
  },
  rekapBulanan(kelas_id: number, bulan: number, tahun: number) {
    return apiClient.get<ApiResponse<AbsensiRekapBulanan[]>>(
      '/absensi/rekap/bulanan',
      { params: { kelas_id, bulan, tahun } }
    )
  },
  rekapHarian(tanggal: string) {
    return apiClient.get<ApiResponse<Record<string, unknown>>>('/absensi/rekap/harian', {
      params: { tanggal },
    })
  },
}
