<template>
    <el-form class="view-profile" label-width="80px">
        <el-form-item>
            <avatar :id="user.ID" :name="user.Name" :size="60"></avatar>
        </el-form-item>
        <el-form-item label="备注信息">
            <div>{{user.Name}}</div>
        </el-form-item>
        <el-form-item label="用户名称">
            <div>{{user.Name}}</div>
        </el-form-item>
        <el-form-item label="用户ID">
            <div>{{user.ID}}</div>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="createSession">发送消息</el-button>
        </el-form-item>
    </el-form>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router';
import useStore from '@/stores'

const route = useRoute()
const router = useRouter()

const { sessionStore, userStore } = useStore()

const user = ref({})

onMounted(async () => {
    await refresh();
})

onMounted(() => {
    watch(() => route.params.id, async () => {
        await refresh();
    });
});

async function refresh() {
    user.value = await userStore.getUser(route.params.id)
}

async function createSession() {
    let session = await sessionStore.createSession({
        Kind: "USER",
        DestinationID: route.params.id
    })
    router.push(`/sessions/${session.ID}`)
}
</script>