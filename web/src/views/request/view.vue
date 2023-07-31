<template>
    <el-form class="view-profile" label-width="80px">
        <el-form-item label="用户头像">
            <avatar :id="requestView.user.ID" :name="requestView.user.Name" :size="50"></avatar>
        </el-form-item>
        <el-form-item label="用户名称">
            <div>{{requestView.user.Name}}</div>
        </el-form-item>
        <el-form-item label="微信号">
            <div>{{requestView.user.ID}}</div>
        </el-form-item>
        <el-form-item label="请求内容">
            <div>{{requestView.request.Description}}</div>
        </el-form-item>
        <el-form-item label="请求时间">
            <div>{{this.$utils.formatTime(requestView.request.CreatedAt)}}</div>
        </el-form-item>
        <el-form-item v-if="!requestView.request.Status">
            <el-button type="primary" @click="ackRequest">同意请求</el-button>
        </el-form-item>
    </el-form>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'

import { onMounted, watch, reactive } from 'vue';
import { useRoute } from 'vue-router';

import useStore from '@/stores'
const { userStore, groupStore, requestStore } = useStore()

const route = useRoute()

const requestView = reactive({
    user: {},
    request: {}
})

onMounted(async() => {
    await refresh()
    watch(() => route.params.id, async () => {
        await refresh();
    });
});

async function refresh() {
    requestView.request = await requestStore.getRequest(route.params.id)
    requestView.user = await userStore.getUser(requestView.request.OwnerID)

    if (requestView.request.Kind == 'GROUP') {
        let group = await groupStore.getGroup(requestView.request.GroupID)
        let destination = await userStore.getUser(requestView.request.DestinationID)
        requestView.request.Description = `[${requestView.user.Name}] 邀请 [${destination.Name}] 加入 [${group.DisplayName}] 群聊`
    } else {
        requestView.request.Description = `${requestView.user.Name} 请求添加您为好友`
    }
}
        
async function ackRequest() {
    await requestStore.ackRequest(requestView.request.ID)
    await refresh()
}
</script>