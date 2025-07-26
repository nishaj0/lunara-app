import api from '@/lib/api';
import { RegisterRequest, LoginRequest, RegisterResponse, AuthResponse } from '@/types/auth';

export const authService = {
  // Register new user
  register: async (data: RegisterRequest): Promise<RegisterResponse> => {
    const response = await api.post('/auth/register', data);
    return response.data;
  },

  // Login user
  login: async (data: LoginRequest): Promise<AuthResponse> => {
    const response = await api.post('/auth/login', data);
    
    // Store token in localStorage
    if (response.data.token) {
      localStorage.setItem('token', response.data.token);
    }
    
    // return in expected format
    return {
      token: response.data.token,
      user: response.data.user
    };
  },

  // Logout user
  logout: () => {
    localStorage.removeItem('token');
  },

  // Check if user is authenticated
  isAuthenticated: (): boolean => {
    return !!localStorage.getItem('token');
  },

  // Get stored token
  getToken: (): string | null => {
    return localStorage.getItem('token');
  }
};