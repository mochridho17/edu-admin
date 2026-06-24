<template>
    <div>
        <!-- Filter -->
        <el-card shadow="hover" class="mb-4">
            <template #header>
                <div class="flex justify-between items-center">
                    <span>Filter</span>
                    <el-button
                        type="primary"
                        size="small"
                        @click="openDialog()"
                    >
                        + Tambah Jurnal
                    </el-button>
                </div>
            </template>
            <el-row :gutter="16">
                <el-col :span="6">
                    <el-input
                        v-model="filterNamaGuru"
                        placeholder="Nama Guru"
                        clearable
                        class="w-full"
                        @change="loadData"
                        @clear="loadData"
                    />
                </el-col>
                <el-col :span="4">
                    <el-select
                        v-model="filterMapel"
                        placeholder="Mata Pelajaran"
                        clearable
                        class="w-full"
                        @change="loadData"
                    >
                        <el-option
                            v-for="m in mapelStore.list"
                            :key="m.id"
                            :label="m.nama"
                            :value="m.id"
                        />
                    </el-select>
                </el-col>
                <el-col :span="4">
                    <el-select
                        v-model="filterKelas"
                        placeholder="Kelas"
                        clearable
                        class="w-full"
                        @change="loadData"
                    >
                        <el-option
                            v-for="k in kelasStore.list"
                            :key="k.id"
                            :label="k.nama"
                            :value="k.id"
                        />
                    </el-select>
                </el-col>
                <el-col :span="4">
                    <el-select
                        v-model="filterSemester"
                        placeholder="Semester"
                        clearable
                        class="w-full"
                        @change="loadData"
                    >
                        <el-option label="Ganjil" value="ganjil" />
                        <el-option label="Genap" value="genap" />
                    </el-select>
                </el-col>
                <el-col :span="4">
                    <el-select
                        v-model="filterTA"
                        placeholder="Tahun Ajaran"
                        clearable
                        class="w-full"
                        @change="loadData"
                    >
                        <el-option
                            v-for="ta in tahunAjaranStore.list"
                            :key="ta.id"
                            :label="`${ta.nama} ${semesterLabel(ta.semester)}`"
                            :value="ta.id"
                        />
                    </el-select>
                </el-col>
                <el-col :span="2">
                    <el-button type="primary" @click="loadData" class="w-full">
                        <el-icon><Search /></el-icon> Cari
                    </el-button>
                </el-col>
            </el-row>
        </el-card>

        <!-- Table -->
        <el-card shadow="hover">
            <el-table
                :data="pagedData"
                stripe
                v-loading="jurnalStore.loading"
                style="width: 100%"
            >
                <el-table-column type="index" label="No" width="60" />
                <el-table-column label="Hari/Tanggal" width="180">
                    <template #default="{ row }">
                        {{ formatTanggal(row.tanggal) }}
                    </template>
                </el-table-column>
                <el-table-column
                    label="Nama Guru"
                    width="200"
                    prop="nama_guru"
                />
                <el-table-column label="Mata Pelajaran" min-width="180">
                    <template #default="{ row }">
                        {{ row.mapel?.nama || "-" }}
                    </template>
                </el-table-column>
                <el-table-column label="Kelas" width="120">
                    <template #default="{ row }">
                        {{ row.kelas?.nama || "-" }}
                    </template>
                </el-table-column>
                <el-table-column label="Kegiatan Pembelajaran" min-width="250">
                    <template #default="{ row }">
                        <div class="kegiatan-cell">{{ row.kegiatan }}</div>
                    </template>
                </el-table-column>
                <el-table-column label="Semester" width="100">
                    <template #default="{ row }">
                        <el-tag
                            :type="
                                row.semester === 'ganjil'
                                    ? 'primary'
                                    : 'warning'
                            "
                            size="small"
                        >
                            {{ semesterLabel(row.semester) }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="Aksi" width="200" fixed="right">
                    <template #default="{ row }">
                        <el-button size="small" @click="openDialog(row)"
                            >Edit</el-button
                        >
                        <el-button
                            size="small"
                            type="danger"
                            @click="handleDelete(row.id)"
                            >Hapus</el-button
                        >
                    </template>
                </el-table-column>
            </el-table>

            <div
                class="flex justify-center mt-4"
                v-if="jurnalStore.list.length > pageSize"
            >
                <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :total="jurnalStore.list.length"
                    :page-sizes="[10, 20, 50]"
                    layout="total, sizes, prev, pager, next, jumper"
                    background
                />
            </div>

            <el-empty
                v-if="!jurnalStore.loading && jurnalStore.list.length === 0"
                description="Belum ada data jurnal"
            />
        </el-card>

        <!-- Dialog Create / Edit -->
        <el-dialog
            v-model="dialogVisible"
            :title="isEdit ? 'Edit Jurnal Guru' : 'Tambah Jurnal Guru'"
            width="600px"
            :close-on-click-modal="false"
        >
            <el-form
                ref="formRef"
                :model="form"
                :rules="rules"
                label-position="top"
                v-loading="submitting"
            >
                <el-row :gutter="16">
                    <el-col :span="12">
                        <el-form-item label="Hari/Tanggal" prop="tanggal">
                            <el-date-picker
                                v-model="form.tanggal"
                                type="date"
                                placeholder="Pilih Tanggal"
                                class="w-full"
                                value-format="YYYY-MM-DD"
                            />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="Nama Guru" prop="nama_guru">
                            <el-input
                                v-model="form.nama_guru"
                                placeholder="Masukkan nama guru"
                            />
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row :gutter="16">
                    <el-col :span="12">
                        <el-form-item label="Mata Pelajaran" prop="mapel_id">
                            <el-select
                                v-model="form.mapel_id"
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
                    </el-col>
                    <el-col :span="12">
                        <el-form-item label="Kelas" prop="kelas_id">
                            <el-select
                                v-model="form.kelas_id"
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
                    </el-col>
                </el-row>

                <el-row :gutter="16">
                    <el-col :span="12">
                        <el-form-item label="Semester" prop="semester">
                            <el-select
                                v-model="form.semester"
                                placeholder="Pilih Semester"
                                class="w-full"
                            >
                                <el-option label="Ganjil" value="ganjil" />
                                <el-option label="Genap" value="genap" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="12">
                        <el-form-item
                            label="Tahun Ajaran"
                            prop="tahun_ajaran_id"
                        >
                            <el-select
                                v-model="form.tahun_ajaran_id"
                                placeholder="Pilih TA"
                                class="w-full"
                            >
                                <el-option
                                    v-for="ta in tahunAjaranStore.list"
                                    :key="ta.id"
                                    :label="`${ta.nama} ${semesterLabel(ta.semester)}`"
                                    :value="ta.id"
                                />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-form-item label="Kegiatan Pembelajaran" prop="kegiatan">
                    <el-input
                        v-model="form.kegiatan"
                        type="textarea"
                        :rows="3"
                        placeholder="Deskripsikan kegiatan pembelajaran"
                    />
                </el-form-item>
            </el-form>

            <template #footer>
                <el-button @click="dialogVisible = false">Batal</el-button>
                <el-button
                    type="primary"
                    :loading="submitting"
                    @click="handleSubmit"
                >
                    {{ isEdit ? "Update" : "Simpan" }}
                </el-button>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import type { FormInstance, FormRules } from "element-plus";
import { useJurnalStore } from "../stores/jurnal";
import { useKelasStore } from "../stores/kelas";
import { useMapelStore } from "../stores/mapel";
import { useTahunAjaranStore } from "../stores/tahunAjaran";
import { SEMESTER_LABEL } from "../utils/format";
import type { Semester, JurnalGuruForm } from "../types";

const jurnalStore = useJurnalStore();
const kelasStore = useKelasStore();
const mapelStore = useMapelStore();
const tahunAjaranStore = useTahunAjaranStore();

const currentPage = ref(1);
const pageSize = ref(10);
const pagedData = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    return jurnalStore.list.slice(start, start + pageSize.value);
});

