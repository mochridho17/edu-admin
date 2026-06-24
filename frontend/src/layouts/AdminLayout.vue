<template>
    <div class="app-wrapper">
        <!-- Sidebar -->
        <aside class="sidebar" :class="{ collapsed: isCollapsed }">
            <!-- Logo -->
            <div class="sidebar-logo">
                <span class="logo-text" v-show="!isCollapsed">Edu-Admin</span>
            </div>

            <!-- Menu -->
            <el-scrollbar class="sidebar-scroll">
                <el-menu
                    :default-active="activeMenu"
                    router
                    :collapse="isCollapsed"
                    background-color="#1e1e2d"
                    text-color="#a2a3b7"
                    active-text-color="#ffffff"
                    class="sidebar-menu"
                >
                    <el-menu-item index="/">
                        <el-icon><Odometer /></el-icon>
                        <template #title>Dashboard</template>
                    </el-menu-item>

                    <el-sub-menu index="data-master">
                        <template #title>
                            <el-icon><FolderOpened /></el-icon>
                            <span>Data Master</span>
                        </template>
                        <el-menu-item index="/kelas">
                            <el-icon><School /></el-icon>
                            <template #title>Kelas</template>
                        </el-menu-item>
                        <el-menu-item index="/siswa">
                            <el-icon><User /></el-icon>
                            <template #title>Siswa</template>
                        </el-menu-item>
                        <el-menu-item index="/mapel">
                            <el-icon><Reading /></el-icon>
                            <template #title>Mata Pelajaran</template>
                        </el-menu-item>
                        <el-menu-item
                            v-if="authStore.isSuperAdmin"
                            index="/tahun-ajaran"
                        >
                            <el-icon><Calendar /></el-icon>
                            <template #title>Tahun Ajaran</template>
                        </el-menu-item>
                    </el-sub-menu>

                    <el-sub-menu index="absensi-menu">
                        <template #title>
                            <el-icon><Edit /></el-icon>
                            <span>Absensi</span>
                        </template>
                        <el-menu-item index="/absensi">
                            <el-icon><EditPen /></el-icon>
                            <template #title>Input Absensi</template>
                        </el-menu-item>
                        <el-menu-item index="/absensi/rekap">
                            <el-icon><DataAnalysis /></el-icon>
                            <template #title>Rekap Absensi</template>
                        </el-menu-item>
                    </el-sub-menu>

                    <el-sub-menu index="nilai-menu">
                        <template #title>
                            <el-icon><Trophy /></el-icon>
                            <span>Nilai</span>
                        </template>
                        <el-menu-item index="/nilai">
                            <el-icon><EditPen /></el-icon>
                            <template #title>Input Nilai</template>
                        </el-menu-item>
                        <el-menu-item index="/nilai/rekap">
                            <el-icon><DataBoard /></el-icon>
                            <template #title>Rekap Nilai</template>
                        </el-menu-item>
                    </el-sub-menu>

                    <el-menu-item index="/jurnal-guru">
                        <el-icon><Notebook /></el-icon>
                        <template #title>Jurnal Guru</template>
                    </el-menu-item>

                    <el-menu-item index="/laporan">
                        <el-icon><Document /></el-icon>
                        <template #title>Laporan</template>
                    </el-menu-item>

                    <el-menu-item v-if="authStore.isSuperAdmin" index="/users">
                        <el-icon><UserFilled /></el-icon>
                        <template #title>Manajemen User</template>
                    </el-menu-item>

                    <!-- Bottom section -->
                    <div class="sidebar-bottom" v-show="!isCollapsed">
                        <el-divider class="sidebar-divider" />
                        <div class="sidebar-user-info">
                            <el-avatar :size="32" class="sidebar-avatar">
                                {{
                                    authStore.user?.username
                                        ?.charAt(0)
                                        .toUpperCase()
                                }}
                            </el-avatar>
                            <div class="user-details">
                                <span class="user-name">{{
                                    authStore.user?.username
                                }}</span>
                                <span class="user-role">{{
                                    authStore.isSuperAdmin
                                        ? "Super Admin"
                                        : "Admin"
                                }}</span>
                            </div>
                        </div>
                    </div>
                </el-menu>
            </el-scrollbar>
        </aside>

        <!-- Main Content -->
        <div class="main-container" :class="{ expanded: isCollapsed }">
            <!-- Topbar -->
            <header class="topbar">
                <div class="topbar-left">
                    <el-button
                        @click="isCollapsed = !isCollapsed"
                        text
                        class="toggle-btn"
                    >
                        <el-icon :size="20"
                            ><Fold v-if="!isCollapsed" /><Expand v-else
                        /></el-icon>
                    </el-button>
                    <el-breadcrumb separator="/">
                        <el-breadcrumb-item :to="{ path: '/' }"
                            >Home</el-breadcrumb-item
                        >
                        <el-breadcrumb-item v-if="pageTitle !== 'Dashboard'">{{
                            pageTitle
                        }}</el-breadcrumb-item>
                    </el-breadcrumb>
                </div>
                <div class="topbar-right">
                    <el-tag
                        v-if="authStore.isDemoMode"
                        type="warning"
                        size="small"
                        effect="dark"
                        class="demo-badge"
                    >
                        🧪 DEMO
                    </el-tag>
                    <el-dropdown trigger="click">
                        <span class="user-dropdown">
                            <el-avatar :size="32" class="user-avatar">
                                {{
                                    authStore.user?.username
                                        ?.charAt(0)
                                        .toUpperCase()
                                }}
                            </el-avatar>
                            <span class="user-name-top">{{
                                authStore.user?.username
                            }}</span>
                            <el-icon><ArrowDown /></el-icon>
                        </span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item @click="handleLogout">
                                    <el-icon><SwitchButton /></el-icon>
                                    Logout
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </div>
            </header>

            <!-- Page Title -->
            <div class="page-header">
                <h2 class="page-title">{{ pageTitle }}</h2>
                <p class="page-description" v-if="pageDescriptions[pageTitle]">
                    {{ pageDescriptions[pageTitle] }}
                </p>
            </div>

            <!-- Page Content -->
            <main class="page-content">
                <router-view />
            </main>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router";
