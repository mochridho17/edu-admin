<template>
  <div class="w-full max-w-md">
    <el-card shadow="always" class="p-6">
      <div class="text-center mb-6">
        <h1 class="text-2xl font-bold text-gray-800 mb-2">Edu-Admin</h1>
        <p class="text-gray-500">Admin Dashboard</p>
      </div>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-position="top"
        @submit.prevent="handleLogin"
      >
        <el-form-item label="Username" prop="username">
          <el-input
            v-model="form.username"
            placeholder="Masukkan username"
            :prefix-icon="User"
            size="large"
          />
        </el-form-item>

        <el-form-item label="Password" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="Masukkan password"
            :prefix-icon="Lock"
            size="large"
            show-password
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="w-full"
            :loading="authStore.loading"
            native-type="submit"
          >
            Masuk
          </el-button>
        </el-form-item>
      </el-form>

      <!-- Demo mode hint - hanya muncul saat backend tidak terhubung -->
      <el-alert
        v-if="authStore.isDemoMode"
        title="🧪 Mode Demo — Data bersifat contoh"
        type="warning"
        :closable="false"
        show-icon
        class="mt-4"
      />

      <!-- Login info untuk user - tampil terus -->
      <div class="mt-4 text-center">
        <p class="text-xs text-gray-400">
          Login dengan akun yang sudah didaftarkan oleh admin
        </p>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { User, Lock } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const formRef = ref<FormInstance>()

const form = reactive({
  username: '',
  password: '',
})

const rules: FormRules = {
  username: [{ required: true, message: 'Username harus diisi', trigger: 'blur' }],
  password: [{ required: true, message: 'Password harus diisi', trigger: 'blur' }],
}

async function handleLogin() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  const success = await authStore.login(form)
  if (success) {
    ElMessage.success('Login berhasil')
    router.push('/')
  }
}
</script>
