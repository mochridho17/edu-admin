import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Absensi, AbsensiBatchForm, AbsensiRekapBulanan } from '../types'
import { absensiApi } from '../api/modules/absensi'

export const useAbsensiStore = defineStore('absensi', () => {
  const list = ref<Absensi[]>([])
  const rekapBulanan = ref<AbsensiRekapBulanan[]>([])
  const loading = ref(false)

  async function fetchByKelasTanggal(kelas_id: number, tanggal: string) {
    loading.value = true
    try {
      const res = await absensiApi.getByKelasTanggal(kelas_id, tanggal)
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function submitBatch(data: AbsensiBatchForm) {
    const res = await absensiApi.submitBatch(data)
    list.value = res.data.data
    return res.data
  }

  async function fetchRekapBulanan(kelas_id: number, bulan: number, tahun: number) {
    loading.value = true
    try {
      const res = await absensiApi.rekapBulanan(kelas_id, bulan, tahun)
      rekapBulanan.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  return { list, rekapBulanan, loading, fetchByKelasTanggal, submitBatch, fetchRekapBulanan }
})
