import useGroupStore from './group'
import useMomentStore from './moment'
import useRequestStore from './request'
import useSessionStore from './session'
import useUserStore from './user'

export default function useStore() {
    return {
        groupStore: useGroupStore(),
        momentStore: useMomentStore(),
        requestStore: useRequestStore(),
        sessionStore: useSessionStore(),
        userStore: useUserStore()
    }
}