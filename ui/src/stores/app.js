// 保持当前系统的运行状态
import { useStorage } from '@vueuse/core'

// 把LocalStroage封装成一个响应式的Ref对象
export const state = useStorage(
    'vblog',
    {
        isLogin: false,
        token: {},
    },
    localStorage,
    { mergeDefaults: true }
)