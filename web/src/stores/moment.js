import http from '@/js/axios'
import { defineStore } from 'pinia'

const useMomentStore = defineStore('moment', {
	actions: {
		async listMoments() {
			return await http.get('/api/moments')
		},
	
		async getMoment(id) {
			return await http.get(`/api/moments/${id}`)
		},
	
		async createMoment(monment) {
			return await http.post('/api/moments', monment)
		},
	
		async deleteMoment(id) {
			return await http.delete(`/api/moments/${id}`)
		},
	
		async updateMoment(id, moment) {
			return await http.put(`/api/moments/${id}`, moment)
		},
	
		async createComment(moment) {
			return await http.post('/api/moments', moment)
		},
	
		async deleteComment(id) {
			return await http.delete(`/api/moments/${id}`)
		},

		async createComment(comment) {
			return await http.post(`/api/comment`, comment)
		},

		async deleteComment(id) {
			return await http.delete(`/api/comment/${id}`)
		}
	}
})

export default useMomentStore