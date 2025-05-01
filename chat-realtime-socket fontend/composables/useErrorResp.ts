import { swalToast } from './useAlertDialog'
import { logout } from '@/service/auth.service'


export const errorResp = (error: any) => {
    if (error) {
      
      var httpCode = error.status
      
      var statusData = error._data;
      // console.log(error._data)
        var message = statusData?.message || "An error occurred";
              
        var payloadCode = statusData?.code;
        // var message = error.data.status.message
        // var payloadCode = error.data.status.code

        if (httpCode == 403 || httpCode == 401) {
            swalToast({
              icon: 'error',
              title: 'กรุณาเข้าสู่ระบบอีกครั้ง',
            })
            setTimeout(() => {
              logout()
            }, 1000)
            return false
          }

          switch (payloadCode) {
            case '0001':
              message = 'token หมดอายุกรุณาเข้าสู่ระบบใหม่'
              break
          }
      

        swalToast({
            icon: 'error',
            title: message,
        })
    }
}
