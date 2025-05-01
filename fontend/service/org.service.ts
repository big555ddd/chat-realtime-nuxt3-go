import { client } from "./httpClient";

export function list(){
    return client(`/ogn/list`,{
        method : "GET",
    })
}