# PRD — Edu-Admin Dashboard

## 1. Ringkasan Proyek

**Edu-Admin** adalah aplikasi web admin dashboard untuk merekap absensi harian dan rekap nilai siswa. Digunakan oleh guru/wali kelas untuk memantau kehadiran dan performa akademik dalam satu tempat.

> **Keputusan Teknis**: Vue 3 + TypeScript (Frontend) · Gin + GORM + PostgreSQL (Backend) · Single-sekolah

---

## 2. Tujuan & Sasaran

- Mengganti rekap manual (Excel/kertas) dengan sistem digital
- Menyediakan dashboard rekap absensi harian yang cepat
- Menyediakan dashboard rekap nilai per siswa per mata pelajaran
- Mengekspor data ke format PDF/Excel untuk laporan

---

## 3. Target Pengguna

| Role | Deskripsi |
|---|---|
| **Admin / Guru** | Melihat & mengelola absensi dan nilai |
| **Super Admin** | Kelola user, kelas, mapel, dan data master |

---

## 4. Fitur (Functional Requirements)

### 4.1 Autentikasi & Manajemen User
- Login / Logout (JWT)
- Role-based access: Admin, Super Admin
- CRUD user (Super Admin only)
- Profil & ubah password

### 4.2 Data Master
- **Kelas**: CRUD kelas (nama kelas, tingkat, wali kelas)
- **Siswa**: CRUD siswa (NIS, nama, kelas, jenis kelamin)
- **Mata Pelajaran**: CRUD mapel (kode, nama, KKM)
- **Tahun Ajaran / Semester**: CRUD tahun ajaran & semester aktif

### 4.3 Absensi Harian
- **Input Absen**: Pilih kelas + tanggal, centang hadir/sakit/izin/alfa per siswa
- **Rekap Harian**: Tampilkan ringkasan per kelas per hari
- **Rekap Bulanan**: Grafik & tabel rekap per siswa per bulan
- **Filter**: Kelas, tanggal range, status kehadiran

### 4.4 Nilai
- **Input Nilai**: Pilih kelas + mapel, input nilai tugas, UTS, UAS
- **Rekap Nilai**: Tabel nilai per siswa per mapel
- **Rata-rata & Peringkat**: Otomatis hitung rata-rata dan peringkat kelas
- **KKM**: Tandai siswa yang belum tuntas (di bawah KKM)

### 4.5 Dashboard & Laporan
- **Home Dashboard**: Statistik (total siswa, rata-rata kehadiran, grafik)
- **Ekspor**: Export rekap absen & nilai ke PDF / Excel
- **Pencarian**: Cari siswa berdasarkan NIS atau nama

---

## 5. Non-Functional Requirements

| Parameter | Spesifikasi |
|---|---|
| **Performance** | Response time < 2 detik untuk halaman rekap |
| **Security** | JWT expiry, bcrypt password, CORS, input validation |
| **Reliability** | PostgreSQL backup, error handling graceful |
| **Usability** | Responsive (mobile-friendly), dark/light mode opsional |
| **Scalability** | Siap ditambahkan fitur baru (arsitektur modular) |

---

## 6. Tech Stack (Final)

### 6.1 Backend (Go + Gin + GORM)

| Komponen | Pilihan |
|---|---|
| **Bahasa** | Go 1.22+ |
| **HTTP Router** | **Gin** |
| **Database** | PostgreSQL 16 |
| **ORM** | **GORM** |
| **Migration** | golang-migrate |
| **Auth** | golang-jwt (JWT) + bcrypt |
| **Validator** | go-playground/validator |
| **API Docs** | swaggo/swagger (OpenAPI) |
| **Logging** | zerolog (JSON structured) |
| **Config** | viper (`.env`) |
| **Container** | Docker + docker-compose |

### 6.2 Frontend (Vue 3)

| Komponen | Pilihan |
|---|---|
| **Framework** | **Vue 3** + TypeScript |
| **Build Tool** | **Vite** |
| **Routing** | Vue Router |
| **State Management** | **Pinia** |
| **UI Library** | **Element Plus** + Tailwind CSS |
| **HTTP Client** | **Axios** (dengan interceptor JWT) |
| **Charting** | **vue-chartjs** |
| **Form & Validation** | **VeeValidate** + **Zod** |
| **Table** | Element Plus Table (built-in) |

