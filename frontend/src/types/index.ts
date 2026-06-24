// ====== Enums ======
export type Role = "super_admin" | "admin";
export type Semester = "ganjil" | "genap";
export type Tingkat = "1" | "2" | "3" | "4" | "5" | "6";
export type JenisKelamin = "L" | "P";
export type StatusAbsen = "hadir" | "sakit" | "izin" | "alfa";

// ====== Auth ======
export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface JwtPayload {
  user_id: number;
  role: Role;
  exp: number;
}

// ====== User ======
export interface User {
  id: number;
  username: string;
  email: string;
  role: Role;
  created_at: string;
  updated_at: string;
}

export interface UserForm {
  username: string;
  email: string;
  password?: string;
  role: Role;
}

// ====== Tahun Ajaran ======
export interface TahunAjaran {
  id: number;
  nama: string;
  semester: Semester;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface TahunAjaranForm {
  nama: string;
  semester: Semester;
}

// ====== Kelas ======
export interface Kelas {
  id: number;
  nama: string;
  tingkat: Tingkat;
  tahun_ajaran_id: number;
  wali_kelas: string;
  tahun_ajaran?: TahunAjaran;
  created_at: string;
  updated_at: string;
}

export interface KelasForm {
  nama: string;
  tingkat: Tingkat;
  tahun_ajaran_id: number;
  wali_kelas: string;
}

// ====== Siswa ======
export interface Siswa {
  id: number;
  nis: string;
  nama: string;
  jenis_kelamin: JenisKelamin;
  kelas_id: number;
  kelas?: Kelas;
  created_at: string;
  updated_at: string;
}

export interface SiswaForm {
  nis: string;
  nama: string;
  jenis_kelamin: JenisKelamin;
  kelas_id: number;
}

// ====== Mata Pelajaran ======
export interface MataPelajaran {
  id: number;
  kode: string;
  nama: string;
  kkm: number;
  created_at: string;
  updated_at: string;
}

export interface MataPelajaranForm {
  kode: string;
  nama: string;
  kkm: number;
}

// ====== Absensi ======
export interface Absensi {
  id: number;
  siswa_id: number;
  tanggal: string;
  status: StatusAbsen;
  keterangan: string | null;
  siswa?: Siswa;
  created_at: string;
  updated_at: string;
}

export interface AbsensiForm {
  siswa_id: number;
  tanggal: string;
  status: StatusAbsen;
  keterangan?: string;
}

export interface AbsensiBatchForm {
  kelas_id: number;
  tanggal: string;
  data: { siswa_id: number; status: StatusAbsen; keterangan?: string }[];
}

export interface AbsensiRekapBulanan {
  siswa_id: number;
  nama_siswa: string;
  nis: string;
  hadir: number;
  sakit: number;
  izin: number;
  alfa: number;
  total: number;
}

// ====== Nilai ======
export interface Nilai {
  id: number;
  siswa_id: number;
  mapel_id: number;
  semester: Semester;
  tahun_ajaran_id: number;
  tugas: number | null;
  uts: number | null;
  uas: number | null;
  akhir: number | null;
  siswa?: Siswa;
  mapel?: MataPelajaran;
  created_at: string;
  updated_at: string;
}

export interface NilaiBatchForm {
  kelas_id: number;
  mapel_id: number;
  semester: Semester;
  tahun_ajaran_id: number;
  data: { siswa_id: number; tugas?: number; uts?: number; uas?: number }[];
}

export interface NilaiRekapSiswa {
  mapel_nama: string;
  mapel_kode: string;
  kkm: number;
  tugas: number | null;
  uts: number | null;
  uas: number | null;
  akhir: number | null;
  tuntas: boolean;
}

export interface NilaiRankingKelas {
  siswa_id: number;
  nama_siswa: string;
  nis: string;
  rata_rata: number;
  peringkat: number;
}

// ====== Jurnal Guru ======
export interface JurnalGuru {
  id: number;
  tanggal: string;
  nama_guru: string;
  mapel_id: number;
  kelas_id: number;
  kegiatan: string;
  semester: Semester;
  tahun_ajaran_id: number;
  mapel?: MataPelajaran;
  kelas?: Kelas;
  tahun_ajaran?: TahunAjaran;
  created_at: string;
  updated_at: string;
}

export interface JurnalGuruForm {
  tanggal: string;
  nama_guru: string;
  mapel_id: number;
  kelas_id: number;
  kegiatan: string;
  semester: Semester;
  tahun_ajaran_id: number;
}

// ====== Dashboard ======
export interface DashboardStats {
  total_siswa: number;
  total_kelas: number;
  total_mapel: number;
  rata_rata_kehadiran: string;
  absensi_hari_ini: {
    hadir: number;
    sakit: number;
    izin: number;
    alfa: number;
  };
}

// ====== Pagination ======
export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  limit: number;
}

// ====== Api Response ======
export interface ApiResponse<T = unknown> {
  success: boolean;
  message: string;
  data: T;
}
