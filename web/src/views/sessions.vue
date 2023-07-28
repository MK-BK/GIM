<template>
    <el-container>
        <createSession :visible="visible" v-on:close="visible=!visible"></createSession>
        <el-aside width="200px">
            <div class="ui-search">
				<el-input placeholder="请输入内容" v-model="search" clearable></el-input>
				<el-icon @click="visible=!visible"><Plus /></el-icon>
			</div>
            <div v-if="!list.length" class="empty-content">session列表为空</div>
            <div v-else v-for="session of list" class="pannel-item" :class="active(session)" @click="viewSession(session)"  :key="session.ID">
                <avatar :id="session.DestinationID" :name="session.DisplayName"></avatar>
                <div class="session-item">
                    <div class="name">{{session.DisplayName}}</div>
                    <div class="content">{{session.LatestMessage}}</div>
                </div>
                <div class="time">{{this.$utils.formatTime(session.UpdatedAt)}}</div>
                <div class="ui-icon-10 ui-icon-delete" @click="deleteSession(session)"></div>
            </div>
        </el-aside>
        <el-main>
            <router-view :key="key"></router-view>
        </el-main>
    </el-container>
</template>

<script setup>
import createSession from '@/views/session/create'
import avatar from '@/views/components/avatar.vue'

import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import useStore from '@/stores'

const router = useRouter()
const { sessionStore } = useStore()

const search = ref('')  
const visible = ref(false)
const sessions = ref([])

onMounted(async () => {
    await refresh()
})

const key = computed(()=> {
    return router.currentRoute.name !== undefined? router.currentRoute.name + new Date(): router.currentRoute + new Date()
})

const list = computed(() =>{
    if (search != '') {
        return (sessions.value || []).filter((session) => {
            return session.DisplayName.indexOf(search.value) > -1
        })
    }
    return sessions.value
})

async function refresh() {
    sessions.value = await sessionStore.listSessions()
}

async function viewSession(session) {
    sessionStore.setSession(session)
    router.push('/sessions/' + session.ID)
}

function active(session) {
    return {
        'active': sessionStore.session.ID == session.ID
    }
}

async function deleteSession(session) {
    await sessionStore.deleteSession(session.ID)
    await refresh()
}
</script>
