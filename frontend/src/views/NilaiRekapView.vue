<template>
  <div>
    <el-tabs v-model="activeTab">
      <el-tab-pane label="Ranking Kelas" name="ranking">
        <!-- Filter -->
        <el-card shadow="hover" class="mb-4">
          <el-row :gutter="16">
            <el-col :span="6">
              <el-select v-model="rankFilterKelas" placeholder="Pilih Kelas" class="w-full" @change="loadRanking">
                <el-option v-for="k in kelasStore.list" :key="k.id" :label="k.nama" :value="k.id" />
              </el-select>
            </el-col>
            <el-col :span="4">
              <el-select v-model="rankFilterSemester" placeholder="Semester" class="w-full" @change="loadRanking">
                <el-option label="Ganjil" value="ganjil" />
                <el-option label="Genap" value="genap" />
              </el-select>
            </el-col>
            <el-col :span="4">
              <el-select v-model="rankFilterTA" placeholder="Tahun Ajaran" class="w-full" @change="loadRanking">
                <el-option v-for="ta in tahunAjaranStore.list" :key="ta.id"
                  :label="`${ta.nama} ${ta.semester}`" :value="ta.id" />
              </el-select>
            </el-col>
          </el-row>
        </el-card>

        <el-card shadow="hover" v-if="nilaiStore.rankingKelas.length > 0">
          <el-table :data="nilaiStore.rankingKelas" stripe v-loading="nilaiStore.loading" style="width: 100%">
            <el-table-column label="Peringkat" width="100">
              <template #default="{ row }">
                <el-tag :type="row.peringkat === 1 ? 'danger' : row.peringkat <= 3 ? 'warning' : 'info'" size="large">
                  #{{ row.peringkat }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="nis" label="NIS" width="120" />
            <el-table-column prop="nama_siswa" label="Nama" min-width="200" />
            <el-table-column prop="rata_rata" label="Rata-rata" width="120">
              <template #default="{ row }">
                <strong>{{ row.rata_rata.toFixed(1) }}</strong>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-tab-pane>

      <el-tab-pane label="Rekap Per Siswa" name="siswa">
        <el-card shadow="hover" class="mb-4">
          <el-row :gutter="16">
            <el-col :span="8">
              <el-select
                v-model="rekapSiswaId"
                placeholder="Pilih Siswa"
                filterable
                class="w-full"
                @change="loadRekapSiswa"
              >
                <el-option
                  v-for="s in semuaSiswa"
                  :key="s.id"
                  :label="`${s.nis} - ${s.nama}`"
                  :value="s.id"
                />
              </el-select>
            </el-col>
          </el-row>
        </el-card>

        <el-card shadow="hover" v-if="nilaiStore.rekapSiswa.length > 0">
          <el-table :data="nilaiStore.rekapSiswa" stripe style="width: 100%">
            <el-table-column prop="mapel_kode" label="Kode" width="100" />
            <el-table-column prop="mapel_nama" label="Mata Pelajaran" min-width="200" />
            <el-table-column prop="kkm" label="KKM" width="80" />
            <el-table-column prop="tugas" label="Tugas" width="80" />
            <el-table-column prop="uts" label="UTS" width="80" />
            <el-table-column prop="uas" label="UAS" width="80" />
            <el-table-column prop="akhir" label="Akhir" width="80">
              <template #default="{ row }">
                <el-tag :type="row.tuntas ? 'success' : 'danger'">{{ row.akhir }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="Status" width="100">
              <template #default="{ row }">
                <el-tag :type="row.tuntas ? 'success' : 'danger'" size="small">
                  {{ row.tuntas ? 'Tuntas' : 'Remidi' }}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useKelasStore } from '../stores/kelas'
import { useTahunAjaranStore } from '../stores/tahunAjaran'
import { useNilaiStore } from '../stores/nilai'
import { siswaApi } from '../api/modules/siswa'
import type { Siswa } from '../types'

const kelasStore = useKelasStore()
const tahunAjaranStore = useTahunAjaranStore()
const nilaiStore = useNilaiStore()

const activeTab = ref('ranking')
const rankFilterKelas = ref<number>()
const rankFilterSemester = ref('ganjil')
const rankFilterTA = ref<number>()
const rekapSiswaId = ref<number>()
const semuaSiswa = ref<Siswa[]>([])

onMounted(() => {
  kelasStore.fetchAll()
  tahunAjaranStore.fetchAll()
})

function loadRanking() {
  if (!rankFilterKelas.value || !rankFilterTA.value) return
  nilaiStore.fetchRankingKelas({
    kelas_id: rankFilterKelas.value,
    semester: rankFilterSemester.value,
    tahun_ajaran_id: rankFilterTA.value,
  })
}

async function loadRekapSiswa() {
  if (!rekapSiswaId.value) return
  await nilaiStore.fetchRekapSiswa(rekapSiswaId.value)
}

// Watch tab change to load siswa list for rekap
import { watch } from 'vue'
watch(activeTab, async (tab) => {
  if (tab === 'siswa' && semuaSiswa.value.length === 0) {
    const res = await siswaApi.getAll()
    semuaSiswa.value = res.data.data
  }
})
</script>
