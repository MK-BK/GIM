<template>
	<el-container>
		<el-aside width="200px">
			<div class="ui-search">
				<el-input placeholder="请输入内容" v-model="object.search" clearable></el-input>
				<el-icon @click="object.visible=!object.visible"><Plus /></el-icon>
			</div>
			<createFriend :visible="object.visible" v-on:close="object.visible=!object.visible"></createFriend>
			<el-collapse v-model="object.activeNames">
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
				<el-collapse-item title="好友列表" name="users">
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
            <router-view></router-view>
        </el-main>
	</el-container>
</template>

<script setup>
import createFriend from '@/views/user/create'
import avatar from '@/views/components/avatar'

import { computed, onMounted, reactive } from 'vue'

import { useRouter } from 'vue-router'
import useStore from '@/stores'
const { userStore, groupStore } = useStore()
const router = useRouter()

const object = reactive({
	groups: [],
	users: [],
	visible: false,
	activeNames: ['users','groups'],
	search: ''
})

onMounted(async() => {
	await refresh()
})

const groupList = computed(() => {
	return object.groups.filter((group) => {
		return group.DisplayName.indexOf(object.search) > -1;
	})
})

const usersList = computed(() => {
	return object.users.filter((user) => {
		return user.Name.indexOf(object.search) > -1;
	})
})

function viewUser(user) {
	router.push('/users/' + user.ID)
}

function viewGroup(group) {
	groupStore.setGroup(group)
	router.push('/groups/' + group.ID)
}
		
async function deleteUser(user) {
	await userStore.deleteUser(user.ID)
	await refresh()
}

async function deleteGroup(group) {
	await groupStore.deleteGroup(group.ID)
	await refresh()
}

async function refresh() {
	object.users = await userStore.listFriends()
	object.groups = await groupStore.listGroups()
}
</script>

<style scoped>
.el-collapse {
	padding-left: 20px;
}
</style>