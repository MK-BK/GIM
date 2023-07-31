<template>
    <el-dialog v-model="props.visible" title="创建群聊">
        <el-container>
            <el-header>
                <el-input placeholder="搜索" v-model="view.search" clearable></el-input>
            </el-header>
            <el-main>
                <div v-if="!userLists.length" class="empty-content">好友列表为空</div>
                <div v-else v-for="user of userLists" class="user-item" :key="user.ID">
                    <el-checkbox v-model="user.checked" :label="user.Name" size="large" />
                </div>
            </el-main>
            <el-footer>
                <el-button @click="emits('close')">取消</el-button>
                <el-button type="primary" @click="create" :disabled="!disabled">创建群聊</el-button>
            </el-footer>
        </el-container>
  </el-dialog>
</template>

<script setup>
import { computed, onMounted, reactive } from "vue"
import useStore from '@/stores'

const { sessionStore, userStore, groupStore } = useStore()

const props = defineProps(['visible'])
const emits = defineEmits(['close'])

const view = reactive({
    users: [],
    search: ''
})

onMounted(async () => {
    await refresh()
})

const disabled = computed(() => {
    return view.users.some(function(user) {
        return user.checked;
    })
})

const userLists = computed(() => {
    if (view.search !== ''){
        return (view.users || []).filter((user) => {
            return user.Name.indexOf(view.search) > -1
        })
    }
    return view.users
})

async function refresh() {
    view.users = (await userStore.listFriends()).map(element => {
        element.checked = false;
        return element;
    });
}

async function create() {
    const selectUsers = []
    view.users.forEach((user) => {
        if (user.checked) {
            selectUsers.push(user.ID)
        }
    })

    if (selectUsers.length > 1) {
        const group = await groupStore.createGroup({
            Users: selectUsers
        })

        await sessionStore.createSession({
            Kind: "GROUP",
            GroupID: view.group.ID
        })
    }
    emits('close')
}
</script>