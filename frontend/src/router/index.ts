import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    component: () => import('../layouts/AuthLayout.vue'),
    meta: { requiresAuth: false },
    children: [
      {
        path: '',
        name: 'Login',
        component: () => import('../views/LoginView.vue'),
        meta: { requiresAuth: false },
      },
    ],
  },
  {
    path: '/',
    component: () => import('../layouts/AdminLayout.vue'),
    meta: { layout: 'admin', requiresAuth: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('../views/DashboardView.vue'),
        meta: { title: 'Dashboard', icon: 'Odometer' },
      },
      {
        path: 'kelas',
        name: 'Kelas',
        component: () => import('../views/KelasView.vue'),
        meta: { title: 'Data Kelas', icon: 'School' },
      },
      {
        path: 'siswa',
        name: 'Siswa',
        component: () => import('../views/SiswaView.vue'),
        meta: { title: 'Data Siswa', icon: 'User' },
      },
      {
        path: 'mapel',
        name: 'Mapel',
        component: () => import('../views/MapelView.vue'),
        meta: { title: 'Mata Pelajaran', icon: 'Reading' },
      },
      {
        path: 'tahun-ajaran',
        name: 'TahunAjaran',
        component: () => import('../views/TahunAjaranView.vue'),
        meta: { title: 'Tahun Ajaran', icon: 'Calendar', roles: ['super_admin'] },
      },
      {
        path: 'absensi',
        name: 'Absensi',
        component: () => import('../views/AbsensiView.vue'),
        meta: { title: 'Input Absensi', icon: 'Edit' },
      },
      {
        path: 'absensi/rekap',
        name: 'AbsensiRekap',
        component: () => import('../views/AbsensiRekapView.vue'),
        meta: { title: 'Rekap Absensi', icon: 'DataAnalysis' },
      },
      {
        path: 'nilai',
        name: 'Nilai',
        component: () => import('../views/NilaiView.vue'),
        meta: { title: 'Input Nilai', icon: 'EditPen' },
      },
      {
        path: 'nilai/rekap',
        name: 'NilaiRekap',
        component: () => import('../views/NilaiRekapView.vue'),
        meta: { title: 'Rekap Nilai', icon: 'Trophy' },
      },
      {
        path: 'laporan',
        name: 'Laporan',
        component: () => import('../views/LaporanView.vue'),
        meta: { title: 'Laporan', icon: 'Document' },
      },
      {
        path: 'users',
        name: 'Users',
        component: () => import('../views/UsersView.vue'),
        meta: { title: 'Manajemen User', icon: 'UserFilled', roles: ['super_admin'] },
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Navigation guard
router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth === false) {
    next()
    return
  }

  if (!authStore.isLoggedIn) {
    // Coba restore dari localStorage
    authStore.restoreSession()
    if (!authStore.isLoggedIn) {
      next('/login')
      return
    }
  }

  // Cek role-based access
  const roles = to.meta.roles as string[] | undefined
  if (roles && !roles.includes(authStore.user?.role || '')) {
    next('/')
    return
  }

  next()
})

export default router
