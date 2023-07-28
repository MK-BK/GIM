<template>
	<el-container>
		<el-aside width="200px">
			<createFriend :visible="visible" v-on:close="visible=!visible"></createFriend>

			<div class="ui-search">
				<el-input placeholder="请输入内容" v-model="search" clearable></el-input>
				<el-icon @click="visible=!visible"><Plus /></el-icon>
			</div>

			<el-collapse v-model="activeNames">
				<el-collapse-item title="群聊列表" name="groups">
					<div v-if="!groupList.length">群聊列表为空</div>
					<div v-else v-for="group of groupList" class="pannel-item none-border" @click="viewGroup(group)" :key="group.ID">
						<div class="avatar-item">
							<avatar :id="group.ID" :name="group.DisplayName"></avatar>
							<div v-if="group.DisplayName">{{group.DisplayName}}</div>
							<div v-else>{{group.ID}}</div>
						</div>
						<div class="ui-icon-10 ui-icon-delete" @click="deleteGroup(group)"></div>
					</div>
				</el-collapse-item>

				<el-collapse-item title="好友列表" name="uses">
					<div v-if="!usersList.length">用户列表为空</div>
					<div v-else v-for="user of usersList" class="pannel-item none-border" @click="viewUser(user)" :key="user.ID">
						<div class="avatar-item">
							<avatar :id="user.ID" :name="user.Name" ></avatar>
							<div>{{user.Name}}</div>
						</div>
						<div class="ui-icon-10 ui-icon-delete" @click="deleteUser(user)"></div>
					</div>
				</el-collapse-item>
			</el-collapse>
		</el-aside>
		<el-main>
            <router-view :key="key"></router-view>
        </el-main>
	</el-container>
</template>

<script setup>
import createFriend from '@/views/user/create'
import avatar from '@/views/components/avatar'

import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

import useStore from '@/stores'

const { userStore, groupStore } = useStore()
const router = useRouter()

const search = ref('')
const activeNames= ref(['users'])

const visible = ref(false)

const groups = ref([])
const users = ref([])

onMounted(async() => {
	await refresh()
})

const key = computed(()=> {
    return router.currentRoute.name !== undefined? router.currentRoute.name + new Date(): router.currentRoute + new Date()
})

const groupList = computed(() => {
	return groups.value.filter((group) => {
		return group.DisplayName.indexOf(search.value) > -1;
	})
})

const usersList = computed(() => {
	return users.value.filter((user) => {
		return user.Name.indexOf(search.value) > -1;
	})
})

function viewUser(user) {
	// userStore.setUser(user)
	router.push('/users/' + user.ID)
}

function viewGroup(group) {
	groupStore.setGroup(group)
	router.push('/groups/' + group.ID)
}
		
function deleteUser(user) {
	userStore.deleteUser(user.ID)
	refresh()
}

function deleteGroup(group) {
	groupStore.deleteGroup(group.ID)
	refresh()
}

async function refresh() {
	users.value = await userStore.listFriends()
	groups.value = await groupStore.listGroups()
}
</script>

<style scoped>
.el-collapse {
	padding-left: 20px;
}
</style>