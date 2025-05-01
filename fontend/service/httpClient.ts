import { $fetch, type FetchOptions } from "ofetch";
import { setCookie } from "~/composables/useCookie";
import type { Response } from "~/models/http.models";

export function client<T, R extends Response<T> = Response<T>>(url: string, opts?: FetchOptions<"json">) {
   // const config = useRuntimeConfig()
   const { $router } = useNuxtApp()

   const controller = new AbortController()
   setTimeout(() => {
      controller.abort()
   }, 60000);

   const token = setCookie('token')
   const headers: any = {
      Authorization: token.value && `Bearer ${token.value}`,
      'ngrok-skip-browser-warning': 'true',
      ...opts?.headers
   }

   const config = useRuntimeConfig() as unknown as { public: { NUXT_PUBLIC_BASE_URL: string } };


   return $fetch<R>(url, {
      baseURL: config?.public?.NUXT_PUBLIC_BASE_URL || "http://localhost:3000",
      signal: controller.signal,
      method: opts?.method?.toUpperCase(),
      headers,
      ...opts,
      onResponseError({ response }) {
         if (response?.status == 403 || response?.status === 401) {
            token.value = undefined
            setTimeout(() => {
               $router.push('/login')
            }, 500);
         }
      }
   })
}










// // เชื่อมต่อ SERVICES API URL

// import type { option } from '~/models/axios.model'
// import axios from 'axios'

// export const client = (ops: option) => {
//   const config: any = useRuntimeConfig()

//   let api: any = axios.create()

//   api.defaults.baseURL = config.public.API_BASE  // config API base URL for file .env
//   api.defaults.headers.common['Content-Type'] = 'application/json'

//   // ตรวจสอบ Token ที่ใช้ใน Cookie ก่อนทำการ request ไปยัง API
//   api.interceptors.request.use(
//     async function (config: any) {
//       let token = await useCookie('token')
//       config.headers['Authorization'] = 'Bearer ' + token.value
//       return config
//     },
//     function (error: any) {
//       return Promise.reject(error)
//     },
//   )
  
//   // ตรวจสอบ Status Code ที่ได้รับจาก API ก่อนทำการจัดการ response
//   api.defaults.validaStatus = false
//   ops.method = ops.method.toLowerCase()
  
//   return api({
//     ...ops,
//   })
// }