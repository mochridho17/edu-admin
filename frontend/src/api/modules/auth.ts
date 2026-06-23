import apiClient from '../client'
import type { ApiResponse, LoginRequest, LoginResponse, User, UserForm } from '../../types'

export const authApi = {
  login(data: LoginRequest) {
    return apiClient.post<ApiResponse<LoginResponse>>('/auth/login', data)
  },
  refresh() {
    return apiClient.post<ApiResponse<{ token: string }>>('/auth/refresh')
  },
}

export const userApi = {
  getAll() {
    return apiClient.get<ApiResponse<User[]>>('/users')
  },
  create(data: UserForm) {
    return apiClient.post<ApiResponse<User>>('/users', data)
  },
  update(id: number, data: Partial<UserForm>) {
    return apiClient.put<ApiResponse<User>>(`/users/${id}`, data)
  },
  delete(id: number) {
    return apiClient.delete<ApiResponse<null>>(`/users/${id}`)
  },
}
