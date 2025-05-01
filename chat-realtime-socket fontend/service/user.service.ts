import type { UpdateProfile, UserResponse } from "~/models/user.models";
import { client } from './httpClient';
import type { Param } from "~/models/http.models";

export function getUser() {
   return client("/api/v1/auth/user/detail", {
      method: "GET",
   });
}

export function getByID(id: string) {
   return client(`/api/v1/users/${id}`, {
      method: "GET",
   });
}

export function updateByID(id: string ,body:UpdateProfile){
   return client(`/api/v1/users/${id}`,{
      method: "PATCH",
      body: body
   })
}

export function listUser(param: Param){
   return client(`/api/v1/users/list`,{
      method: "GET",
      params: param
   })
}

export function listfriend(id:string,param:Param){
   return client(`/api/v1/users/listfriend/${id}`,{
      method: "GET",
      params: param
   })
}