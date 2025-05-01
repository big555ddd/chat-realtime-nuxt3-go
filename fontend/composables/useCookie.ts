import useStatefulCookie from "./useStatefulCookie";

export function setCookie(name: string, options: any = { path: '/' }) {
   return useStatefulCookie<string | undefined>(name, options)
}

export function removeCookie(name: string) {
   const cookie = useCookie(name)
   cookie.value = null

   const state = useState<string | undefined>(name, () => undefined)
   watch(state, () => { cookie.value = null })

   return state
}