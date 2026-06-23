import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { TahunAjaran, TahunAjaranForm } from '../types'
import { tahunAjaranApi } from '../api/modules/tahunAjaran'

export const useTahunAjaranStore = defineStore('tahunAjaran', () => {
  const list = ref<TahunAjaran[]>([])
  const loading = ref(false)

  async function fetchAll() {
    loading.value = true
    try {
      const res = await tahunAjaranApi.getAll()
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function create(data: TahunAjaranForm) {
    const res = await tahunAjaranApi.create(data)
    list.value.push(res.data.data)
    return res.data
  }

  async function update(id: number, data: Partial<TahunAjaranForm>) {
    const res = await tahunAjaranApi.update(id, data)
    const idx = list.value.findIndex((t) => t.id === id)
    if (idx !== -1) list.value[idx] = res.data.data
    return res.data
  }

  async function remove(id: number) {
    await tahunAjaranApi.delete(id)
    list.value = list.value.filter((t) => t.id !== id)
  }

  async function activate(id: number) {
    // Deactivate all first, then activate the selected one
    const res = await tahunAjaranApi.activate(id)
    await fetchAll()
    return res.data
  }

  const activeTahunAjaran = () => list.value.find((t) => t.is_active)

  return { list, loading, fetchAll, create, update, remove, activate, activeTahunAjaran }
})
