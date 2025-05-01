export interface Response<T> {
    status: ResponseMessage
    data?: T
 }
 
 export interface ResponseMessage {
    code: string,
    message: string
 }
 
 export interface ResponseWithPagination<T> extends Response<T> {
    paginate: Paginate
 }
 
 export interface Paginate {
    current_page: number,
    per_page: number,
    total_pages: number,
    total : number
 }

 export interface Param{
   limit? : number,
   page? : number,
   search? : string
 }
