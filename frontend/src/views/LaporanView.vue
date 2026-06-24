<template>
    <div>
        <el-row :gutter="20">
            <!-- Jurnal Guru -->
            <el-col :xs="24" :sm="12" :lg="8">
                <el-card shadow="hover" class="report-card">
                    <div
                        class="report-header"
                        style="
                            background: linear-gradient(
                                135deg,
                                #009ef7,
                                #7239ea
                            );
                        "
                    >
                        <el-icon :size="32" color="#fff"><Notebook /></el-icon>
                        <div class="report-title">
                            <h3>Jurnal Guru</h3>
                            <p>Export kegiatan pembelajaran harian</p>
                        </div>
                    </div>
                    <div class="report-body">
                        <el-form label-position="top" size="small">
                            <el-row :gutter="12">
                                <el-col :span="12">
                                    <el-form-item label="Nama Guru">
                                        <el-input
                                            v-model="jurnalForm.nama_guru"
                                            placeholder="Opsional"
                                            clearable
                                        />
                                    </el-form-item>
                                </el-col>
                                <el-col :span="12">
                                    <el-form-item label="Mapel">
                                        <el-select
                                            v-model="jurnalForm.mapel_id"
                                            placeholder="Semua"
                                            clearable
                                            class="w-full"
                                        >
                                            <el-option
                                                v-for="m in mapelStore.list"
                                                :key="m.id"
                                                :label="m.nama"
                                                :value="m.id"
                                            />
                                        </el-select>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                            <el-row :gutter="12">
                                <el-col :span="12">
                                    <el-form-item label="Kelas">
                                        <el-select
                                            v-model="jurnalForm.kelas_id"
                                            placeholder="Semua"
                                            clearable
                                            class="w-full"
                                        >
                                            <el-option
                                                v-for="k in kelasStore.list"
                                                :key="k.id"
                                                :label="k.nama"
                                                :value="k.id"
                                            />
                                        </el-select>
                                    </el-form-item>
                                </el-col>
                                <el-col :span="12">
                                    <el-form-item label="Semester">
                                        <el-select
                                            v-model="jurnalForm.semester"
                                            placeholder="Semua"
                                            clearable
                                            class="w-full"
                                        >
                                            <el-option
                                                label="Ganjil"
                                                value="ganjil"
                                            />
                                            <el-option
                                                label="Genap"
                                                value="genap"
                                            />
                                        </el-select>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                            <el-form-item label="Tahun Ajaran">
                                <el-select
                                    v-model="jurnalForm.tahun_ajaran_id"
                                    placeholder="Semua"
                                    clearable
                                    class="w-full"
                                >
                                    <el-option
                                        v-for="ta in tahunAjaranStore.list"
                                        :key="ta.id"
                                        :label="`${ta.nama} ${ta.semester === 'ganjil' ? 'Ganjil' : 'Genap'}`"
                                        :value="ta.id"
                                    />
                                </el-select>
                            </el-form-item>
                            <el-button
                                type="primary"
                                :icon="Download"
                                :loading="loadingJurnal"
                                @click="exportJurnal"
                                class="w-full"
                            >
                                Export Excel
                            </el-button>
                        </el-form>
                    </div>
                </el-card>
            </el-col>

            <!-- Absensi -->
            <el-col :xs="24" :sm="12" :lg="8">
                <el-card shadow="hover" class="report-card">
                    <div
                        class="report-header"
                        style="
                            background: linear-gradient(
                                135deg,
                                #50cd89,
                                #009ef7
                            );
                        "
                    >
                        <el-icon :size="32" color="#fff"><Document /></el-icon>
                        <div class="report-title">
                            <h3>Absensi</h3>
                            <p>Rekap kehadiran siswa bulanan</p>
                        </div>
                    </div>
                    <div class="report-body">
                        <el-form label-position="top" size="small">
                            <el-form-item label="Kelas">
                                <el-select
                                    v-model="absensiForm.kelas_id"
                                    placeholder="Pilih Kelas"
                                    class="w-full"
                                >
                                    <el-option
                                        v-for="k in kelasStore.list"
                                        :key="k.id"
                                        :label="k.nama"
                                        :value="k.id"
                                    />
                                </el-select>
                            </el-form-item>
                            <el-row :gutter="12">
                                <el-col :span="12">
                                    <el-form-item label="Bulan">
                                        <el-select
                                            v-model="absensiForm.bulan"
                                            placeholder="Bulan"
                                            class="w-full"
                                        >
                                            <el-option
                                                v-for="(b, i) in bulanList"
                                                :key="i"
                                                :label="b"
                                                :value="i + 1"
                                            />
                                        </el-select>
                                    </el-form-item>
                                </el-col>
                                <el-col :span="12">
                                    <el-form-item label="Tahun">
                                        <el-select
                                            v-model="absensiForm.tahun"
                                            placeholder="Tahun"
                                            class="w-full"
                                        >
                                            <el-option
                                                v-for="t in tahunList"
                                                :key="t"
                                                :label="String(t)"
                                                :value="t"
                                            />
                                        </el-select>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                            <el-button
                                type="primary"
                                :icon="Download"
                                :loading="loadingAbsen"
                                @click="exportAbsensi"
                                class="w-full"
                            >
                                Export Excel
                            </el-button>
                        </el-form>
                    </div>
                </el-card>
            </el-col>

            <!-- Nilai -->
            <el-col :xs="24" :sm="12" :lg="8">
                <el-card shadow="hover" class="report-card">
                    <div
                        class="report-header"
                        style="
                            background: linear-gradient(
                                135deg,
                                #f1416c,
                                #ffa800
                            );
                        "
                    >
                        <el-icon :size="32" color="#fff"><Trophy /></el-icon>
                        <div class="report-title">
                            <h3>Nilai</h3>
                            <p>Export nilai siswa per mata pelajaran</p>
                        </div>
                    </div>
                    <div class="report-body">
                        <el-form label-position="top" size="small">
                            <el-form-item label="Kelas">
                                <el-select
                                    v-model="nilaiForm.kelas_id"
                                    placeholder="Pilih Kelas"
                                    class="w-full"
                                >
                                    <el-option
                                        v-for="k in kelasStore.list"
                                        :key="k.id"
                                        :label="k.nama"
                                        :value="k.id"
                                    />
                                </el-select>
                            </el-form-item>
                            <el-form-item label="Mata Pelajaran">
                                <el-select
                                    v-model="nilaiForm.mapel_id"
                                    placeholder="Pilih Mapel"
                                    class="w-full"
                                >
                                    <el-option
                                        v-for="m in mapelStore.list"
                                        :key="m.id"
                                        :label="m.nama"
                                        :value="m.id"
                                    />
                                </el-select>
                            </el-form-item>
                            <el-button
                                type="primary"
                                :icon="Download"
                                :loading="loadingNilai"
                                @click="exportNilai"
                                class="w-full"
                            >
                                Export Excel
                            </el-button>
                        </el-form>
                    </div>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { Document, Download, Notebook, Trophy } from "@element-plus/icons-vue";
