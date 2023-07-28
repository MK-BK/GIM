<template>
    <el-container>
        <el-aside width="200px">
            <div class="ui-search">
                <el-input placeholder="请输入内容" v-model="search" clearable></el-input>
            </div>
            <div v-if="!list.length" class="empty-content">请求列表为空</div>
            <div v-else v-for="request of list" class="pannel-item" :class="active(request)" @click="viewRequest(request)" :key="request.ID">
                <avatar :id="request.OwnerID" :name="request.OwnerName"></avatar>
                <div class="session-item">
                    <div class="name">{{request.DisplayName}}</div>
                    <div class="content">{{request.Description}}</div>
                </div>
                <div class="time">{{this.$utils.formatTime(request.UpdatedAt)}}</div>
                <div class="ui-icon-10 ui-icon-delete" @click="deleteRequest(request)"></div>
            </div>
        </el-aside>
        <el-main>
            <router-view :key="key"></router-view>
        </el-main>
    </el-container>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'
import { ref, computed, onMounted } from 'vue'
import useStore from '@/stores'
import { useRouter } from 'vue-router'


const { requestStore } = useStore()
const router = useRouter()

const search = ref('')
const requests = ref([])

const key = computed(()=> {
    return router.currentRoute.name !== undefined? router.currentRoute.name + new Date(): router.currentRoute + new Date()
})

onMounted(async () => {
    await refresh()
})

const list = computed(() =>{
    if (search.value != '') {
        return (requests.value || []).filter((request) => {
            return request.OwnerName.indexOf(search.value) > -1
        })
    }
    return requests.value
})


async function refresh() {
    requests.value = await requestStore.listRequests()
}

function viewRequest(request) {
    requestStore.setRquest(request)
    router.push('/requests/' + request.ID)
}

function active(request) {
    return {
        'active': requestStore.request == request.ID
    }
}

async function deleteRequest(request) {
    await requestStore.deleteRequest(request.ID)
    refresh()
}
</script>