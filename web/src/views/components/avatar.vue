<template>
    <el-avatar v-if="!avatar.isAvatar" shape="square" :size="size">{{this.name}}</el-avatar>
    <el-avatar v-else :src="avatar.src" shape="square" :size="size"/>
</template>

<script setup>
import { onMounted , reactive, ref } from "vue"

const props = defineProps(['id', 'name', 'size'])

const avatar = reactive({
    src: `http://localhost:8080/api/avatar/${props.id}`,
    isAvatar: true
})

onMounted(async() => {
    await refresh()
})

async function refresh() {
    let item = sessionStorage.getItem('avatar/'+ props.id)
    if (item != undefined)  {
        if (item == "") {
            avatar.isAvatar = false
        } else {
            avatar.src = sessionStorage.getItem('avatar/'+ props.id)
        }
    } else {
        try {
            let resp = await this.$store.dispatch('getAvatar', {
                ID: props.id,
            })
            sessionStorage.setItem('avatar/'+ props.id, resp)
        } catch(e) {
            avatar.isAvatar = false
            sessionStorage.setItem('avatar/'+ props.id, '')
        }
    }
}
</script>

<style scoped>
.el-avatar {
    font-size: 13px;
    color: black;
}
</style>