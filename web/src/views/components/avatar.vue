<template>
    <el-avatar v-if="!avatar" shape="square" :size="size">{{this.name}}</el-avatar>
    <el-avatar v-else :src="src" shape="square" :size="size"/>
</template>

<script setup>
import { onMounted , ref } from "vue"

const props = defineProps(['id', 'name', 'size'])

const src = ref(`http://localhost:8080/api/avatar/` + props.id)
const avatar = ref(true)

onMounted(async() => {
    await refresh()
})

async function refresh() {
    let item = sessionStorage.getItem('avatar/'+ props.id)

    if (item != undefined)  {
        if (item == "") {
            avatar.value = false
        } else {
            src = sessionStorage.getItem('avatar/'+ props.id)
        }
    } else {
        try {
            let resp = await this.$store.dispatch('getAvatar', {
                ID: props.id,
            })
            sessionStorage.setItem('avatar/'+ props.id, resp)
        } catch(e) {
            avatar.value = false
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