import { useAuthStore } from "../stores/auth";
import {
    Odometer,
    FolderOpened,
    School,
    User,
    Reading,
    Calendar,
    Edit,
    EditPen,
    DataAnalysis,
    Trophy,
    DataBoard,
    Document,
    Notebook,
    UserFilled,
    Fold,
    Expand,
    ArrowDown,
    SwitchButton,
} from "@element-plus/icons-vue";

const route = useRoute();
const authStore = useAuthStore();

const isCollapsed = ref(false);
const activeMenu = computed(() => route.path);
const pageTitle = computed(() => (route.meta.title as string) || "Dashboard");

const pageDescriptions: Record<string, string> = {
    "Data Kelas": "Kelola data kelas dan wali kelas",
    "Data Siswa": "Kelola data siswa per kelas",
    "Mata Pelajaran": "Kelola mata pelajaran dan KKM",
    "Tahun Ajaran": "Atur tahun ajaran dan semester aktif",
    "Input Absensi": "Catat kehadiran siswa harian",
    "Rekap Absensi": "Lihat rekap kehadiran bulanan",
    "Input Nilai": "Input nilai tugas, UTS, dan UAS",
    "Rekap Nilai": "Lihat peringkat dan rekap nilai siswa",
    "Jurnal Guru": "Catat dan lihat jurnal harian kegiatan pembelajaran",
    Laporan: "Export laporan absensi dan nilai",
    "Manajemen User": "Kelola akun pengguna sistem",
};

function handleLogout() {
    authStore.logout();
}
</script>

<style scoped>
.app-wrapper {
    display: flex;
    height: 100vh;
    background: #f0f2f5;
}

/* ====== SIDEBAR ====== */
.sidebar {
    width: 260px;
    min-width: 260px;
    background: #1e1e2d;
    display: flex;
    flex-direction: column;
    transition: all 0.3s ease;
    z-index: 100;
}
.sidebar.collapsed {
    width: 64px;
    min-width: 64px;
}

