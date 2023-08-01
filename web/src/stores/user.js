import http  from '@/js/axios'
import { defineStore } from 'pinia'

const useUserStore = defineStore('user', {
	state: () => ({ 
		token: '',
		user: {}
	}),
	actions: {
		setToken(token) {
			this.token = token
		},

		setUser(user) {
			this.user = user
		},

		async login(payload) {
			let resp = await http.post('/api/login', payload)
			if (resp.headers['token']) {
				this.setToken(resp.headers['token'])
				sessionStorage.setItem('token', resp.headers['token'])
			}

			this.setUser(resp.data)
		},

		async register(payload) {
			return http.post('/api/register', payload)
		},

		async listFriends() {
			return await http.get('/api/friends')
		},

		async getUser(id) {
			return await http.get(`/api/users/${id}`);
		},

		async updateUser(id, user) {
			return await http.put(`/api/users/${id}`, user)
		},

		async addFriend(payload) {
			return await http.post('/api/friends', payload)
		},

		async deleteUser(id) {
			return http.delete(`/api/friends/${id}`)
		},

		async searchUsers(payload) {
			return await http.post(`/api/users`, payload)
		},

		async getAvatar(id) {
			return await http.get(`/api/avatar/${id}`)
		}
	},
	persist: {
		key: 'user',
		storage: sessionStorage
	},
})

export default useUserStore