// Filters
const filterNamaGuru = ref("");
const filterMapel = ref<number>();
const filterKelas = ref<number>();
const filterSemester = ref<string>();
const filterTA = ref<number>();

// Dialog
const dialogVisible = ref(false);
const isEdit = ref(false);
const editingId = ref<number | null>(null);
const submitting = ref(false);
const formRef = ref<FormInstance>();

const form = reactive<JurnalGuruForm>({
    tanggal: "",
    nama_guru: "",
    mapel_id: 0,
    kelas_id: 0,
    kegiatan: "",
    semester: "ganjil",
    tahun_ajaran_id: 0,
});

const rules: FormRules = {
    tanggal: [{ required: true, message: "Pilih tanggal", trigger: "change" }],
    nama_guru: [
        { required: true, message: "Masukkan nama guru", trigger: "blur" },
    ],
    mapel_id: [
        {
            required: true,
            message: "Pilih mata pelajaran",
            trigger: "change",
            type: "number",
        },
    ],
    kelas_id: [
        {
            required: true,
            message: "Pilih kelas",
            trigger: "change",
            type: "number",
        },
    ],
    kegiatan: [
        {
            required: true,
            message: "Isi kegiatan pembelajaran",
            trigger: "blur",
        },
    ],
    semester: [
        { required: true, message: "Pilih semester", trigger: "change" },
    ],
    tahun_ajaran_id: [
        {
            required: true,
            message: "Pilih tahun ajaran",
            trigger: "change",
            type: "number",
        },
    ],
};

