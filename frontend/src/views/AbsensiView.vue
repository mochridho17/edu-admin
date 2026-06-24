<template>
    <div>
        <!-- Pilih Kelas & Tanggal -->
        <el-card shadow="hover" class="mb-4">
            <el-row :gutter="12">
                <el-col :span="3">
                    <el-select
                        v-model="selectedKelas"
                        placeholder="Pilih Kelas"
                        class="w-full"
                        @change="loadAbsensi"
                    >
                        <el-option
                            v-for="k in kelasStore.list"
                            :key="k.id"
                            :label="k.nama"
                            :value="k.id"
                        />
                    </el-select>
                </el-col>
                <el-col :span="6">
                    <div style="display: flex; gap: 8px">
                        <el-date-picker
                            v-model="selectedTanggal"
                            type="date"
                            placeholder="Pilih Tanggal"
                            style="flex: 1"
                            value-format="YYYY-MM-DD"
                            @change="loadAbsensi"
                        />
                        <el-button
                            type="primary"
                            :loading="absensiStore.loading"
                            @click="loadAbsensi"
                            style="white-space: nowrap"
                        >
                            <el-icon><Search /></el-icon> Tampilkan
                        </el-button>
                    </div>
                </el-col>
            </el-row>
        </el-card>

        <!-- Table Absensi -->
        <el-card shadow="hover" v-if="selectedKelas && selectedTanggal">
            <div class="mb-4 flex justify-between items-center">
                <span class="text-gray-600">
                    <strong>{{ selectedKelasNama }}</strong> —
                    {{ formatTanggal(selectedTanggal) }}
                </span>
                <el-button
                    type="success"
                    :loading="submitting"
                    @click="handleSimpan"
                >
                    <el-icon><Check /></el-icon> Simpan Absensi
                </el-button>
            </div>

            <el-table :data="tableData" stripe style="width: 100%">
                <el-table-column type="index" label="No" width="60" />
                <el-table-column prop="nis" label="NIS" width="120" />
                <el-table-column
                    prop="nama"
                    label="Nama Siswa"
                    min-width="200"
                />
                <el-table-column label="Status Kehadiran" width="400">
                    <template #default="{ row }">
                        <el-radio-group
                            v-model="row.status"
                            class="status-group"
                        >
                            <el-radio-button value="hadir"
                                >Hadir</el-radio-button
                            >
                            <el-radio-button value="sakit"
                                >Sakit</el-radio-button
                            >
                            <el-radio-button value="izin">Izin</el-radio-button>
                            <el-radio-button value="alfa">Alfa</el-radio-button>
                        </el-radio-group>
                    </template>
                </el-table-column>
                <el-table-column label="Keterangan" width="200">
                    <template #default="{ row }">
                        <el-input
                            v-model="row.keterangan"
                            placeholder="Opsional"
                            size="small"
                        />
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

        <!-- Empty State -->
        <el-empty
            v-else
            description="Pilih kelas dan tanggal untuk memulai absensi"
        />
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { Search, Check } from "@element-plus/icons-vue";
import { useKelasStore } from "../stores/kelas";
import { useAbsensiStore } from "../stores/absensi";
import type { StatusAbsen } from "../types";

const kelasStore = useKelasStore();
const absensiStore = useAbsensiStore();

const selectedKelas = ref<number>();
const selectedTanggal = ref<string>(new Date().toISOString().split("T")[0]);
const submitting = ref(false);

interface TableRow {
    siswa_id: number;
    nis: string;
    nama: string;
    status: StatusAbsen;
    keterangan: string;
}

const tableData = ref<TableRow[]>([]);

const selectedKelasNama = computed(() => {
    const k = kelasStore.list.find((k) => k.id === selectedKelas.value);
    return k?.nama || "";
});

onMounted(() => {
    kelasStore.fetchAll();
});

function formatTanggal(t: string) {
    return new Date(t + "T00:00:00").toLocaleDateString("id-ID", {
        weekday: "long",
        day: "numeric",
        month: "long",
        year: "numeric",
    });
}

async function loadAbsensi() {
    if (!selectedKelas.value || !selectedTanggal.value) return;

    try {
        await absensiStore.fetchByKelasTanggal(
            selectedKelas.value,
            selectedTanggal.value,
        );
        if (absensiStore.list.length > 0) {
            // Map existing data
            tableData.value = absensiStore.list.map((a) => ({
                siswa_id: a.siswa_id,
                nis: a.siswa?.nis || "",
                nama: a.siswa?.nama || "",
                status: a.status,
                keterangan: a.keterangan || "",
            }));
        } else {
            // Load siswa list if no absensi yet
            const { siswaApi } = await import("../api/modules/siswa");
            const resSiswa = await siswaApi.getAll({
                kelas_id: selectedKelas.value,
            });
            tableData.value = resSiswa.data.data.map((s) => ({
                siswa_id: s.id,
                nis: s.nis,
                nama: s.nama,
                status: "hadir" as StatusAbsen,
                keterangan: "",
            }));
        }
    } catch {
        tableData.value = [];
    }
}

async function handleSimpan() {
    if (!selectedKelas.value || !selectedTanggal.value) return;
    submitting.value = true;
    try {
        await absensiStore.submitBatch({
            kelas_id: selectedKelas.value,
            tanggal: selectedTanggal.value,
            data: tableData.value.map((r) => ({
                siswa_id: r.siswa_id,
                status: r.status,
                keterangan: r.keterangan || undefined,
            })),
        });
        ElMessage.success("Absensi berhasil disimpan");
    } finally {
        submitting.value = false;
    }
}
</script>

<style scoped>
.status-group .el-radio-button__inner {
    padding: 8px 16px;
}
</style>
