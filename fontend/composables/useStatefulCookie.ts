import type { CookieOptions } from "nuxt/app";

export default function useStatefulCookie<T>(key: string, _opts: CookieOptions<T | undefined>) {
   const cookie = useCookie<T | undefined>(key, { ..._opts, readonly: false })

   const state = useState<T | undefined>(key, () => cookie.value)

   watch(state, () => {
      cookie.value = state.value
   })

   return state
}



