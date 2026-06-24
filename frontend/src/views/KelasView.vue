<template>
    <div>
        <!-- Filter -->
        <el-card shadow="hover" class="mb-4">
            <template #header>
                <div class="flex justify-between items-center">
                    <span>Filter</span>
                    <el-button type="primary" size="small" @click="openDialog()"
                        >+ Tambah Kelas</el-button
                    >
                </div>
            </template>
            <el-row :gutter="16">
                <el-col :span="8">
                    <el-select
                        v-model="filterTingkat"
                        placeholder="Filter Tingkat"
                        clearable
                        class="w-full"
                        @change="loadData"
                    >
                        <el-option label="Kelas 1" value="1" />
                        <el-option label="Kelas 2" value="2" />
                        <el-option label="Kelas 3" value="3" />
                        <el-option label="Kelas 4" value="4" />
                        <el-option label="Kelas 5" value="5" />
                        <el-option label="Kelas 6" value="6" />
                    </el-select>
                </el-col>
                <el-col :span="8">
                    <el-select
                        v-model="filterTahunAjaran"
                        placeholder="Filter Tahun Ajaran"
                        clearable
                        class="w-full"
                        @change="loadData"
                    >
                        <el-option
                            v-for="ta in tahunAjaranStore.list"
                            :key="ta.id"
                            :label="`${ta.nama} - ${ta.semester}${ta.is_active ? ' (Aktif)' : ''}`"
                            :value="ta.id"
                        />
                    </el-select>
                </el-col>
            </el-row>
        </el-card>

        <!-- Table -->
        <el-card shadow="hover">
            <el-table
                :data="pagedData"
                v-loading="kelasStore.loading"
                stripe
                style="width: 100%"
            >
                <el-table-column prop="nama" label="Nama Kelas" />
                <el-table-column prop="tingkat" label="Tingkat" width="120" />
                <el-table-column
                    prop="wali_kelas"
                    label="Wali Kelas"
                    min-width="150"
                />
                <el-table-column label="Tahun Ajaran" min-width="200">
                    <template #default="{ row }">
                        {{ row.tahun_ajaran?.nama }} -
                        {{ row.tahun_ajaran?.semester }}
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
                v-if="kelasStore.list.length > pageSize"
            >
                <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :total="kelasStore.list.length"
                    :page-sizes="[10, 20, 50]"
                    layout="total, sizes, prev, pager, next, jumper"
                    background
                />
            </div>
        </el-card>

        <!-- Dialog Form -->
        <el-dialog
            v-model="dialogVisible"
            :title="isEdit ? 'Edit Kelas' : 'Tambah Kelas'"
            width="500px"
        >
            <el-form
                ref="formRef"
                :model="form"
                :rules="rules"
                label-width="140px"
            >
                <el-form-item label="Nama Kelas" prop="nama">
                    <el-input
                        v-model="form.nama"
                        placeholder="Contoh: XII IPA 1"
                    />
                </el-form-item>
                <el-form-item label="Tingkat" prop="tingkat">
                    <el-select v-model="form.tingkat" class="w-full">
                        <el-option label="Kelas 1" value="1" />
                        <el-option label="Kelas 2" value="2" />
                        <el-option label="Kelas 3" value="3" />
                        <el-option label="Kelas 4" value="4" />
                        <el-option label="Kelas 5" value="5" />
                        <el-option label="Kelas 6" value="6" />
                    </el-select>
                </el-form-item>
                <el-form-item label="Tahun Ajaran" prop="tahun_ajaran_id">
                    <el-select v-model="form.tahun_ajaran_id" class="w-full">
                        <el-option
                            v-for="ta in tahunAjaranStore.list"
                            :key="ta.id"
                            :label="`${ta.nama} - ${ta.semester}`"
                            :value="ta.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="Wali Kelas" prop="wali_kelas">
                    <el-input
                        v-model="form.wali_kelas"
                        placeholder="Nama wali kelas"
                    />
                </el-form-item>
            </el-form>
            <template #footer>
                <el-button @click="dialogVisible = false">Batal</el-button>
                <el-button
                    type="primary"
                    :loading="submitting"
                    @click="handleSubmit"
                    >Simpan</el-button
                >
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import type { FormInstance, FormRules } from "element-plus";
import { useKelasStore } from "../stores/kelas";
import { useTahunAjaranStore } from "../stores/tahunAjaran";
import type { KelasForm } from "../types";

const kelasStore = useKelasStore();
const tahunAjaranStore = useTahunAjaranStore();

const dialogVisible = ref(false);
const isEdit = ref(false);
const editingId = ref<number | null>(null);
const submitting = ref(false);
const formRef = ref<FormInstance>();
const filterTingkat = ref("");
const filterTahunAjaran = ref<number>();

const currentPage = ref(1);
const pageSize = ref(10);
const pagedData = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    return kelasStore.list.slice(start, start + pageSize.value);
});

const form = reactive<KelasForm>({
    nama: "",
    tingkat: "1",
    tahun_ajaran_id: 0,
    wali_kelas: "",
});

const rules: FormRules = {
    nama: [
        { required: true, message: "Nama kelas harus diisi", trigger: "blur" },
    ],
    tingkat: [{ required: true, message: "Pilih tingkat", trigger: "change" }],
    tahun_ajaran_id: [
        {
            required: true,
            message: "Pilih tahun ajaran",
            trigger: "change",
            type: "number",
        },
    ],
    wali_kelas: [
        {
            required: true,
            message: "Nama wali kelas harus diisi",
            trigger: "blur",
        },
    ],
};

onMounted(() => {
    loadData();
    tahunAjaranStore.fetchAll();
});

function loadData() {
    currentPage.value = 1;
    kelasStore.fetchAll({
        tingkat: filterTingkat.value || undefined,
        tahun_ajaran_id: filterTahunAjaran.value || undefined,
    });
}

function openDialog(row?: KelasForm & { id?: number }) {
    if (row) {
        isEdit.value = true;
        editingId.value = row.id!;
        form.nama = row.nama;
        form.tingkat = row.tingkat;
        form.tahun_ajaran_id = row.tahun_ajaran_id;
        form.wali_kelas = row.wali_kelas;
    } else {
        isEdit.value = false;
        editingId.value = null;
        form.nama = "";
        form.tingkat = "1";
        form.tahun_ajaran_id =
            tahunAjaranStore.list.find((t) => t.is_active)?.id || 0;
        form.wali_kelas = "";
    }
    dialogVisible.value = true;
}

async function handleSubmit() {
    const valid = await formRef.value?.validate().catch(() => false);
    if (!valid) return;

    submitting.value = true;
    try {
        if (isEdit.value && editingId.value) {
            await kelasStore.update(editingId.value, form);
            ElMessage.success("Kelas berhasil diupdate");
        } else {
            await kelasStore.create(form);
            ElMessage.success("Kelas berhasil ditambahkan");
        }
        dialogVisible.value = false;
    } finally {
        submitting.value = false;
    }
}

async function handleDelete(id: number) {
    try {
        await ElMessageBox.confirm(
            "Yakin ingin menghapus kelas ini?",
            "Konfirmasi",
            {
                type: "warning",
            },
        );
        await kelasStore.remove(id);
        ElMessage.success("Kelas berhasil dihapus");
    } catch {
        // cancelled
    }
}
</script>
