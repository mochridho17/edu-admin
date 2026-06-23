import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Kelas, KelasForm } from '../types'
import { kelasApi } from '../api/modules/kelas'

export const useKelasStore = defineStore('kelas', () => {
  const list = ref<Kelas[]>([])
  const loading = ref(false)

  async function fetchAll(params?: { tingkat?: string; tahun_ajaran_id?: number }) {
    loading.value = true
    try {
      const res = await kelasApi.getAll(params)
      list.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function create(data: KelasForm) {
    const res = await kelasApi.create(data)
    list.value.push(res.data.data)
    return res.data
  }

  async function update(id: number, data: Partial<KelasForm>) {
    const res = await kelasApi.update(id, data)
    const idx = list.value.findIndex((k) => k.id === id)
    if (idx !== -1) list.value[idx] = res.data.data
    return res.data
  }

  async function remove(id: number) {
    await kelasApi.delete(id)
    list.value = list.value.filter((k) => k.id !== id)
  }

  return { list, loading, fetchAll, create, update, remove }
})
