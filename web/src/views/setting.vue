<template>
    <el-container>
        <el-header>用户设置</el-header>
        <el-main>
            <el-form label-position="left" label-width="100px" style="max-width: 460px">
                <el-form-item label="头像" class="avatar">
                    <avatar :id="user.ID" :name="user.Name" :size="80"></avatar>
                    <el-upload class="avatar-uploader"
                        :action="url"
                        :headers="headers"
                        :show-file-list="false"
                        :with-credentials='true'
                        :on-success="handler"
                        :before-upload="handler">
                        <el-button size="small" type="primary">点击更换</el-button>
                    </el-upload>
                </el-form-item>
                <el-form-item label="名称">
                    <el-input v-model="user.Name" />
                </el-form-item>
                <el-form-item label="密码">
                    <el-input type="password" v-model="user.Password" show-password />
                </el-form-item>
                <el-form-item label="邮件">
                    <el-input v-model="user.Email" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="update">更新</el-button>
                </el-form-item>
            </el-form>
        </el-main>
    </el-container>
</template>

<script setup>
import avatar from '@/views/components/avatar'

import { ref, onMounted } from 'vue'
import useStore from '@/stores'

const { userStore } = useStore()

const user = ref({})
const url = ref('/api/avatar')
const headers = ref({
    'token': sessionStorage.getItem('token')
})

onMounted(async () => {
    await refresh()
})

async function refresh() {
    user.value = await userStore.getUser(userStore.user.ID)
}

async function update() {
    await userStore.updateUser(userStore.user.ID, {
        Name: user.value.Name,
        Password: user.value.Password,
        Email: user.value.Email,
    })
    refresh()
}

async function handler() {
    refresh()
}
</script>

<style lang="less" scoped>
.el-container {
    padding: 20px 20px;

    .el-header {
        text-align: left;
        padding-left: 0;
    }
    .el-main {
        padding: 10px 20px;
        .avatar {
            .avatar-uploader {
                margin-left: 20px;
                align-self: end;
            }
        }
    }
}
</style>