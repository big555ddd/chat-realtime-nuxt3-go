export interface Message {
    id : string,
    sender : string,
    recipient : string,
    type : string,
    detail : string,
    created_at : number,
    updated_at : number,
}

export interface Contact {
    id : string,
    username : string,
    first_name : string,
    last_name : string,
    display_name : string,
    status : string,
    created_at : number,
    avatar: string | null;
    type : string,
    detail : string,
    unread : number,
    message_time : number,
}

export interface CreateChat{
    recipient : string,
    type : string,
    detail : string;
}

export interface AddFriend {
    user_id: string;
}


