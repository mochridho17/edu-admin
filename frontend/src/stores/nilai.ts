import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Nilai, NilaiBatchForm, NilaiRekapSiswa, NilaiRankingKelas } from '../types'
import { nilaiApi } from '../api/modules/nilai'

export const useNilaiStore = defineStore('nilai', () => {
  const list = ref<Nilai[]>([])
  const rekapSiswa = ref<NilaiRekapSiswa[]>([])
  const rankingKelas = ref<NilaiRankingKelas[]>([])
  const loading = ref(false)

  async function fetchByKelasMapel(params: {
    kelas_id: number
    mapel_id: number
    semester: string
    tahun_ajaran_id: number
  }) {
    loading.value = true
    try {
      const res = await nilaiApi.getByKelasMapel(params)
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function submitBatch(data: NilaiBatchForm) {
    const res = await nilaiApi.submitBatch(data)
    list.value = res.data.data
    return res.data
  }

  async function fetchRekapSiswa(siswa_id: number) {
    loading.value = true
    try {
      const res = await nilaiApi.rekapSiswa(siswa_id)
      rekapSiswa.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function fetchRankingKelas(params: {
    kelas_id: number
    semester: string
    tahun_ajaran_id: number
  }) {
    loading.value = true
    try {
      const res = await nilaiApi.rankingKelas(params)
      rankingKelas.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  return {
    list, rekapSiswa, rankingKelas, loading,
    fetchByKelasMapel, submitBatch, fetchRekapSiswa, fetchRankingKelas,
  }
})
