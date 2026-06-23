<template>
  <div>
    <!-- Welcome + Demo Badge -->
    <div class="welcome-section">
      <div>
        <h3 class="welcome-text">Halo, <strong>{{ authStore.user?.username }}</strong>! Selamat datang kembali.</h3>
        <p class="welcome-sub">
          {{ hariIni }}
          <el-tag v-if="authStore.isDemoMode" type="warning" size="small" class="ml-2">🧪 Mode Demo</el-tag>
        </p>
      </div>
    </div>

    <!-- Stat Cards -->
    <el-row :gutter="20" class="mb-5">
      <el-col :span="6" v-for="stat in stats" :key="stat.label">
        <el-card shadow="never" class="stat-card">
          <div class="stat-content">
            <div class="stat-info">
              <p class="stat-value">{{ stat.value }}</p>
              <p class="stat-label">{{ stat.label }}</p>
            </div>
            <div class="stat-icon-wrapper" :style="{ background: stat.bg }">
              <el-icon :size="22" :color="stat.color">
                <component :is="stat.icon" />
              </el-icon>
            </div>
          </div>
          <div class="stat-footer">
            <span class="stat-trend" :class="stat.trend">
              <el-icon :size="12"><ArrowUp v-if="stat.trend === 'up'" /><ArrowDown v-else /></el-icon>
              {{ stat.trendText }}
            </span>
            <span class="stat-period">Bulan ini</span>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- Charts Row -->
    <el-row :gutter="20">
      <!-- Chart Absensi -->
      <el-col :span="16">
        <el-card shadow="never">
          <template #header>
            <div class="flex justify-between items-center">
              <span>📊 Absensi Hari Ini</span>
              <el-tag type="success" size="small" effect="plain">{{ hariIni }}</el-tag>
            </div>
          </template>
          <div class="chart-container">
            <div class="chart-bars">
              <div class="bar-item" v-for="item in chartData" :key="item.label">
                <span class="bar-value">{{ item.value }}</span>
                <div
                  class="bar"
                  :style="{ height: item.percent + '%', background: item.color }"
                ></div>
                <span class="bar-label">{{ item.label }}</span>
              </div>
            </div>
            <div class="chart-summary">
              <div class="summary-item" v-for="item in chartData" :key="item.label">
                <span class="summary-dot" :style="{ background: item.color }"></span>
                <span>{{ item.label }}: <strong>{{ item.value }}</strong></span>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>

      <!-- Info Sekolah -->
      <el-col :span="8">
        <el-card shadow="never">
          <template #header>
            <span>🏫 Ringkasan Sekolah</span>
          </template>
          <div class="info-list">
            <div class="info-item" v-for="item in infoSekolah" :key="item.label">
              <span class="info-label">{{ item.label }}</span>
              <span class="info-value">{{ item.value }}</span>
            </div>
            <el-divider class="info-divider" />
            <div class="info-active">
              <el-icon :size="16" color="#50cd89"><Check /></el-icon>
              <span>Tahun Ajaran Aktif: <strong>2025/2026 - Ganjil</strong></span>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import {
  UserFilled, School, Reading, Document,
  ArrowUp, ArrowDown, Check,
} from '@element-plus/icons-vue'

const authStore = useAuthStore()
const hariIni = new Date().toLocaleDateString('id-ID', {
  weekday: 'long', day: 'numeric', month: 'long', year: 'numeric',
})

const stats = ref([
  {
    label: 'Total Siswa', value: '256 Siswa', icon: UserFilled,
    bg: '#e8f3ff', color: '#009ef7', trend: 'up', trendText: '+12',
  },
  {
    label: 'Total Kelas', value: '8 Kelas', icon: School,
    bg: '#e8fff3', color: '#50cd89', trend: 'up', trendText: '+1',
  },
  {
    label: 'Mata Pelajaran', value: '12 Mapel', icon: Reading,
    bg: '#f0e8ff', color: '#7239ea', trend: 'up', trendText: '+2',
  },
  {
    label: 'Rata-rata Kehadiran', value: '92%', icon: Document,
    bg: '#fff4de', color: '#ffa800', trend: 'down', trendText: '-3%',
  },
])

const chartData = ref([
  { label: 'Hadir', value: 42, percent: 85, color: '#50cd89' },
  { label: 'Sakit', value: 3, percent: 15, color: '#ffa800' },
  { label: 'Izin', value: 2, percent: 10, color: '#009ef7' },
  { label: 'Alfa', value: 1, percent: 5, color: '#f1416c' },
])

const infoSekolah = ref([
  { label: 'Nama Sekolah', value: 'SDN 01 Menteng' },
  { label: 'NPSN', value: '20105678' },
  { label: 'Kepala Sekolah', value: 'Hj. Siti Aminah, S.Pd.' },
  { label: 'Jumlah Guru', value: '12 Orang' },
  { label: 'Akreditasi', value: 'A (Unggul)' },
])

onMounted(() => {
  if (!authStore.isDemoMode) {
    // TODO: fetch dari API
  }
})
</script>

<style scoped>
.welcome-section {
  margin-bottom: 24px;
}
.welcome-text {
  font-size: 16px;
  font-weight: 400;
  color: #6c757d;
}
.welcome-text strong {
  color: #1e1e2d;
}
.welcome-sub {
  font-size: 13px;
  color: #a2a3b7;
  margin-top: 4px;
}

/* ====== STAT CARDS ====== */
.stat-card {
  transition: all 0.2s ease;
  cursor: pointer;
}
.stat-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08) !important;
}
.stat-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
.stat-value {
  font-size: 24px;
  font-weight: 700;
  color: #1e1e2d;
  margin: 0;
}
.stat-label {
  font-size: 13px;
  color: #a2a3b7;
  margin-top: 4px;
}
.stat-icon-wrapper {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.stat-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 16px;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
}
.stat-trend {
  font-size: 12px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 2px;
}
.stat-trend.up { color: #50cd89; }
.stat-trend.down { color: #f1416c; }
.stat-period {
  font-size: 11px;
  color: #a2a3b7;
}

/* ====== CHART ====== */
.chart-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}
.chart-bars {
  display: flex;
  justify-content: center;
  gap: 32px;
  align-items: flex-end;
  height: 200px;
  padding: 0 20px;
}
.bar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}
.bar-value {
  font-size: 14px;
  font-weight: 700;
  color: #1e1e2d;
}
.bar {
  width: 48px;
  border-radius: 6px 6px 0 0;
  min-height: 8px;
  transition: height 0.5s ease;
}
.bar-label {
  font-size: 12px;
  color: #6c757d;
  font-weight: 500;
}
.chart-summary {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
  justify-content: center;
}
.summary-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #6c757d;
}
.summary-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  display: inline-block;
}

/* ====== INFO LIST ====== */
.info-list {
  display: flex;
  flex-direction: column;
  gap: 0;
}
.info-item {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid #f0f0f0;
}
.info-item:last-child {
  border-bottom: none;
}
.info-label {
  font-size: 13px;
  color: #a2a3b7;
}
.info-value {
  font-size: 13px;
  font-weight: 600;
  color: #1e1e2d;
}
.info-divider {
  margin: 4px 0;
}
.info-active {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #50cd89;
}
.info-active strong {
  color: #1e1e2d;
}
</style>
