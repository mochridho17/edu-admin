import apiClient from '../client'
import type {
  ApiResponse,
  Nilai,
  NilaiBatchForm,
  NilaiRekapSiswa,
  NilaiRankingKelas,
} from '../../types'

export const nilaiApi = {
  getByKelasMapel(params: {
    kelas_id: number
    mapel_id: number
    semester: string
    tahun_ajaran_id: number
  }) {
    return apiClient.get<ApiResponse<Nilai[]>>('/nilai', { params })
  },
  submitBatch(data: NilaiBatchForm) {
    return apiClient.post<ApiResponse<Nilai[]>>('/nilai/batch', data)
  },
  rekapSiswa(siswa_id: number) {
    return apiClient.get<ApiResponse<NilaiRekapSiswa[]>>(`/nilai/rekap/siswa/${siswa_id}`)
  },
  rankingKelas(params: { kelas_id: number; semester: string; tahun_ajaran_id: number }) {
    return apiClient.get<ApiResponse<NilaiRankingKelas[]>>('/nilai/rekap/kelas', {
      params,
    })
  },
}
