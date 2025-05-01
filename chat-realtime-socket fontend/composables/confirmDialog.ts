import Swal, { type SweetAlertIcon } from 'sweetalert2'

import { logout } from '@/service/auth.service'

interface Toast {
  icon: string
  title: string
  text: string
  confirmButtonText?: string
  cancelButtonText?: string
}

export const swalConfirmDialog = (payload: Toast): Promise<boolean> => {
  return Swal.fire({
    icon: payload.icon as SweetAlertIcon,
    title: payload.title,
    text: payload.text,
    showCancelButton: true,
    confirmButtonText: payload.confirmButtonText || 'Confirm',
    cancelButtonText: payload.cancelButtonText || 'Cancel',
  }).then((result) => {
    if (result.isConfirmed) {
      if (payload.title === 'Unauthorized') {
        setTimeout(() => {
          logout()
        }, 1000)
      }
      return true
    } else {
      return false
    }
  })
}