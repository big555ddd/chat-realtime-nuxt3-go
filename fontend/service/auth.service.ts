import type { ChangePasswordResetRequest, LoginRequest, ResetRequest } from "@/models/auth.models";
import { client } from './httpClient';

export function login(body: LoginRequest) {
   return client("/api/v1/auth/login", {
      method: "POST",
      body: body,
   });
}


export function update(body?: any) {
    return client(``, {
       method: "PATCH",
       body: body,
    });
 }


 
export function logout() {
   return client("/auth/logout", {
      method: "POST",
   });
}

export function register(body: any) {
   return client("/api/v1/auth/register", {
      method: "POST",
      body: JSON.stringify(body),
   });
}

export function resetpassword(body: ResetRequest) {
   return client("/api/v1/auth/reset-password", {
      method: "POST",
      body: body,
   });
}

export function changepassword(body: ChangePasswordResetRequest) {
   return client("/api/v1/auth/change-password", {
      method: "POST",
      body: body,
   });
}