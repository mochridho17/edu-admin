<template>
    <div>
        <el-card shadow="hover">
            <template #header>
                <div class="flex justify-between items-center">
                    <span>Daftar Tahun Ajaran</span>
                    <el-button type="primary" size="small" @click="openDialog()"
                        >+ Tambah Tahun Ajaran</el-button
                    >
                </div>
            </template>
            <el-table
                :data="pagedData"
                v-loading="tahunAjaranStore.loading"
                stripe
                style="width: 100%"
            >
                <el-table-column prop="nama" label="Tahun Ajaran" width="150" />
                <el-table-column prop="semester" label="Semester" width="100">
                    <template #default="{ row }">
                        <el-tag
                            :type="
                                row.semester === 'ganjil'
                                    ? 'primary'
                                    : 'warning'
                            "
                        >
                            {{ row.semester === "ganjil" ? "Ganjil" : "Genap" }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="Status" width="120">
                    <template #default="{ row }">
                        <el-tag v-if="row.is_active" type="success"
                            >Aktif</el-tag
                        >
                        <el-tag v-else type="info">Tidak Aktif</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="Aksi" width="280" fixed="right">
                    <template #default="{ row }">
                        <el-button size="small" @click="openDialog(row)"
                            >Edit</el-button
                        >
                        <el-button
                            size="small"
                            :type="row.is_active ? 'warning' : 'success'"
                            @click="handleActivate(row.id)"
                        >
                            {{ row.is_active ? "Nonaktifkan" : "Aktifkan" }}
                        </el-button>
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
                v-if="tahunAjaranStore.list.length > pageSize"
            >
                <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :total="tahunAjaranStore.list.length"
                    :page-sizes="[10, 20, 50]"
                    layout="total, sizes, prev, pager, next, jumper"
                    background
                />
            </div>
        </el-card>

        <el-dialog
            v-model="dialogVisible"
            :title="isEdit ? 'Edit Tahun Ajaran' : 'Tambah Tahun Ajaran'"
            width="500px"
        >
            <el-form
                ref="formRef"
                :model="form"
                :rules="rules"
                label-width="140px"
            >
                <el-form-item label="Tahun Ajaran" prop="nama">
                    <el-input
                        v-model="form.nama"
                        placeholder="Contoh: 2025/2026"
                    />
                </el-form-item>
                <el-form-item label="Semester" prop="semester">
                    <el-select v-model="form.semester" class="w-full">
                        <el-option label="Ganjil" value="ganjil" />
                        <el-option label="Genap" value="genap" />
                    </el-select>
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
import { useTahunAjaranStore } from "../stores/tahunAjaran";
import type { TahunAjaranForm } from "../types";

const tahunAjaranStore = useTahunAjaranStore();

const currentPage = ref(1);
const pageSize = ref(10);
const pagedData = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    return tahunAjaranStore.list.slice(start, start + pageSize.value);
});

const dialogVisible = ref(false);
const isEdit = ref(false);
const editingId = ref<number | null>(null);
const submitting = ref(false);
const formRef = ref<FormInstance>();

const form = reactive<TahunAjaranForm>({ nama: "", semester: "ganjil" });
const rules: FormRules = {
    nama: [
        {
            required: true,
            message: "Tahun ajaran harus diisi",
            trigger: "blur",
        },
    ],
    semester: [
        { required: true, message: "Pilih semester", trigger: "change" },
    ],
};

onMounted(() => {
    currentPage.value = 1;
    tahunAjaranStore.fetchAll();
});

function openDialog(row?: TahunAjaranForm & { id?: number }) {
    if (row) {
        isEdit.value = true;
        editingId.value = row.id!;
        form.nama = row.nama;
        form.semester = row.semester;
    } else {
        isEdit.value = false;
        editingId.value = null;
        form.nama = "";
        form.semester = "ganjil";
    }
    dialogVisible.value = true;
}

async function handleSubmit() {
    const valid = await formRef.value?.validate().catch(() => false);
    if (!valid) return;
    submitting.value = true;
    try {
        if (isEdit.value && editingId.value) {
            await tahunAjaranStore.update(editingId.value, form);
            ElMessage.success("Tahun ajaran berhasil diupdate");
        } else {
            await tahunAjaranStore.create(form);
            ElMessage.success("Tahun ajaran berhasil ditambahkan");
        }
        dialogVisible.value = false;
    } finally {
        submitting.value = false;
    }
}

async function handleActivate(id: number) {
    await tahunAjaranStore.activate(id);
    ElMessage.success("Status tahun ajaran berhasil diubah");
}

async function handleDelete(id: number) {
    try {
        await ElMessageBox.confirm(
            "Yakin ingin menghapus tahun ajaran ini?",
            "Konfirmasi",
            { type: "warning" },
        );
        await tahunAjaranStore.remove(id);
        ElMessage.success("Tahun ajaran berhasil dihapus");
    } catch {
        /* cancelled */
    }
}
</script>
