<template>
	<el-form label-position="top" label-width="100px">
		<el-form-item label="群聊名称">
			<el-input v-model="groupView.group.DisplayName" @keyup.native.enter="updateGroup"></el-input>
		</el-form-item>
		<el-form-item label="群公告">
			<el-input type="textarea" v-model="groupView.group.Announcement" :autosize="{minRows:4, maxRows:8}" @keyup.native.enter="updateGroup"></el-input>
		</el-form-item>
		<el-form-item label="开启群验证">
			<el-switch v-model="groupView.group.Verify" @change="updateGroup"></el-switch>
		</el-form-item>
		<el-form-item label="群聊成员">
			<el-container>
				<div class="ui-search">
					<el-input placeholder="请输入内容" v-model="groupView.search" clearable></el-input>
					<el-icon @click="groupView.visible=!groupView.visible"><Plus /></el-icon>
				</div>
				<sessionAdd :visible="groupView.visible" :groupID="groupView.group.ID" v-on:close="groupView.visible=!groupView.visible"></sessionAdd>
				<el-main v-for="user of groupView.group.Users" :key="user.ID" class="pannel-item none-border">
					<avatar :id="user.ID" :name="user.Name"></avatar>
					<div>{{user.Name}}</div>
					<div>{{user.Auth}}</div>
					<div>
						<div v-if="user.CanDelete" class="ui-icon-10 ui-icon-delete" @click="leaveGroup(user.ID)"></div>
					</div>
				</el-main>
				<el-footer></el-footer>
			</el-container>
		</el-form-item>
		<el-form-item>
			<el-button v-if="isManager()" type="danger" @click="deleteGroup()">解散群聊</el-button>
			<el-button v-else type="danger" @click="leaveGroup(userStore.user.ID)">退出群聊</el-button>
		</el-form-item>
	</el-form>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'
import sessionAdd from '@/views/components/sessionAdd'

import { onMounted, reactive } from 'vue'
import useStore from '@/stores';

const { groupStore, userStore } = useStore()
const props = defineProps(['groupID'])

const groupView = reactive({
	group: {},
	visible: false,
	search: ''
})

onMounted(async () => {
	await refresh()
})

async function refresh() {
	groupView.group = await groupStore.getGroup(props.groupID)
	for (let user of groupView.group.Users) {
		user.Auth = user.ID ==  groupView.group.ManagerID ? '管理员':'成员'
		user.CanDelete = canDelete(user.ID)
	}
}

function isManager() {
	return userStore.user.ID == groupView.group.ManagerID
}

function canDelete(userID) {
	if (userID == groupView.group.ManagerID) {
		return false
	}
	return isManager()
}

async function leaveGroup(userID) {
	await groupStore.leaveGroup(groupView.group.ID, userID)
	await refresh()
}

async function deleteGroup() {
	await groupStore.deleteGroup(groupView.group.ID)
}

async function updateGroup() {
	await groupStore.updateGroup(groupView.group.ID, {
		Name: groupView.group.DisplayName,
		Announcement: groupView.group.Announcement,
		verify: groupView.group.Verify,
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