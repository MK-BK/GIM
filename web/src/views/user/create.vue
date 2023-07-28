<template>
    <el-dialog v-model="props.visible" title="添加用户">
        <el-container>
            <el-header>
                <el-input placeholder="微信号/名称" v-model="search" @keyup.enter.native="searchUser" clearable></el-input>
            </el-header>
            <el-main>
                <div v-if="!list.length">暂无查询结果</div>
                <div v-else v-for="user of list" class="user-item" :key="user.ID">
                    <div>{{user.Name}}</div>
                    <div v-if="hasBeenFriend(user)">已经是好友了</div>
                    <el-button v-else type="primary" @click="createFriend(user)">添加好友</el-button>
                </div>
            </el-main>
        </el-container>
    </el-dialog>
</template>

<script setup>
import { ref, computed } from "vue";
import useStore from '@/stores'

const { userStore } = useStore()

const props = defineProps(['visible'])
const emits = defineEmits(['close'])

const search = ref('')
const users = ref([])

const list = computed(() => {
    return users.value.filter((user) => {
        return user.ID !== userStore.user.ID
    })
})

async function searchUser() {
    try {
        users.value = await userStore.searchUsers({
            Search: search.value
        });
    } catch (e) {
        console.log(e)
        users.value = []
    }
}

async function createFriend(user) {
    await userStore.addFriend({
        ID: user.ID,
        Description: 'test'
    })
    emits('close')
}

function hasBeenFriend(user) {
    return (userStore.user.Relations || []).some(relation => {
        return relation == user.ID
    });
}
</script>