import apiClient from "../api/client";
import { useKelasStore } from "../stores/kelas";
import { useMapelStore } from "../stores/mapel";
import { useTahunAjaranStore } from "../stores/tahunAjaran";
import { BULAN_INDONESIA } from "../utils/format";

const kelasStore = useKelasStore();
const mapelStore = useMapelStore();
const tahunAjaranStore = useTahunAjaranStore();
const loadingAbsen = ref(false);
const loadingNilai = ref(false);
const loadingJurnal = ref(false);

const bulanList = BULAN_INDONESIA;
const tahunList = Array.from(
    { length: 5 },
    (_, i) => new Date().getFullYear() - 2 + i,
);

const absensiForm = reactive({
    kelas_id: undefined as number | undefined,
    bulan: new Date().getMonth() + 1,
    tahun: new Date().getFullYear(),
});
const nilaiForm = reactive({
    kelas_id: undefined as number | undefined,
    mapel_id: undefined as number | undefined,
});
const jurnalForm = reactive({
    nama_guru: "",
    mapel_id: undefined as number | undefined,
    kelas_id: undefined as number | undefined,
    semester: undefined as string | undefined,
    tahun_ajaran_id: undefined as number | undefined,
});

onMounted(() => {
    kelasStore.fetchAll();
    mapelStore.fetchAll();
    tahunAjaranStore.fetchAll();
});

