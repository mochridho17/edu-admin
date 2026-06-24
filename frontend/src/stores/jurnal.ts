import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { JurnalGuru, JurnalGuruForm } from '../types'
import { jurnalApi } from '../api/modules/jurnal'

export const useJurnalStore = defineStore('jurnal', () => {
  const list = ref<JurnalGuru[]>([])
  const loading = ref(false)

  async function fetchAll(params?: {
    nama_guru?: string
    mapel_id?: number
    kelas_id?: number
    semester?: string
    tahun_ajaran_id?: number
  }) {
    loading.value = true
    try {
      const res = await jurnalApi.getAll(params)
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function create(data: JurnalGuruForm) {
    const res = await jurnalApi.create(data)
    list.value.unshift(res.data.data)
    return res.data
  }

  async function update(id: number, data: Partial<JurnalGuruForm>) {
    const res = await jurnalApi.update(id, data)
    const idx = list.value.findIndex((j) => j.id === id)
    if (idx !== -1) list.value[idx] = res.data.data
    return res.data
  }

  async function remove(id: number) {
    await jurnalApi.delete(id)
    list.value = list.value.filter((j) => j.id !== id)
  }

  return { list, loading, fetchAll, create, update, remove }
})