function semesterLabel(s: string) {
    return SEMESTER_LABEL[s] || s;
}

function formatTanggal(t: string) {
    return new Date(t + "T00:00:00").toLocaleDateString("id-ID", {
        weekday: "long",
        day: "numeric",
        month: "long",
        year: "numeric",
    });
}

onMounted(() => {
    Promise.all([
        kelasStore.fetchAll(),
        mapelStore.fetchAll(),
        tahunAjaranStore.fetchAll(),
    ]).then(() => {
        loadData();
    });
});

function loadData() {
    currentPage.value = 1;
    jurnalStore.fetchAll({
        nama_guru: filterNamaGuru.value || undefined,
        mapel_id: filterMapel.value,
        kelas_id: filterKelas.value,
        semester: filterSemester.value,
        tahun_ajaran_id: filterTA.value,
    });
}

function openDialog(row?: {
    id: number;
    tanggal: string;
    nama_guru: string;
    mapel_id: number;
    kelas_id: number;
    kegiatan: string;
    semester: Semester;
    tahun_ajaran_id: number;
}) {
    if (row) {
        isEdit.value = true;
        editingId.value = row.id;
        form.tanggal = row.tanggal;
        form.nama_guru = row.nama_guru;
        form.mapel_id = row.mapel_id;
        form.kelas_id = row.kelas_id;
        form.kegiatan = row.kegiatan;
        form.semester = row.semester;
        form.tahun_ajaran_id = row.tahun_ajaran_id;
    } else {
        isEdit.value = false;
        editingId.value = null;
        form.tanggal = new Date().toISOString().split("T")[0];
        form.nama_guru = "";
        form.mapel_id = 0;
        form.kelas_id = 0;
        form.kegiatan = "";
        form.semester = "ganjil";
        form.tahun_ajaran_id = 0;
    }
    dialogVisible.value = true;
}

async function handleSubmit() {
    const valid = await formRef.value?.validate().catch(() => false);
    if (!valid) return;

    submitting.value = true;
    try {
        if (isEdit.value && editingId.value) {
            await jurnalStore.update(editingId.value, { ...form });
            ElMessage.success("Jurnal berhasil diupdate");
        } else {
            await jurnalStore.create({ ...form });
            ElMessage.success("Jurnal berhasil ditambahkan");
        }
        dialogVisible.value = false;
        loadData();
    } finally {
        submitting.value = false;
    }
}

async function handleDelete(id: number) {
    try {
        await ElMessageBox.confirm(
            "Yakin ingin menghapus jurnal ini?",
            "Konfirmasi",
            {
                confirmButtonText: "Ya, Hapus",
                cancelButtonText: "Batal",
                type: "warning",
            },
        );
        await jurnalStore.remove(id);
        ElMessage.success("Jurnal berhasil dihapus");
        loadData();
    } catch {
        // cancelled
    }
}
</script>

<style scoped>
.kegiatan-cell {
    white-space: pre-line;
    line-height: 1.5;
}
.w-full {
    width: 100%;
}
.mb-4 {
    margin-bottom: 16px;
}
</style>
