import apiClient from "../client";

export interface SchoolInfo {
  nama_sekolah: string;
  npsn: string;
  kepala_sekolah: string;
  jumlah_guru: string;
  akreditasi: string;
}

export const schoolApi = {
  getInfo() {
    return apiClient.get<{
      success: boolean;
      message: string;
      data: SchoolInfo;
    }>("/school/info");
  },
};
