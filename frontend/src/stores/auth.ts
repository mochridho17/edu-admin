import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User, LoginRequest } from '../types'
import { authApi } from '../api/modules/auth'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const loading = ref(false)
  const isDemoMode = ref(false)

  const isLoggedIn = computed(() => !!token.value && !!user.value)
  const isSuperAdmin = computed(() => user.value?.role === 'super_admin')

  async function login(data: LoginRequest) {
    loading.value = true
    try {
      const res = await authApi.login(data)
      const { token: t, user: u } = res.data.data
      token.value = t
      user.value = u
      localStorage.setItem('token', t)
      localStorage.setItem('user', JSON.stringify(u))
      isDemoMode.value = false
      return true
    } catch (err: unknown) {
      // Jika koneksi ditolak (backend belum jalan), fallback ke demo mode
      if (err instanceof Error && (err.message?.includes('Network Error') || err.message?.includes('ERR_CONNECTION'))) {
        return demoLogin(data)
      }
      return false
    } finally {
      loading.value = false
    }
  }

  function demoLogin(data: LoginRequest) {
    // Demo mode - accept any username/password
    const demoUser: User = {
      id: 1,
      username: data.username || 'admin',
      email: 'admin@sekolah.sch.id',
      role: 'super_admin',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    }
    const demoToken = 'demo-token-edu-admin-2025'

    token.value = demoToken
    user.value = demoUser
    isDemoMode.value = true

    localStorage.setItem('token', demoToken)
    localStorage.setItem('user', JSON.stringify(demoUser))
    localStorage.setItem('demo_mode', 'true')

    return true
  }

  function logout() {
    token.value = null
    user.value = null
    isDemoMode.value = false
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    localStorage.removeItem('demo_mode')
    window.location.href = '/login'
  }

  function restoreSession() {
    const savedToken = localStorage.getItem('token')
    const savedUser = localStorage.getItem('user')
    if (savedToken && savedUser) {
      token.value = savedToken
      user.value = JSON.parse(savedUser)
      isDemoMode.value = localStorage.getItem('demo_mode') === 'true'
    }
  }

  return {
    user,
    token,
    loading,
    isLoggedIn,
    isSuperAdmin,
    isDemoMode,
    login,
    logout,
    restoreSession,
  }
})
