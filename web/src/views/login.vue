<template>
    <div id="login">
        <el-form label-position="top">
            <template v-if="activeForm=='Login'">
                <el-form-item label="请输入用户名">
                    <el-input v-model="loginData.userName"></el-input>
                </el-form-item>
                <el-form-item label="请输入密码">
                    <el-input v-model="loginData.password" show-password></el-input>
                </el-form-item>
            </template>

            <template v-if="activeForm=='Register'">
                <el-form-item label="请输入用户名">
                    <el-input v-model="registerData.userName"></el-input>
                </el-form-item>
                <el-form-item label="请输入密码">
                    <el-input v-model="registerData.password" show-password></el-input>
                </el-form-item>
                <el-form-item label="请输入邮箱">
                    <el-input v-model="registerData.email"></el-input>
                </el-form-item>
            </template>
           
            <el-form-item>
                <el-button v-if="activeForm=='Login'"  @click="login">登录</el-button>
                <el-button type="primary" @click="register">注册</el-button>
                <el-button v-if="activeForm=='Register'" @click="activeForm='Login'">取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import useStore from '@/stores'

const { userStore } = useStore()
const router = useRouter()

const loginData = reactive({
    userName: '',
    password: '',
})

const registerData = reactive({
    userName: '',
    password: '',
    email: '',
})

const activeForm = ref('Login')

async function login() {
    try {
        await userStore.login({
            name: loginData.userName,
            password: loginData.password,
        })
        router.push('/index')
    } catch(err) {
        console.log(err)
    }
}

async function register() {
    if (activeForm.value == 'Login' ) {
        activeForm.value = 'Register'
        return;
    }

    try {
        await userStore.register({
            name: registerData.userName,
            password: registerData.password,
            email: registerData.email
        })
        activeForm.value = 'Login';
    } catch (e) {
        console.log(err)
    }
}
</script>