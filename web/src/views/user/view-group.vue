<template>
    <el-form class="view-profile"  label-width="80px">
        <el-form-item>
            <avatar :id="group.ID" :name="group.DisplayName" :size="60"></avatar>
        </el-form-item>
        <el-form-item label="群聊名称">
            <div>{{group.DisplayName}}</div>
        </el-form-item>
        <el-form-item label="群聊ID">
            <div>{{group.ID}}</div>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="createSession">发送消息</el-button>
        </el-form-item>
    </el-form>
</template>

<script setup>
import avatar from '@/views/components/avatar.vue'
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import useStore from '@/stores'

const route = useRoute()
const router = useRouter()

const { groupStore, sessionStore } = useStore()

const group = ref({})

onMounted(async () => {
    await refresh();
})

onMounted(() => {
    watch(() => route.params.id, async () => {
        await refresh();
    });
});

async function refresh() {
    group.value = await groupStore.getGroup(route.params.id)
}

async function createSession() {
    const session = await sessionStore.createSession({
        Kind: "GROUP",
        GroupID: route.params.id
    })

    router.push(`/sessions/${session.ID}`)
}
</script>