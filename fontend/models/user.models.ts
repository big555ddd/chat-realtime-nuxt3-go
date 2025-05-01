export interface UserResponse {
    id : string,
    username : string,
    email : string,
    password : string,
    first_name : string,
    last_name : string,
    display_name : string,
    role_id : number,
    status : string,
    created_at : number,
    updated_at : number,
}

export interface UpdateProfile {
    first_name : string,
    last_name : string,
    display_name : string,
}


