<template>
  <div>
    <!-- Filter & Search -->
    <el-card shadow="hover" class="mb-4">
      <template #header>
        <div class="flex justify-between items-center">
          <span>Filter</span>
          <el-button type="primary" size="small" @click="openDialog()">+ Tambah Siswa</el-button>
        </div>
      </template>
      <el-row :gutter="16">
        <el-col :span="8">
          <el-select
            v-model="filterKelas"
            placeholder="Filter Kelas"
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
        <el-col :span="8">
          <el-input
            v-model="searchText"
            placeholder="Cari NIS / Nama..."
            clearable
            @input="loadData"
          />
        </el-col>
      </el-row>
    </el-card>

    <!-- Table -->
    <el-card shadow="hover">
      <el-table :data="siswaStore.list" v-loading="siswaStore.loading" stripe style="width: 100%">
        <el-table-column prop="nis" label="NIS" width="120" />
        <el-table-column prop="nama" label="Nama" min-width="180" />
        <el-table-column prop="jenis_kelamin" label="JK" width="80">
          <template #default="{ row }">
            {{ row.jenis_kelamin === 'L' ? 'Laki-laki' : 'Perempuan' }}
          </template>
        </el-table-column>
        <el-table-column label="Kelas" width="150">
          <template #default="{ row }">
            {{ row.kelas?.nama }}
          </template>
        </el-table-column>
        <el-table-column label="Aksi" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="openDialog(row)">Edit</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.id)">Hapus</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Dialog -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? 'Edit Siswa' : 'Tambah Siswa'"
      width="500px"
    >
      <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="NIS" prop="nis">
          <el-input v-model="form.nis" placeholder="Nomor Induk Siswa" />
        </el-form-item>
        <el-form-item label="Nama" prop="nama">
          <el-input v-model="form.nama" placeholder="Nama lengkap" />
        </el-form-item>
        <el-form-item label="Jenis Kelamin" prop="jenis_kelamin">
          <el-select v-model="form.jenis_kelamin" class="w-full">
            <el-option label="Laki-laki" value="L" />
            <el-option label="Perempuan" value="P" />
          </el-select>
        </el-form-item>
        <el-form-item label="Kelas" prop="kelas_id">
          <el-select v-model="form.kelas_id" class="w-full">
            <el-option
              v-for="k in kelasStore.list"
              :key="k.id"
              :label="k.nama"
              :value="k.id"
            />
          </el-select>
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
import { useSiswaStore } from '../stores/siswa'
import { useKelasStore } from '../stores/kelas'
import type { SiswaForm } from '../types'

const siswaStore = useSiswaStore()
const kelasStore = useKelasStore()

const dialogVisible = ref(false)
const isEdit = ref(false)
const editingId = ref<number | null>(null)
const submitting = ref(false)
const formRef = ref<FormInstance>()
const filterKelas = ref<number>()
const searchText = ref('')

const form = reactive<SiswaForm>({
  nis: '',
  nama: '',
  jenis_kelamin: 'L',
  kelas_id: 0,
})

const rules: FormRules = {
  nis: [{ required: true, message: 'NIS harus diisi', trigger: 'blur' }],
  nama: [{ required: true, message: 'Nama harus diisi', trigger: 'blur' }],
  jenis_kelamin: [{ required: true, message: 'Pilih jenis kelamin', trigger: 'change' }],
  kelas_id: [{ required: true, message: 'Pilih kelas', trigger: 'change', type: 'number' }],
}

onMounted(() => {
  loadData()
  kelasStore.fetchAll()
})

function loadData() {
  siswaStore.fetchAll({
    kelas_id: filterKelas.value || undefined,
    search: searchText.value || undefined,
  })
}

function openDialog(row?: SiswaForm & { id?: number }) {
  if (row) {
    isEdit.value = true
    editingId.value = row.id!
    form.nis = row.nis
    form.nama = row.nama
    form.jenis_kelamin = row.jenis_kelamin
    form.kelas_id = row.kelas_id
  } else {
    isEdit.value = false
    editingId.value = null
    form.nis = ''
    form.nama = ''
    form.jenis_kelamin = 'L'
    form.kelas_id = 0
  }
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitting.value = true
  try {
    if (isEdit.value && editingId.value) {
      await siswaStore.update(editingId.value, form)
      ElMessage.success('Siswa berhasil diupdate')
    } else {
      await siswaStore.create(form)
      ElMessage.success('Siswa berhasil ditambahkan')
    }
    dialogVisible.value = false
  } finally {
    submitting.value = false
  }
}

async function handleDelete(id: number) {
  try {
    await ElMessageBox.confirm('Yakin ingin menghapus siswa ini?', 'Konfirmasi', { type: 'warning' })
    await siswaStore.remove(id)
    ElMessage.success('Siswa berhasil dihapus')
  } catch { /* cancelled */ }
}
</script>
