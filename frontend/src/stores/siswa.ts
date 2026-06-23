import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Siswa, SiswaForm } from '../types'
import { siswaApi } from '../api/modules/siswa'

export const useSiswaStore = defineStore('siswa', () => {
  const list = ref<Siswa[]>([])
  const loading = ref(false)

  async function fetchAll(params?: { kelas_id?: number; search?: string }) {
    loading.value = true
    try {
      const res = await siswaApi.getAll(params)
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function create(data: SiswaForm) {
    const res = await siswaApi.create(data)
    list.value.push(res.data.data)
    return res.data
  }

  async function update(id: number, data: Partial<SiswaForm>) {
    const res = await siswaApi.update(id, data)
    const idx = list.value.findIndex((s) => s.id === id)
    if (idx !== -1) list.value[idx] = res.data.data
    return res.data
  }

  async function remove(id: number) {
    await siswaApi.delete(id)
    list.value = list.value.filter((s) => s.id !== id)
  }

  return { list, loading, fetchAll, create, update, remove }
})
