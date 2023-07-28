<template>
	<el-form label-position="top" label-width="100px">
		<el-form-item label="群聊名称">
			<el-input v-model="group.DisplayName" @keyup.native.enter="updateGroup"></el-input>
		</el-form-item>
		<el-form-item label="群公告">
			<el-input type="textarea" v-model="group.Announcement" :autosize="{minRows:4, maxRows:8}" @keyup.native.enter="updateGroup"></el-input>
		</el-form-item>
		<el-form-item label="开启群验证">
			<el-switch v-model="group.Verify" @change="updateGroup"></el-switch>
		</el-form-item>
		<el-form-item label="群聊成员">
			<el-container>
				<div class="ui-search">
					<el-input placeholder="请输入内容" v-model="search" clearable></el-input>
					<el-icon @click="visible=!visible"><Plus /></el-icon>
				</div>
				
				<sessionAdd :visible="visible" :groupID="group.ID" v-on:close="visible=!visible"></sessionAdd>

				<el-main v-for="user of group.Users" :key="user.ID" class="pannel-item none-border">
					<avatar :id="user.ID" :name="user.Name"></avatar>
					<div>{{user.Name}}</div>
					<div>{{getAuth(user.ID)}}</div>
					<div v-if="canDelete(user.ID)" class="ui-icon-10 ui-icon-delete" @click="leaveGroup(user.ID)"></div>
				</el-main>
				<el-footer>
					<el-button v-if="isManager()" type="danger" @click="deleteGroup()">解散群聊</el-button>
					<el-button v-else type="danger" @click="leaveGroup(userStore.user.ID)">退出群聊</el-button>
				</el-footer>
			</el-container>
		</el-form-item>
	</el-form>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'
import sessionAdd from '@/views/components/sessionAdd'

import { ref, onMounted } from 'vue'
import useStore from '@/stores';

const { groupStore, userStore } = useStore()

const props = defineProps(['groupID'])

const search = ref('')
const group = ref({})
const visible = ref(false)

onMounted(async () => {
	await refresh()
})

async function refresh() {
	group.value = await groupStore.getGroup(props.groupID)
}

function isManager() {
	if (userStore.user.ID == group.value.ManagerID) {
		return true
	}
	return false
}

function canDelete(key) {
	if (key == group.value.ManagerID) {
		return false
	}
	return isManager()
}

function getAuth(key) {
	if (key == group.value.ManagerID) {
		return '管理员'
	}
	return '成员'
}

async function leaveGroup(userID) {
	await groupStore.leaveGroup(group.value.ID, userID)
	await refresh()
}

async function deleteGroup() {
	await groupStore.deleteGroup(group.value.ID)
}

async function updateGroup() {
	await groupStore.updateGroup(group.value.ID, {
		Name: group.value.DisplayName,
		Announcement: group.value.Announcement,
		verify: group.value.Verify,
	})
	await refresh()
}
</script>

<style lang="less" scoped>
form {
	padding: 10px 10px;
	border-left: 1px solid #c1b8b8;
	font-size: 14px;
	.el-container {
		.pannel-item {
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 10px 10px;
			border-bottom: 1px solid #00000029;
			color: black;
			margin-top: 10px;
		}
		.none-border {
			border-bottom: none;
			padding-left: 0;
		}
	}
}
</style>