.sidebar-logo {
    height: 64px;
    display: flex;
    align-items: center;
    padding: 0 20px;
    gap: 10px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}
.sidebar-logo .logo-icon {
    font-size: 24px;
}
.sidebar-logo .logo-text {
    font-size: 18px;
    font-weight: 700;
    color: #ffffff;
    white-space: nowrap;
}

.sidebar-scroll {
    flex: 1;
    overflow: hidden;
}
.sidebar-scroll :deep(.el-scrollbar__view) {
    height: 100%;
}

.sidebar-menu {
    border-right: none !important;
    height: 100%;
}
.sidebar-menu:not(.el-menu--collapse) {
    width: 260px;
}
.sidebar-menu.el-menu--collapse {
    width: 64px;
}
.sidebar-menu :deep(.el-menu-item),
.sidebar-menu :deep(.el-sub-menu__title) {
    height: 46px;
    line-height: 46px;
}
.sidebar-menu :deep(.el-menu-item.is-active) {
    background: linear-gradient(90deg, #009ef7, #009ef7dd) !important;
    color: #fff !important;
}
.sidebar-menu :deep(.el-menu-item:hover),
.sidebar-menu :deep(.el-sub-menu__title:hover) {
    background-color: rgba(255, 255, 255, 0.05) !important;
}
.sidebar-menu :deep(.el-sub-menu .el-menu) {
    background-color: #1a1a28 !important;
}
.sidebar-menu :deep(.el-menu-item .el-icon),
.sidebar-menu :deep(.el-sub-menu__title .el-icon) {
    color: #a2a3b7;
}
.sidebar-menu :deep(.el-menu-item.is-active .el-icon) {
    color: #fff;
}

.sidebar-bottom {
    padding: 8px 16px;
}
.sidebar-divider {
    margin: 4px 0;
    border-color: rgba(255, 255, 255, 0.08);
}
.sidebar-user-info {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 8px 4px;
}
.sidebar-avatar {
    background: #009ef7;
    color: #fff;
    flex-shrink: 0;
}
.user-details {
    display: flex;
    flex-direction: column;
    overflow: hidden;
}
.user-name {
    font-size: 13px;
    font-weight: 600;
    color: #e0e0e0;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
.user-role {
    font-size: 11px;
    color: #a2a3b7;
}

/* ====== MAIN ====== */
.main-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    transition: margin-left 0.3s;
    min-width: 0;
}

/* ====== TOPBAR ====== */
.topbar {
    height: 64px;
    background: #fff;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 24px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.06);
    position: sticky;
    top: 0;
    z-index: 10;
}
.topbar-left {
    display: flex;
    align-items: center;
    gap: 16px;
}
.toggle-btn {
    font-size: 18px;
    color: #6c757d;
}
.topbar :deep(.el-breadcrumb) {
    font-size: 14px;
}
.topbar :deep(.el-breadcrumb__inner) {
    color: #6c757d !important;
}
.topbar :deep(.el-breadcrumb__inner.is-link) {
    color: #009ef7 !important;
    font-weight: 500;
}
.topbar-right {
    display: flex;
    align-items: center;
    gap: 16px;
}
.demo-badge {
    font-size: 11px;
    letter-spacing: 0.5px;
}
.user-dropdown {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    padding: 4px 8px;
    border-radius: 6px;
    transition: background 0.2s;
}
.user-dropdown:hover {
    background: #f0f2f5;
}
.user-avatar {
    background: #009ef7;
    color: #fff;
}
.user-name-top {
    font-size: 14px;
    font-weight: 500;
    color: #333;
}

/* ====== PAGE HEADER ====== */
.page-header {
    padding: 20px 24px 0;
}
.page-title {
    font-size: 22px;
    font-weight: 600;
    color: #1e1e2d;
    margin: 0;
}
.page-description {
    font-size: 13px;
    color: #a2a3b7;
    margin: 4px 0 0;
}

/* ====== PAGE CONTENT ====== */
.page-content {
    flex: 1;
    overflow: auto;
    padding: 20px 24px 24px;
}
</style>