### 6.3 DevOps

- Docker Compose (Go app + PostgreSQL + Frontend via Nginx)
- Git (GitHub)
- Environment variable via `.env`

---

## 7. Arsitektur Aplikasi

```
┌─────────────────────────────────────┐
│         Browser (Vue 3 SPA)         │
│  localhost:5173 (dev) / :80 (prod)  │
└─────────────┬───────────────────────┘
              │ HTTP (JSON)
              ▼
┌─────────────────────────────────────┐
│    Go API (Gin) — localhost:8080     │
│  - Middleware: Auth, CORS, Logger    │
│  - Handler → Service → Repository   │
└─────────────┬───────────────────────┘
              │ SQL
              ▼
┌─────────────────────────────────────┐
│       PostgreSQL — localhost:5432     │
│  Database: edu_admin                 │
└─────────────────────────────────────┘
```

### Struktur Layer Backend
```
backend/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/       # env config loader
│   ├── database/     # koneksi & migration
│   ├── middleware/    # auth, cors, logger
│   ├── model/        # GORM model structs
│   ├── handler/      # HTTP handler (gin context)
│   ├── service/      # business logic
│   ├── repository/   # DB queries via GORM
│   └── dto/          # request/response structs
├── router/
│   └── router.go     # route definitions
├── go.mod
└── Dockerfile
```

---

## 8. Data Model (Final — Single Sekolah)

> **Catatan**: Single-sekolah → tidak perlu field `sekolah_id`, semua data milik satu institusi.

```
Users
├── id (PK, UUID)
├── username (unique)
├── email (unique)
├── password (bcrypt hash)
├── role (enum: super_admin, admin)
├── created_at
└── updated_at

TahunAjaran
├── id (PK, serial)
├── nama (ex: "2025/2026")
├── semester (enum: ganjil, genap)
├── is_active (boolean)
├── created_at
└── updated_at

Kelas
├── id (PK, serial)
├── nama (ex: "XII IPA 1")
├── tingkat (enum: X, XI, XII)
├── tahun_ajaran_id (FK)
├── wali_kelas (string)
├── created_at
└── updated_at

Siswa
├── id (PK, serial)
├── nis (unique, string)
├── nama (string)
├── jenis_kelamin (enum: L, P)
├── kelas_id (FK)
├── created_at
└── updated_at

MataPelajaran
├── id (PK, serial)
├── kode (unique, ex: "MTK-WAJIB")
├── nama (ex: "Matematika Wajib")
├── kkm (int, ex: 75)
├── created_at
└── updated_at

Absensi
├── id (PK, serial)
├── siswa_id (FK)
├── tanggal (date)
├── status (enum: hadir, sakit, izin, alfa)
├── keterangan (text, nullable)
├── created_at
└── updated_at
│   UNIQUE(siswa_id, tanggal)

Nilai
├── id (PK, serial)
├── siswa_id (FK)
├── mapel_id (FK)
├── semester (enum: ganjil, genap)
├── tahun_ajaran_id (FK)
├── tugas (float, nullable)
├── uts (float, nullable)
├── uas (float, nullable)
├── akhir (float, computed)
├── created_at
└── updated_at
│   UNIQUE(siswa_id, mapel_id, semester, tahun_ajaran_id)
```

---

## 9. API Endpoints

### Auth
| Method | Endpoint | Deskripsi |
|---|---|---|
| POST | `/api/auth/login` | Login user |
| POST | `/api/auth/refresh` | Refresh token |

### Users (Super Admin only)
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/users` | List user |
| POST | `/api/users` | Create user |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Hapus user |

### Tahun Ajaran
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/tahun-ajaran` | List |
| POST | `/api/tahun-ajaran` | Create |
| PUT | `/api/tahun-ajaran/:id` | Update |
| DELETE | `/api/tahun-ajaran/:id` | Delete |
| PATCH | `/api/tahun-ajaran/:id/activate` | Set active |

### Kelas
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/kelas` | List (filter by tingkat & tahun_ajaran_id) |
| POST | `/api/kelas` | Create |
| PUT | `/api/kelas/:id` | Update |
| DELETE | `/api/kelas/:id` | Delete |

### Siswa
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/siswa` | List (filter kelas, search NIS/nama) |
| POST | `/api/siswa` | Create |
| PUT | `/api/siswa/:id` | Update |
| DELETE | `/api/siswa/:id` | Delete |

