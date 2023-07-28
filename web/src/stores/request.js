import http from '@/js/axios'
import { defineStore } from 'pinia'

const useRequestStore = defineStore('request',{
	state: () => ({ 
		request: {}
	}),
	actions: {
		setRquest(request) {
			this.request = request
		},

		async listRequests() {
			return await http.get('/api/requests')
		},

		async getRequest(id) {
			return await http.get(`/api/requests/${id}`)
		},

		async ackRequest(id) {
			return await http.post(`/api/requests/${id}`)
		},

		async deleteRequest(id) {
			return await http.delete(`/api/requests/${id}`)
		}
	}
})

export default useRequestStore