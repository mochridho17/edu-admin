import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { MataPelajaran, MataPelajaranForm } from '../types'
import { mapelApi } from '../api/modules/mapel'

export const useMapelStore = defineStore('mapel', () => {
  const list = ref<MataPelajaran[]>([])
  const loading = ref(false)

  async function fetchAll() {
    loading.value = true
    try {
      const res = await mapelApi.getAll()
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function create(data: MataPelajaranForm) {
    const res = await mapelApi.create(data)
    list.value.push(res.data.data)
    return res.data
  }

  async function update(id: number, data: Partial<MataPelajaranForm>) {
    const res = await mapelApi.update(id, data)
    const idx = list.value.findIndex((m) => m.id === id)
    if (idx !== -1) list.value[idx] = res.data.data
    return res.data
  }

  async function remove(id: number) {
    await mapelApi.delete(id)
    list.value = list.value.filter((m) => m.id !== id)
  }

  return { list, loading, fetchAll, create, update, remove }
})
