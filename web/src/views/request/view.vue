<template>
    <el-form class="view-profile" label-width="80px">
        <el-form-item label="用户头像">
            <avatar :id="user.ID" :name="user.Name" :size="50"></avatar>
        </el-form-item>
        <el-form-item label="用户名称">
            <div>{{user.Name}}</div>
        </el-form-item>
        <el-form-item label="微信号">
            <div>{{user.ID}}</div>
        </el-form-item>
        <el-form-item label="请求内容">
            <div>{{request.Description}}</div>
        </el-form-item>
        <el-form-item label="请求时间">
            <div>{{this.$utils.formatTime(request.CreatedAt)}}</div>
        </el-form-item>
        <el-form-item v-if="!request.Status">
            <el-button type="primary" @click="ackRequest">同意请求</el-button>
        </el-form-item>
    </el-form>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'

import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import useStore from '@/stores'

const { userStore, requestStore } = useStore()
const route = useRoute()

const user = ref({})
const request = ref({})

onMounted(async() =>{
    await refresh()
})

onMounted(() => {
    watch(() => route.params.id, async () => {
        await refresh();
    });
});

async function refresh() {
    request.value = await requestStore.getRequest(route.params.id)
    user.value = await userStore.getUser(request.value.OwnerID)
}
        
async function ackRequest() {
    await requestStore.ackRequest(request.value.ID)
    await refresh()
}
</script>