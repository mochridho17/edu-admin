<template>
  <div>
    <el-card shadow="hover">
      <template #header>
        <div class="flex justify-between items-center">
          <span>Daftar Mata Pelajaran</span>
          <el-button type="primary" size="small" @click="openDialog()">+ Tambah Mapel</el-button>
        </div>
      </template>
      <el-table :data="mapelStore.list" v-loading="mapelStore.loading" stripe style="width: 100%">
        <el-table-column prop="kode" label="Kode" width="120" />
        <el-table-column prop="nama" label="Nama Mapel" min-width="200" />
        <el-table-column prop="kkm" label="KKM" width="100" />
        <el-table-column label="Aksi" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="openDialog(row)">Edit</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.id)">Hapus</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="isEdit ? 'Edit Mapel' : 'Tambah Mapel'" width="500px">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="Kode" prop="kode">
          <el-input v-model="form.kode" placeholder="Contoh: MTK-WAJIB" />
        </el-form-item>
        <el-form-item label="Nama" prop="nama">
          <el-input v-model="form.nama" placeholder="Nama mata pelajaran" />
        </el-form-item>
        <el-form-item label="KKM" prop="kkm">
          <el-input-number v-model="form.kkm" :min="0" :max="100" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">Batal</el-button>
        <el-button type="primary" :loading="submitting" @click="handleSubmit">Simpan</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useMapelStore } from '../stores/mapel'
import type { MataPelajaranForm } from '../types'

const mapelStore = useMapelStore()
const dialogVisible = ref(false)
const isEdit = ref(false)
const editingId = ref<number | null>(null)
const submitting = ref(false)
const formRef = ref<FormInstance>()

const form = reactive<MataPelajaranForm>({ kode: '', nama: '', kkm: 75 })
const rules: FormRules = {
  kode: [{ required: true, message: 'Kode mapel harus diisi', trigger: 'blur' }],
  nama: [{ required: true, message: 'Nama mapel harus diisi', trigger: 'blur' }],
  kkm: [{ required: true, message: 'KKM harus diisi', trigger: 'blur' }],
}

onMounted(() => mapelStore.fetchAll())

function openDialog(row?: MataPelajaranForm & { id?: number }) {
  if (row) {
    isEdit.value = true; editingId.value = row.id!
    form.kode = row.kode; form.nama = row.nama; form.kkm = row.kkm
  } else {
    isEdit.value = false; editingId.value = null
    form.kode = ''; form.nama = ''; form.kkm = 75
  }
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitting.value = true
  try {
    if (isEdit.value && editingId.value) {
      await mapelStore.update(editingId.value, form)
      ElMessage.success('Mapel berhasil diupdate')
    } else {
      await mapelStore.create(form)
      ElMessage.success('Mapel berhasil ditambahkan')
    }
    dialogVisible.value = false
  } finally { submitting.value = false }
}

async function handleDelete(id: number) {
  try {
    await ElMessageBox.confirm('Yakin ingin menghapus mapel ini?', 'Konfirmasi', { type: 'warning' })
    await mapelStore.remove(id)
    ElMessage.success('Mapel berhasil dihapus')
  } catch { /* cancelled */ }
}
</script>
