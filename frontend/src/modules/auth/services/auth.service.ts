import { apiClient } from '@/core/config/axios-instance.config'
import type { Users } from '@/modules/users/models/users.model'
import type { EmployeeLoginDTO, UserLoginDTO } from '../models/auth.model'

class AuthService {
  private readonly baseUrl = '/auth'

  async loginEmployee(dto: EmployeeLoginDTO) {
    try {
      return apiClient.post(`${this.baseUrl}/employees/login`, dto).then((res) => res.data)
    } catch (error) {
      console.error(`Failed to login employee:`, error)
      throw error
    }
  }

  async logoutEmployee() {
    try {
      return apiClient.post(`${this.baseUrl}/employees/logout`).then((res) => res.data)
    } catch (error) {
      console.error(`Failed to logout employee:`, error)
      throw error
    }
  }

  async loginUser(dto: UserLoginDTO) {
    try {
      return apiClient.post(`${this.baseUrl}/users/login`, dto).then((res) => res.data)
    } catch (error) {
      console.error(`Failed to login user:`, error)
      throw error
    }
  }

  async getCurrentUser() {
    try {
      return apiClient.get<Users>(`${this.baseUrl}/users/me`).then((res) => res.data)
    } catch (error) {
      console.error(`Failed to get current user:`, error)
      throw error
    }
  }

  async logoutUser() {
    try {
      return apiClient.post(`${this.baseUrl}/users/logout`).then((res) => res.data)
    } catch (error) {
      console.error(`Failed to logout user:`, error)
      throw error
    }
  }
}

export const authService = new AuthService()
