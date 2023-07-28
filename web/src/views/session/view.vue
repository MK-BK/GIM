<template>
    <div class="session-profile">
        <el-container>
            <el-header>
                <span>{{session.DisplayName}}</span>
                <i v-if="session.Kind=='GROUP'&&session.Enabled" class="ui-icon ui-icon-20 ui-icon-profile" @click="visible=!visible"></i>
            </el-header>
            <el-main>
                <div v-if="!messages.length" class="empty-message">no message to show</div>
                <div v-else v-for="message of messages" :key="message.ID" >
                    <div class="message-item" :class="getPosition(message)">
                        <avatar :id="message.OwnerID" :name="message.OwnerName" :size="40"></avatar>
                        <div class="content">
                            <div class="title">{{this.$utils.formatTime(message.CreatedAt)}}</div>
                            <div class="message">{{message.Content}}</div>
                        </div>
                    </div>
                    <div v-if:="message.Error" class="error">
                        <div class="ui-icon ui-icon-20 ui-icon-error" :title="getMessageError(message)">
                            {{ getMessageError(message) }}
                        </div>
                    </div>
                </div>
                <div id="bottom"></div>
            </el-main>
            <div class="chat-input-container">
                <div class="chat-input">
                    <textarea
                        ref="textarea"
                        class="chat-textarea"
                        placeholder="请输入消息"
                        v-model="content"
                        @input="onInput"
                        @keydown.enter.prevent="sendMessage"></textarea>
                    <button class="chat-send-btn" @click="sendMessage">发送</button>
                </div>
            </div>
        </el-container>
        <sessionProfile v-if="visible" :groupID="session.GroupID"></sessionProfile>
    </div>
</template>

<script setup>
import sessionProfile from '@/views/session/profile.vue'
import avatar from '@/views/components/avatar.vue'

import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import useStore from '@/stores'

const { sessionStore }  = useStore()
const route = useRoute()

const session = ref({})
const messages =  ref([])
const group = ref({})

const content = ref('')
const visible = ref(false)

const bottom = ref(null)

onMounted(async () => {
    await refresh()
    watch(() => route.params.id, async () => {
        await refresh()
        visible.value = false
    });
});

async function refresh() {
    try {
        session.value = await sessionStore.getSession(route.params.id)
    } catch (e) {
        console.log(e)
    }

    messages.value = await sessionStore.getMessages(route.params.id)
    scorllBottom()
}

function getPosition(message) {
    return {
        'position-right': message.OwnerID === session.value.OwnerID
    }
}

function getMessageError(message) {
    if (message.Kind == "GROUP") {
        return "你已经被移除群聊"
    }

    return "你的消息已被拒绝接收"
}

async function sendMessage() {
    if (content != "") {
       await sessionStore.sendMessage(session.value.ID, {
            Content: content.value,
        })

        content.value = '';
    }
    await refresh()
    scorllBottom()
}

function scorllBottom() {
    document.getElementById("bottom").scrollIntoView();
}

function onInput() {
    // nextTick(() => {
    //     this.textareaHeight = `${this.$refs.input.scrollHeight}px`;
    // });
}
</script>

<style lang="less" scoped>
.session-profile {
    display: flex;
    height: 100%;
    overflow: hidden;
    .el-container {
        .el-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
        }
        .el-main {
            border-top: 1px solid #c1b8b8;
            border-bottom: 1px solid #c1b8b8;
            padding-bottom: 10px !important;
            overflow-y: auto;
            height: 80%;

            .message-item {
                display: flex;
                flex-wrap: wrap;
                align-items: center;
                margin-top: 10px;
                padding-left: 5px;
                padding-right: 5px;
                .content {
                    margin-left: 10px;
                    font-size: 12px;
                    padding: 5px 10px;
                    border-radius: 5px;
                    background-color: aliceblue;
                    max-width: 50%;
                    .title {
                        text-align: left;
                    }
                    .message {
                        margin-top: 10px;
                        overflow-wrap: break-word;
                        word-wrap: break-word;
                        text-align: left;
                    }
                }
                .active {
                    background-color:#00800033;
                }
            }

            .error {
                
            }

            .position-right {
                flex-direction: row-reverse;
                .content {
                    margin-right: 10px;
                }
            }
        }

        .el-footer {
            container {
                position: relative;
                width: 100%;
                height: 100%;
            }

            .textarea {
                width: 100%;
                padding: 10px;
                border: none;
                overflow-y: auto;
            }

            .button {
                position: absolute;
                bottom: 10px;
                right: 10px;
                padding: 10px;
            }
        }
    }
    form {
        width: 20%;
        height: 100%;
    }
    .chat-input-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 80px;
        box-sizing: border-box;
        background-color: #f5f5f5;
    }

    .chat-input {
        display: flex;
        align-items: center;
        width: 100%;
        margin: 0 auto;
        background-color: #fff;
        border-radius: 8px;
        box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
        overflow: hidden;
    }

    .chat-textarea {
        flex-grow: 1;
        width: 100%;
        padding: 12px;
        margin-right: 8px;
        border: none;
        outline: none;
        resize: none;
        font-size: 14px;
        line-height: 20px;
    }

    .chat-send-btn {
        display: inline-block;
        padding: 8px 16px;
        border: none;
        outline: none;
        border-radius: 4px;
        background-color: #007aff;
        color: #fff;
        font-size: 14px;
        line-height: 20px;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .chat-send-btn:hover {
        background-color: #0062cc;
    }
}
</style>