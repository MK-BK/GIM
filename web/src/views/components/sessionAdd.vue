<template>
    <el-dialog v-model="props.visible"  title="添加好友">
        <el-container>
            <el-header>
                <el-input placeholder="搜索" v-model="search" clearable></el-input>
            </el-header>
            <el-main>
                <div v-if="!userLists.length" class="empty-content">好友列表为空</div>
                <div v-else v-for="user of userLists" class="user-item" :key="user.ID">
                    <el-checkbox v-model="user.checked" :label="user.Name" size="large" />
                </div>
            </el-main>
            <el-footer>
                <el-button @click="emits('close')">取消</el-button>
                <el-button type="primary" @click="create" :disabled="!disabled">邀请好友</el-button>
            </el-footer>
        </el-container>
  </el-dialog>
</template>

<script setup>
import { ref, computed, onMounted } from "vue"
import useStore from '@/stores'

const { groupStore, userStore } = useStore()

const props = defineProps(['visible', 'groupID'])
const emits = defineEmits(['close'])

const search = ref('')
const users = ref([])
const group = ref({})

const disabled = computed(() => {
    return users.value.some(function(user) {
        return user.checked;
    })
})

const userLists = computed(() => {
    if (search.value !== ''){
        return (users.value || []).filter((user) => {
            return user.Name.indexOf(search.value) > -1
        })
    }
    return users.value
})

onMounted(async () => {
    await refresh()
})

async function refresh() {
    group.value = await groupStore.getGroup(props.groupID)

    users.value = (await userStore.listFriends()).map(element => {
        element.checked = false;
        return element;
    });
}

async function create() {
    const selectUsers = []
    users.value.forEach((user) => {
        if (user.checked) {
            selectUsers.push(user.ID)
        }
    })


    if (selectUsers.length > 0) {
        await groupStore.joinGroup(props.groupID, selectUsers)
    }
    emits('close')
}
</script>