import Reset from "~/pages/reset.vue"

export interface LoginRequest {
    username: string
    password: string
  }

  export interface ResetRequest {
    redirect_url: string
    email: string
  }

export interface ChangePasswordResetRequest {
    password: string
    token: string
  }