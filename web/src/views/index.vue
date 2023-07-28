<template>
    <el-container class="main-layout">
        <el-aside width="50px">
            <avatar :name="user.Name" :id="user.ID"></avatar>
            <div v-for="menu of menus" class="menu-item" :key="menu.name">
                <router-link :to="menu.path">
                    <div class="ui-icon ui-icon-30"
                        :class="[active(menu.name)?`ui-icon-${menu.name}-active`:`ui-icon-${menu.name}`]">
                    </div>
                </router-link>
            </div>
        </el-aside>
        <el-main class="main-body">
            <router-view></router-view>
        </el-main>
    </el-container>
</template>

<script setup>
import avatar from '@/views/components/avatar'
import { computed, reactive } from 'vue'
import { useRoute } from 'vue-router'

import useStore from '@/stores' 

const { userStore } = useStore()

const route = useRoute()
const user = computed(() => {
    return userStore.user
})

const menus = reactive(
    [
        {name: 'sessions', path: '/sessions'},
        {name: 'users', path: '/users'},
        {name: 'requests', path: '/requests'},
        {name: 'moments', path: '/moments'},
        {name: 'setting', path: '/setting'},
    ]
)

function active(path) {
    return route.path.indexOf(path) > -1
}
</script>