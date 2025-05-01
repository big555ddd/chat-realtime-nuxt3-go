import { defineStore } from 'pinia'


interface state {
  id: string,
  username: string,
  email: string,
  password: string,
  first_name: string,
  last_name: string,
  display_name: string,
  role_id: number,
  status: string,
  created_at: number,
  updated_at: number,
  loading: boolean,
}

export const useIndexStore = defineStore('index', {
  state: (): state => ({
    id: '',
    username: '',
    email: '',
    password: '',
    first_name: '',
    last_name: '',
    display_name: '',
    role_id: 0,
    status: '',
    created_at: 0,
    updated_at: 0,
    loading: false,
  }),
  actions: {
    setID(ID: string) {
      this.id = ID
    },
  },
})
