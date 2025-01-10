import axios from 'axios'

const apiUrl = `${import.meta.env.VITE_API_URL || 'http://localhost:8000'}/api/v1`

export const apiClient = axios.create({
  baseURL: apiUrl,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
})

apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('AUTH_TOKEN')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response && error.response.status === 401) {
      console.error('Unauthorized access - redirecting to login')
    }
    return Promise.reject(error)
  },
)