### Mata Pelajaran
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/mapel` | List |
| POST | `/api/mapel` | Create |
| PUT | `/api/mapel/:id` | Update |
| DELETE | `/api/mapel/:id` | Delete |

### Absensi
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/absensi?kelas_id=&tanggal=` | Lihat absen per kelas per hari |
| POST | `/api/absensi` | Input/update batch absen |
| GET | `/api/absensi/rekap/bulanan?kelas_id=&bulan=&tahun=` | Rekap bulanan |
| GET | `/api/absensi/rekap/harian?tanggal=` | Rekap seluruh kelas hari ini |

### Nilai
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/nilai?kelas_id=&mapel_id=&semester=&tahun_ajaran_id=` | Lihat nilai |
| POST | `/api/nilai/batch` | Input/update batch nilai |
| GET | `/api/nilai/rekap/siswa/:id` | Rekap nilai per siswa |
| GET | `/api/nilai/rekap/kelas?kelas_id=&semester=&tahun_ajaran_id=` | Ranking kelas |

### Dashboard & Laporan
| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/api/dashboard/stats` | Statistik dashboard |
| GET | `/api/laporan/absensi?kelas_id=&bulan=&tahun=` | Export absensi (xlsx) |
| GET | `/api/laporan/nilai?kelas_id=&mapel_id=` | Export nilai (xlsx)

---

## 10. Halaman Frontend (Routes)

| # | Path | Halaman | Auth |
|---|---|---|---|
| 1 | `/login` | Login | Public |
| 2 | `/` | Dashboard (statistik + grafik) | All |
| 3 | `/kelas` | Data Kelas (CRUD) | All |
| 4 | `/siswa` | Data Siswa (CRUD + filter) | All |
| 5 | `/mapel` | Data Mata Pelajaran (CRUD) | All |
| 6 | `/tahun-ajaran` | Tahun Ajaran & Semester | Super Admin |
| 7 | `/absensi` | Input Absen Harian | All |
| 8 | `/absensi/rekap` | Rekap Absensi (tabel + grafik) | All |
| 9 | `/nilai` | Input Nilai | All |
| 10 | `/nilai/rekap` | Rekap Nilai + Ranking | All |
| 11 | `/laporan` | Export Laporan | All |
| 12 | `/users` | Manajemen User | Super Admin |

### Struktur Folder Frontend
```
frontend/
├── src/
│   ├── assets/         # CSS, images
│   ├── components/     # reusable components
│   ├── layouts/        # AdminLayout, AuthLayout
│   ├── views/          # halaman per route
│   ├── router/         # route definitions + guard
│   ├── stores/         # Pinia stores
│   ├── api/            # Axios instance + endpoint modules
│   ├── composables/    # reusable Vue composables
│   ├── utils/          # helpers (format tanggal, dsb)
│   └── types/          # TypeScript interfaces
├── vite.config.ts
├── tailwind.config.js
├── Dockerfile
└── nginx.conf
```

---

## 11. Milestone (Estimasi)

| Fase | Durasi | Output |
|---|---|---|
| **Fase 1: Foundation** | 1 minggu | Setup project (Go + DB + Frontend), Auth, CRUD master data |
| **Fase 2: Absensi** | 1 minggu | Input absen, rekap harian & bulanan |
| **Fase 3: Nilai** | 1 minggu | Input nilai batch, rekap, KKM, peringkat kelas |
| **Fase 4: Dashboard & Laporan** | 1 minggu | Dashboard stats + grafik, export Excel |
| **Fase 5: Polish** | 3 hari | Responsive, error handling, testing manual |

---

## 12. Keputusan yang Telah Diambil

| # | Keputusan | Pilihan |
|---|---|---|
| 1 | **Frontend** | Vue 3 + TypeScript ✅ |
| 2 | **Backend Router** | Gin ✅ |
| 3 | **ORM** | GORM ✅ |
| 4 | **Multi/Single-sekolah** | Single-sekolah ✅ |
| 5 | **Export Prioritas** | Excel (.xlsx) ✅ — pakai library `excelize` |
| 6 | **Dark Mode (MVP)** | Tidak dulu, nanti setelah fitur inti selesai ✅ |
