<template>
    <el-dialog v-model="props.visible" title="发表动态">
        <el-container>
            <el-main>
                <el-input type="textarea" :rows="4" placeholder="请输入内容" v-model="content"></el-input>
            </el-main>
            <el-footer>
                <el-button type="success" @click="create">发表动态</el-button>
            </el-footer>
        </el-container>
    </el-dialog>
</template>

<script setup>
import { ref }from 'vue'
import useStore from '@/stores';

const { momentStore } = useStore()

const emits = defineEmits(['close'])
const props = defineProps(['visible'])

const content = ref('')

async function create() {
    await momentStore.createMoment({
        Content: content.value
    })
    emits('close')
}
</script>