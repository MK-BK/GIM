import http from '@/js/axios'
import { defineStore } from 'pinia'

const useSessionStore = defineStore('session', {
	state: () => ({ 
		session: {}
	}),
	actions: {
		setSession(session) {
			this.session = session
		},

		async listSessions() {
			return await http.get('/api/sessions')
		},

		async createSession(session) {
			return await http.post('/api/sessions', session)
		},

		async getSession(id) {
			return await http.get(`/api/sessions/${id}`)
		},

		async deleteSession(id) {
			return await http.delete(`/api/sessions/${id}`)
		},

		async sendMessage(id, message) {
			return await http.post(`/api/sessions/${id}/messages`, message)
		},

		async getMessages(id) {
			return await http.get(`/api/sessions/${id}/messages`)
		}
	}
})

export default useSessionStore