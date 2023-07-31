import  http  from '@/js/axios'
import { defineStore } from 'pinia'

const useGroupStore = defineStore('group',  {
	state: () => ({
		group: {},
	}),
	actions: {
		setGroup(group) {
			this.group = group
		},
	
		async listGroups() {
			return await http.get('/api/groups')
		},
		
		async createGroup(group) {
			return await http.post('/api/groups', group)
		},
		
		async updateGroup(id, group) {
			return await http.post(`/api/groups/${id}`, group)
		},
	
		async getGroup(id) {
			return await http.get(`/api/groups/${id}`)
		},
	
		async deleteGroup(id) {
			return await http.delete(`/api/groups/${id}`)
		},
	
		async joinGroup(id, users) {
			return await http.post(`/api/groups/${id}/join`, {
				UserIDs: users
			})
		},
	
		async leaveGroup(id, key) {
			return await http.post(`/api/groups/${id}/leave`,{
				UserID: key
			})
		}
	}
})

export default useGroupStore	