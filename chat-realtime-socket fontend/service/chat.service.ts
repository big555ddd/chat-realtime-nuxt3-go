import type { CreateChat,AddFriend } from '~/models/chat.models';
import { client } from './httpClient';
import type { Param } from "~/models/http.models";

export function create(body: CreateChat){
    return client("/api/v1/chats/create",{
        method : "POST",
        body: body
    })
}

export function GetMessage(id: string,param: Param ){
    return client(`/api/v1/chats/${id}`,{
        method: "GET",
        params: param
     })
}

export function AddFriend (body: AddFriend){
    return client("/api/v1/chats/add-friend",{
        method : "POST",
        body: body
    })
}

export function RemoveFriend (body: AddFriend){
    return client("/api/v1/chats/remove-friend",{
        method : "POST",
        body: body
    })
}