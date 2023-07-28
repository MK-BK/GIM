<template>
    <div class="moment-body">
        <div class="moment-background">
            <div class="moment-action ui-icon ui-icon-30 ui-icon-photo" @click="visible=!visible" title="创建动态"></div>
        </div>
        
        <createMoment :visible="visible" v-on:close="visible=!visible"></createMoment>

        <div v-if="!moments.length" class="empty-content">moments列表为空</div>
        <div v-else class="moments" v-for="moment of moments" :key="moment.id">
            <div class="moment">
                <avatar :id="moment.OwnerID" :key="moment.OwnerID" :name="moment.DisplayName"></avatar>
                <div class="content">
                    <div >{{moment.Content}}</div>
                    <div>{{this.$utils.formatTime(moment.CreatedAt)}}</div>
                </div>
                <div class="actions">
                    <el-icon :size="20" @click="moment.checked=!moment.checked"><Comment /></el-icon>
                </div>
            </div>
            <div v-if="moment.checked" class="comment">
                <el-input v-model="moment.Value"></el-input>
                <el-button @click="createComment(moment)">发布</el-button>
            </div>
            <div class="comments">
                <div>评论信息:</div>
                <div class="comment" v-for="comment of moment.Comments" :key="comment.id">
                    <avatar :id="comment.OwnerID" :key="comment.OwnerID" :name="comment.DisplayName"></avatar>
                    <div class="content">
                        <div >{{comment.Content}}</div>
                        <div>{{this.$utils.formatTime(comment.CreatedAt)}}</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import createMoment from '@/views/moment/create'
import avatar from '@/views/components/avatar'

import { ref, onMounted } from 'vue';
import  useStore from '@/stores'

const { momentStore } = useStore()

const moments = ref([])
const visible = ref(false)

onMounted(async () => {
    await refresh()
})

async function refresh() {
    moments.value = await momentStore.listMoments()
}

async function star(moment) {
    await momentStore.updateMoment(moment.ID, {
        Action: 'ADD'
    })
    refresh()
}

async function createComment(moment) {
    await momentStore.createComment({
        MomentID: moment.ID,
        Content:  moment.Value,
    })
    refresh()
}
</script>

<style lang="less" scoped>
.moment-body {
    font-size: 14px;
    height: 100%;
    width: 100%;
    overflow-y: auto;

    .moment-background {
        background-image: url("@/icon/background/moment-background.png");
        height: 200px;
    }

    .moment-action {
        position: relative;
        z-index: 1000;
        left: 96%;
        top: 20px;
    }

    .moments {
        padding: 10px 20px;
        border-bottom: 1px solid #0000004a;
        text-align: left;

        .moment {
            width: 100%;
            font-size: 14px;
            display: flex;
            justify-content: space-between;

            > .content {
                flex: 4;
                padding: 0 10px;
                overflow-wrap: anywhere;
                font-size: 15px;
                align-self: center;
                justify-content: space-between;
                display: flex;
            }

            > .actions {
                display: flex;
                align-self: center;
            }
        }

        .comments {
            margin-top: 10px;
            margin-left: 30px;
        }

        .comment {
            width: 100%;
            font-size: 14px;
            display: flex;
            justify-content: space-between;
            margin-top: 10px;

            > .content {
                flex: 4;
                padding: 0 10px;
                overflow-wrap: anywhere;
                font-size: 15px;
                align-self: center;
                justify-content: space-between;
                display: flex;
            }

            > .actions {
                display: flex;
                align-self: center;
            }
        }
     
    }
}
</style>