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
        <el-col :span="4">
          <el-select v-model="filterBulan" placeholder="Bulan" class="w-full" @change="loadData">
            <el-option v-for="(b, i) in bulanList" :key="i" :label="b" :value="i + 1" />
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-select v-model="filterTahun" placeholder="Tahun" class="w-full" @change="loadData">
            <el-option v-for="t in tahunList" :key="t" :label="String(t)" :value="t" />
          </el-select>
        </el-col>
      </el-row>
    </el-card>

    <!-- Table -->
    <el-card shadow="hover" v-if="absensiStore.rekapBulanan.length > 0">
      <el-table :data="absensiStore.rekapBulanan" stripe v-loading="absensiStore.loading" style="width: 100%">
        <el-table-column type="index" label="No" width="60" />
        <el-table-column prop="nis" label="NIS" width="120" />
        <el-table-column prop="nama_siswa" label="Nama" min-width="180" />
        <el-table-column prop="hadir" label="Hadir" width="80">
          <template #default="{ row }"><el-tag type="success">{{ row.hadir }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="sakit" label="Sakit" width="80">
          <template #default="{ row }"><el-tag type="warning">{{ row.sakit }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="izin" label="Izin" width="80">
          <template #default="{ row }"><el-tag type="info">{{ row.izin }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="alfa" label="Alfa" width="80">
          <template #default="{ row }"><el-tag type="danger">{{ row.alfa }}</el-tag></template>
        </el-table-column>
        <el-table-column label="Kehadiran" width="120">
          <template #default="{ row }">
            <el-progress
              :percentage="Math.round((row.hadir / row.total) * 100)"
              :status="(row.hadir / row.total) >= 0.9 ? 'success' : 'warning'"
            />
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-empty v-else description="Pilih kelas, bulan, dan tahun untuk melihat rekap" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useKelasStore } from '../stores/kelas'
import { useAbsensiStore } from '../stores/absensi'
import { BULAN_INDONESIA } from '../utils/format'

const kelasStore = useKelasStore()
const absensiStore = useAbsensiStore()

const filterKelas = ref<number>()
const filterBulan = ref(new Date().getMonth() + 1)
const filterTahun = ref(new Date().getFullYear())
const bulanList = BULAN_INDONESIA
const tahunList = Array.from({ length: 5 }, (_, i) => new Date().getFullYear() - 2 + i)

onMounted(() => kelasStore.fetchAll())

function loadData() {
  if (!filterKelas.value) return
  absensiStore.fetchRekapBulanan(filterKelas.value, filterBulan.value, filterTahun.value)
}
</script>
