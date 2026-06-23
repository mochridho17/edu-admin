<template>
  <div>
    <!-- Filter -->
    <el-card shadow="hover" class="mb-4">
      <el-row :gutter="16">
        <el-col :span="6">
          <el-select v-model="filterKelas" placeholder="Pilih Kelas" class="w-full" @change="loadData">
            <el-option v-for="k in kelasStore.list" :key="k.id" :label="k.nama" :value="k.id" />
          </el-select>
        </el-col>
        <el-col :span="6">
          <el-select v-model="filterMapel" placeholder="Pilih Mapel" class="w-full" @change="loadData">
            <el-option v-for="m in mapelStore.list" :key="m.id" :label="m.nama" :value="m.id" />
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-select v-model="filterSemester" placeholder="Semester" class="w-full" @change="loadData">
            <el-option label="Ganjil" value="ganjil" />
            <el-option label="Genap" value="genap" />
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-select v-model="filterTA" placeholder="Tahun Ajaran" class="w-full" @change="loadData">
            <el-option v-for="ta in tahunAjaranStore.list" :key="ta.id"
              :label="`${ta.nama} ${ta.semester}`" :value="ta.id" />
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-button type="success" :loading="submitting" @click="handleSimpan">
            <el-icon><Check /></el-icon> Simpan Nilai
          </el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- Table -->
    <el-card shadow="hover" v-if="tableData.length > 0">
      <el-table :data="tableData" stripe style="width: 100%">
        <el-table-column type="index" label="No" width="60" />
        <el-table-column prop="nis" label="NIS" width="120" />
        <el-table-column prop="nama" label="Nama" min-width="180" />
        <el-table-column label="Tugas" width="120">
          <template #default="{ row }">
            <el-input-number v-model="row.tugas" :min="0" :max="100" size="small" controls-position="right" />
          </template>
        </el-table-column>
        <el-table-column label="UTS" width="120">
          <template #default="{ row }">
            <el-input-number v-model="row.uts" :min="0" :max="100" size="small" controls-position="right" />
          </template>
        </el-table-column>
        <el-table-column label="UAS" width="120">
          <template #default="{ row }">
            <el-input-number v-model="row.uas" :min="0" :max="100" size="small" controls-position="right" />
          </template>
        </el-table-column>
        <el-table-column label="Akhir" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.akhir !== null" :type="row.akhir >= (row.kkm || 75) ? 'success' : 'danger'">
              {{ row.akhir }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-empty v-else description="Pilih kelas, mapel, dan semester untuk mulai input nilai" />
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Check } from '@element-plus/icons-vue'
import { useKelasStore } from '../stores/kelas'
import { useMapelStore } from '../stores/mapel'
import { useTahunAjaranStore } from '../stores/tahunAjaran'
import { useNilaiStore } from '../stores/nilai'

const kelasStore = useKelasStore()
const mapelStore = useMapelStore()
const tahunAjaranStore = useTahunAjaranStore()
const nilaiStore = useNilaiStore()

const filterKelas = ref<number>()
const filterMapel = ref<number>()
const filterSemester = ref('ganjil')
const filterTA = ref<number>()
const submitting = ref(false)

interface TableRow {
  siswa_id: number
  nis: string
  nama: string
  kkm: number
  tugas: number | null
  uts: number | null
  uas: number | null
  akhir: number | null
}

const tableData = ref<TableRow[]>([])

onMounted(() => {
  kelasStore.fetchAll()
  mapelStore.fetchAll()
  tahunAjaranStore.fetchAll()
})

async function loadData() {
  if (!filterKelas.value || !filterMapel.value || !filterTA.value) return
  try {
    await nilaiStore.fetchByKelasMapel({
      kelas_id: filterKelas.value,
      mapel_id: filterMapel.value,
      semester: filterSemester.value as 'ganjil' | 'genap',
      tahun_ajaran_id: filterTA.value,
    })
    if (nilaiStore.list.length > 0) {
      const mapel = mapelStore.list.find((m) => m.id === filterMapel.value)
      tableData.value = nilaiStore.list.map((n) => ({
        siswa_id: n.siswa_id,
        nis: n.siswa?.nis || '',
        nama: n.siswa?.nama || '',
        kkm: mapel?.kkm || 75,
        tugas: n.tugas,
        uts: n.uts,
        uas: n.uas,
        akhir: n.akhir,
      }))
    } else {
      // Load from siswa list
      const { siswaApi } = await import('../api/modules/siswa')
      const res = await siswaApi.getAll({ kelas_id: filterKelas.value })
      const mapel = mapelStore.list.find((m) => m.id === filterMapel.value)
      tableData.value = res.data.data.map((s) => ({
        siswa_id: s.id,
        nis: s.nis,
        nama: s.nama,
        kkm: mapel?.kkm || 75,
        tugas: null,
        uts: null,
        uas: null,
        akhir: null,
      }))
    }
  } catch {
    tableData.value = []
  }
}

// Compute nilai akhir
function computeAkhir(tugas: number | null, uts: number | null, uas: number | null): number | null {
  if (tugas === null && uts === null && uas === null) return null
  const t = tugas || 0
  const s = uts || 0
  const a = uas || 0
  return Math.round(t * 0.2 + s * 0.3 + a * 0.5)
}

// Watch for changes to recompute akhir
watch(tableData, (rows) => {
  rows.forEach((row) => {
    row.akhir = computeAkhir(row.tugas, row.uts, row.uas)
  })
}, { deep: true })

async function handleSimpan() {
  if (!filterKelas.value || !filterMapel.value || !filterTA.value) return
  submitting.value = true
  try {
    await nilaiStore.submitBatch({
      kelas_id: filterKelas.value,
      mapel_id: filterMapel.value,
      semester: filterSemester.value as 'ganjil' | 'genap',
      tahun_ajaran_id: filterTA.value,
      data: tableData.value.map((r) => ({
        siswa_id: r.siswa_id,
        tugas: r.tugas || undefined,
        uts: r.uts || undefined,
        uas: r.uas || undefined,
      })),
    })
    ElMessage.success('Nilai berhasil disimpan')
  } finally { submitting.value = false }
}
</script>
