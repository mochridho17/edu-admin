<template>
    <div>
        <el-card shadow="hover">
            <template #header>
                <div class="flex justify-between items-center">
                    <span>Daftar User</span>
                    <el-button type="primary" size="small" @click="openDialog()"
                        >+ Tambah User</el-button
                    >
                </div>
            </template>
            <el-table
                :data="pagedData"
                v-loading="loading"
                stripe
                style="width: 100%"
            >
                <el-table-column
                    prop="username"
                    label="Username"
                    min-width="150"
                />
                <el-table-column prop="email" label="Email" min-width="200" />
                <el-table-column label="Role" width="150">
                    <template #default="{ row }">
                        <el-tag
                            :type="
                                row.role === 'super_admin'
                                    ? 'danger'
                                    : 'primary'
                            "
                        >
                            {{
                                row.role === "super_admin"
                                    ? "Super Admin"
                                    : "Admin"
                            }}
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
                            :disabled="row.id === authStore.user?.id"
                            @click="handleDelete(row.id)"
                        >
                            Hapus
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div
                class="flex justify-center mt-4"
                v-if="users.length > pageSize"
            >
                <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :total="users.length"
                    :page-sizes="[10, 20, 50]"
                    layout="total, sizes, prev, pager, next, jumper"
                    background
                />
            </div>
        </el-card>

        <el-dialog
            v-model="dialogVisible"
            :title="isEdit ? 'Edit User' : 'Tambah User'"
            width="500px"
        >
            <el-form
                ref="formRef"
                :model="form"
                :rules="rules"
                label-width="120px"
            >
                <el-form-item label="Username" prop="username">
                    <el-input v-model="form.username" />
                </el-form-item>
                <el-form-item label="Email" prop="email">
                    <el-input v-model="form.email" type="email" />
                </el-form-item>
                <el-form-item label="Password" prop="password">
                    <el-input
                        v-model="form.password"
                        type="password"
                        show-password
                        :placeholder="
                            isEdit ? 'Kosongkan jika tidak diubah' : 'Password'
                        "
                    />
                </el-form-item>
                <el-form-item label="Role" prop="role">
                    <el-select v-model="form.role" class="w-full">
                        <el-option label="Admin" value="admin" />
                        <el-option label="Super Admin" value="super_admin" />
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
import { useAuthStore } from "../stores/auth";
import { userApi } from "../api/modules/auth";
import type { User, UserForm } from "../types";

const authStore = useAuthStore();
const users = ref<User[]>([]);
const loading = ref(false);

const currentPage = ref(1);
const pageSize = ref(10);
const pagedData = computed(() => {
    const start = (currentPage.value - 1) * pageSize.value;
    return users.value.slice(start, start + pageSize.value);
});

const dialogVisible = ref(false);
const isEdit = ref(false);
const editingId = ref<number | null>(null);
const submitting = ref(false);
const formRef = ref<FormInstance>();

const form = reactive<UserForm & { password: string }>({
    username: "",
    email: "",
    password: "",
    role: "admin",
});

const rules: FormRules = {
    username: [
        { required: true, message: "Username harus diisi", trigger: "blur" },
    ],
    email: [
        { required: true, message: "Email harus diisi", trigger: "blur" },
        { type: "email", message: "Email tidak valid", trigger: "blur" },
    ],
    password: [
        { required: true, message: "Password harus diisi", trigger: "blur" },
        { min: 6, message: "Minimal 6 karakter", trigger: "blur" },
    ],
    role: [{ required: true, message: "Pilih role", trigger: "change" }],
};

onMounted(() => {
    currentPage.value = 1;
    fetchUsers();
});

async function fetchUsers() {
    loading.value = true;
    try {
        const res = await userApi.getAll();
        users.value = res.data.data;
    } finally {
        loading.value = false;
    }
}

function openDialog(row?: User) {
    if (row) {
        isEdit.value = true;
        editingId.value = row.id;
        form.username = row.username;
        form.email = row.email;
        form.password = "";
        form.role = row.role;
        // Password tidak wajib saat edit
        rules.password = [];
    } else {
        isEdit.value = false;
        editingId.value = null;
        form.username = "";
        form.email = "";
        form.password = "";
        form.role = "admin";
        rules.password = [
            {
                required: true,
                message: "Password harus diisi",
                trigger: "blur",
            },
        ];
    }
    dialogVisible.value = true;
}

async function handleSubmit() {
    const valid = await formRef.value?.validate().catch(() => false);
    if (!valid) return;
    submitting.value = true;
    try {
        if (isEdit.value && editingId.value) {
            const payload: Partial<UserForm> = {
                username: form.username,
                email: form.email,
                role: form.role,
            };
            if (form.password) payload.password = form.password;
            await userApi.update(editingId.value, payload);
            ElMessage.success("User berhasil diupdate");
        } else {
            await userApi.create(form);
            ElMessage.success("User berhasil ditambahkan");
        }
        dialogVisible.value = false;
        await fetchUsers();
    } finally {
        submitting.value = false;
    }
}

async function handleDelete(id: number) {
    try {
        await ElMessageBox.confirm(
            "Yakin ingin menghapus user ini?",
            "Konfirmasi",
            { type: "warning" },
        );
        await userApi.delete(id);
        ElMessage.success("User berhasil dihapus");
        await fetchUsers();
    } catch {
        /* cancelled */
    }
}
</script>
