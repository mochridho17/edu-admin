/**
 * Format tanggal ke string Indonesia
 */
export function formatDate(date: string | Date): string {
  const d = new Date(date)
  return d.toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  })
}

/**
 * Format tanggal ke format YYYY-MM-DD (untuk input type=date)
 */
export function toDateInput(date: string | Date): string {
  const d = new Date(date)
  return d.toISOString().split('T')[0]
}

/**
 * Format angka ke format Indonesia (ribuan pakai titik)
 */
export function formatNumber(n: number): string {
  return n.toLocaleString('id-ID')
}

/**
 * Bulan dalam bahasa Indonesia
 */
export const BULAN_INDONESIA = [
  'Januari', 'Februari', 'Maret', 'April', 'Mei', 'Juni',
  'Juli', 'Agustus', 'September', 'Oktober', 'November', 'Desember',
]

/**
 * Dapatkan nama bulan dari nomor (1-based)
 */
export function getNamaBulan(bulan: number): string {
  return BULAN_INDONESIA[bulan - 1] || ''
}

/**
 * Tahun ajaran dari tahun saat ini
 */
export function getTahunAjaranOptions(): string[] {
  const tahun = new Date().getFullYear()
  return [
    `${tahun - 1}/${tahun}`,
    `${tahun}/${tahun + 1}`,
  ]
}

/**
 * Singleton variants untuk Element Plus
 */
export const STATUS_ABSEN_MAP: Record<string, string> = {
  hadir: 'success',
  sakit: 'warning',
  izin: 'info',
  alfa: 'danger',
}

export const STATUS_ABSEN_LABEL: Record<string, string> = {
  hadir: 'Hadir',
  sakit: 'Sakit',
  izin: 'Izin',
  alfa: 'Alfa',
}

export const ROLE_LABEL: Record<string, string> = {
  super_admin: 'Super Admin',
  admin: 'Admin',
}

export const TINGKAT_LABEL: Record<string, string> = {
  '1': 'Kelas 1',
  '2': 'Kelas 2',
  '3': 'Kelas 3',
  '4': 'Kelas 4',
  '5': 'Kelas 5',
  '6': 'Kelas 6',
}

export const SEMESTER_LABEL: Record<string, string> = {
  ganjil: 'Ganjil',
  genap: 'Genap',
}
