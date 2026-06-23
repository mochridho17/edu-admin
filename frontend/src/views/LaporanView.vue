<template>
  <div>
    <el-row :gutter="20">
      <!-- Export Absensi -->
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="flex justify-between items-center">
              <span class="font-semibold">Laporan Absensi</span>
              <el-icon :size="24"><Document /></el-icon>
            </div>
          </template>
          <el-form label-width="140px">
            <el-form-item label="Kelas">
              <el-select v-model="absensiForm.kelas_id" placeholder="Pilih Kelas" class="w-full">
                <el-option v-for="k in kelasStore.list" :key="k.id" :label="k.nama" :value="k.id" />
              </el-select>
            </el-form-item>
            <el-form-item label="Bulan">
              <el-select v-model="absensiForm.bulan" placeholder="Bulan" class="w-full">
                <el-option v-for="(b, i) in bulanList" :key="i" :label="b" :value="i + 1" />
              </el-select>
            </el-form-item>
            <el-form-item label="Tahun">
              <el-select v-model="absensiForm.tahun" placeholder="Tahun" class="w-full">
                <el-option v-for="t in tahunList" :key="t" :label="String(t)" :value="t" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" :icon="Download" :loading="loadingAbsen" @click="exportAbsensi">
                Export Excel
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>

      <!-- Export Nilai -->
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="flex justify-between items-center">
              <span class="font-semibold">Laporan Nilai</span>
              <el-icon :size="24"><Document /></el-icon>
            </div>
          </template>
          <el-form label-width="140px">
            <el-form-item label="Kelas">
              <el-select v-model="nilaiForm.kelas_id" placeholder="Pilih Kelas" class="w-full">
                <el-option v-for="k in kelasStore.list" :key="k.id" :label="k.nama" :value="k.id" />
              </el-select>
            </el-form-item>
            <el-form-item label="Mata Pelajaran">
              <el-select v-model="nilaiForm.mapel_id" placeholder="Pilih Mapel" class="w-full">
                <el-option v-for="m in mapelStore.list" :key="m.id" :label="m.nama" :value="m.id" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" :icon="Download" :loading="loadingNilai" @click="exportNilai">
                Export Excel
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Document, Download } from '@element-plus/icons-vue'
import apiClient from '../api/client'
import { useKelasStore } from '../stores/kelas'
import { useMapelStore } from '../stores/mapel'
import { BULAN_INDONESIA } from '../utils/format'

const kelasStore = useKelasStore()
const mapelStore = useMapelStore()
const loadingAbsen = ref(false)
const loadingNilai = ref(false)

const bulanList = BULAN_INDONESIA
const tahunList = Array.from({ length: 5 }, (_, i) => new Date().getFullYear() - 2 + i)

const absensiForm = reactive({ kelas_id: undefined as number | undefined, bulan: new Date().getMonth() + 1, tahun: new Date().getFullYear() })
const nilaiForm = reactive({ kelas_id: undefined as number | undefined, mapel_id: undefined as number | undefined })

onMounted(() => {
  kelasStore.fetchAll()
  mapelStore.fetchAll()
})

async function downloadFile(url: string, filename: string) {
  try {
    const response = await apiClient.get(url, { responseType: 'blob' })
    const blob = new Blob([response.data])
    const link = document.createElement('a')
    link.href = URL.createObjectURL(blob)
    link.download = filename
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    URL.revokeObjectURL(link.href)
    ElMessage.success('File berhasil diunduh')
  } catch {
    ElMessage.error('Gagal mengunduh file. Hubungi admin.')
  }
}

async function exportAbsensi() {
  if (!absensiForm.kelas_id) {
    ElMessage.warning('Pilih kelas terlebih dahulu')
    return
  }
  loadingAbsen.value = true
  try {
    const kelas = kelasStore.list.find(k => k.id === absensiForm.kelas_id)
    const namaFile = `Laporan_Absensi_${kelas?.nama || ''}_Bulan${absensiForm.bulan}_${absensiForm.tahun}.xlsx`
    const url = `/laporan/absensi?kelas_id=${absensiForm.kelas_id}&bulan=${absensiForm.bulan}&tahun=${absensiForm.tahun}`
    await downloadFile(url, namaFile)
  } finally {
    loadingAbsen.value = false
  }
}

async function exportNilai() {
  if (!nilaiForm.kelas_id || !nilaiForm.mapel_id) {
    ElMessage.warning('Pilih kelas dan mapel terlebih dahulu')
    return
  }
  loadingNilai.value = true
  try {
    const kelas = kelasStore.list.find(k => k.id === nilaiForm.kelas_id)
    const mapel = mapelStore.list.find(m => m.id === nilaiForm.mapel_id)
    const namaFile = `Laporan_Nilai_${kelas?.nama || ''}_${mapel?.nama || ''}.xlsx`
    const url = `/laporan/nilai?kelas_id=${nilaiForm.kelas_id}&mapel_id=${nilaiForm.mapel_id}`
    await downloadFile(url, namaFile)
  } finally {
    loadingNilai.value = false
  }
}
</script>
