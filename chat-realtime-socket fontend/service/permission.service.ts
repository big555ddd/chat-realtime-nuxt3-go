import type { QueryParam } from '~/models/permission.models';
import { client } from './httpClient';

export function list(query: QueryParam){
    return client("permission/list",{
        method : "GET",
        query: query,
    })
}

export function getpermission(id: number){
    return client(`/role/permission/${id}`,{
        method : "GET",
    })
}