async function downloadFile(url: string, filename: string) {
    try {
        const response = await apiClient.get(url, { responseType: "blob" });
        const blob = new Blob([response.data]);
        const link = document.createElement("a");
        link.href = URL.createObjectURL(blob);
        link.download = filename;
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
        URL.revokeObjectURL(link.href);
        ElMessage.success("File berhasil diunduh");
    } catch {
        ElMessage.error("Gagal mengunduh file. Hubungi admin.");
    }
}

async function exportJurnal() {
    loadingJurnal.value = true;
    try {
        const params = new URLSearchParams();
        if (jurnalForm.nama_guru)
            params.append("nama_guru", jurnalForm.nama_guru);
        if (jurnalForm.mapel_id)
            params.append("mapel_id", String(jurnalForm.mapel_id));
        if (jurnalForm.kelas_id)
            params.append("kelas_id", String(jurnalForm.kelas_id));
        if (jurnalForm.semester) params.append("semester", jurnalForm.semester);
        if (jurnalForm.tahun_ajaran_id)
            params.append(
                "tahun_ajaran_id",
                String(jurnalForm.tahun_ajaran_id),
            );
        const qs = params.toString();
        await downloadFile(
            `/laporan/jurnal-guru${qs ? "?" + qs : ""}`,
            "Laporan_Jurnal_Guru.xlsx",
        );
    } finally {
        loadingJurnal.value = false;
    }
}

async function exportAbsensi() {
    if (!absensiForm.kelas_id) {
        ElMessage.warning("Pilih kelas terlebih dahulu");
        return;
    }
    loadingAbsen.value = true;
    try {
        const kelas = kelasStore.list.find(
            (k) => k.id === absensiForm.kelas_id,
        );
        const namaFile = `Laporan_Absensi_${kelas?.nama || ""}_Bulan${absensiForm.bulan}_${absensiForm.tahun}.xlsx`;
        const url = `/laporan/absensi?kelas_id=${absensiForm.kelas_id}&bulan=${absensiForm.bulan}&tahun=${absensiForm.tahun}`;
        await downloadFile(url, namaFile);
    } finally {
        loadingAbsen.value = false;
    }
}

async function exportNilai() {
    if (!nilaiForm.kelas_id || !nilaiForm.mapel_id) {
        ElMessage.warning("Pilih kelas dan mapel terlebih dahulu");
        return;
    }
    loadingNilai.value = true;
    try {
        const kelas = kelasStore.list.find((k) => k.id === nilaiForm.kelas_id);
        const mapel = mapelStore.list.find((m) => m.id === nilaiForm.mapel_id);
        const namaFile = `Laporan_Nilai_${kelas?.nama || ""}_${mapel?.nama || ""}.xlsx`;
        const url = `/laporan/nilai?kelas_id=${nilaiForm.kelas_id}&mapel_id=${nilaiForm.mapel_id}`;
        await downloadFile(url, namaFile);
    } finally {
        loadingNilai.value = false;
    }
}
</script>

<style scoped>
.report-card {
    border-radius: 12px;
    overflow: hidden;
    transition:
        transform 0.2s,
        box-shadow 0.2s;
}
.report-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1) !important;
}
.report-header {
    display: flex;
    align-items: center;
    gap: 14px;
    padding: 20px 20px 18px;
    margin: -1px -1px 0 -1px;
}
.report-title h3 {
    margin: 0;
    font-size: 16px;
    font-weight: 600;
    color: #fff;
}
.report-title p {
    margin: 2px 0 0;
    font-size: 12px;
    color: rgba(255, 255, 255, 0.8);
}
.report-body {
    padding: 16px 0 0;
}
.w-full {
    width: 100%;
}
</style>
