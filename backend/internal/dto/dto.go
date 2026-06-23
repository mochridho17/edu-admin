package dto

// ====== Auth ======
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// ====== User ======
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
	Role     string `json:"role"`
}

// ====== Tahun Ajaran ======
type CreateTahunAjaranRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Semester string `json:"semester" binding:"required"`
}

type UpdateTahunAjaranRequest struct {
	Nama     string `json:"nama"`
	Semester string `json:"semester"`
}

// ====== Kelas ======
type CreateKelasRequest struct {
	Nama          string `json:"nama" binding:"required"`
	Tingkat       string `json:"tingkat" binding:"required"`
	TahunAjaranID uint   `json:"tahun_ajaran_id" binding:"required"`
	WaliKelas     string `json:"wali_kelas" binding:"required"`
}

type UpdateKelasRequest struct {
	Nama          string `json:"nama"`
	Tingkat       string `json:"tingkat"`
	TahunAjaranID uint   `json:"tahun_ajaran_id"`
	WaliKelas     string `json:"wali_kelas"`
}

// ====== Siswa ======
type CreateSiswaRequest struct {
	NIS          string `json:"nis" binding:"required"`
	Nama         string `json:"nama" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	KelasID      uint   `json:"kelas_id" binding:"required"`
}

type UpdateSiswaRequest struct {
	NIS          string `json:"nis"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	KelasID      uint   `json:"kelas_id"`
}

// ====== Mata Pelajaran ======
type CreateMapelRequest struct {
	Kode string `json:"kode" binding:"required"`
	Nama string `json:"nama" binding:"required"`
	KKM  int    `json:"kkm" binding:"required"`
}

type UpdateMapelRequest struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
	KKM  int    `json:"kkm"`
}

// ====== Absensi ======
type AbsensiItem struct {
	SiswaID    uint   `json:"siswa_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
	Keterangan string `json:"keterangan"`
}

type CreateAbsensiBatchRequest struct {
	KelasID uint          `json:"kelas_id" binding:"required"`
	Tanggal string        `json:"tanggal" binding:"required"`
	Data    []AbsensiItem `json:"data" binding:"required"`
}

type AbsensiRekapBulanan struct {
	SiswaID   uint   `json:"siswa_id"`
	NamaSiswa string `json:"nama_siswa"`
	NIS       string `json:"nis"`
	Hadir     int    `json:"hadir"`
	Sakit     int    `json:"sakit"`
	Izin      int    `json:"izin"`
	Alfa      int    `json:"alfa"`
	Total     int    `json:"total"`
}

// ====== Nilai ======
type NilaiItem struct {
	SiswaID uint     `json:"siswa_id" binding:"required"`
	Tugas   *float64 `json:"tugas"`
	UTS     *float64 `json:"uts"`
	UAS     *float64 `json:"uas"`
}

type CreateNilaiBatchRequest struct {
	KelasID       uint        `json:"kelas_id" binding:"required"`
	MapelID       uint        `json:"mapel_id" binding:"required"`
	Semester      string      `json:"semester" binding:"required"`
	TahunAjaranID uint        `json:"tahun_ajaran_id" binding:"required"`
	Data          []NilaiItem `json:"data" binding:"required"`
}

type NilaiRekapSiswa struct {
	MapelNama string   `json:"mapel_nama"`
	MapelKode string   `json:"mapel_kode"`
	KKM       int      `json:"kkm"`
	Tugas     *float64 `json:"tugas"`
	UTS       *float64 `json:"uts"`
	UAS       *float64 `json:"uas"`
	Akhir     *float64 `json:"akhir"`
	Tuntas    bool     `json:"tuntas"`
}

type NilaiRankingKelas struct {
	SiswaID   uint    `json:"siswa_id"`
	NamaSiswa string  `json:"nama_siswa"`
	NIS       string  `json:"nis"`
	RataRata  float64 `json:"rata_rata"`
	Peringkat int     `json:"peringkat"`
}

// ====== Dashboard ======
type DashboardStats struct {
	TotalSiswa        int64            `json:"total_siswa"`
	TotalKelas        int64            `json:"total_kelas"`
	TotalMapel        int64            `json:"total_mapel"`
	RataRataKehadiran string           `json:"rata_rata_kehadiran"`
	AbsensiHariIni    map[string]int64 `json:"absensi_hari_ini"`
	NamaSekolah       string           `json:"nama_sekolah"` // ← TAMBAH INI
}

// ====== Generic ======
type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(data interface{}) ApiResponse {
	return ApiResponse{Success: true, Message: "Berhasil", Data: data}
}

func Created(data interface{}) ApiResponse {
	return ApiResponse{Success: true, Message: "Berhasil dibuat", Data: data}
}

func Error(message string) ApiResponse {
	return ApiResponse{Success: false, Message: message, Data: nil}
}
