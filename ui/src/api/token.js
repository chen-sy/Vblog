import client from './client'

// 对接后端 Login Api
export var LOGIN = (data) => {
    return client({
        url: '/tokens',
        method: 'post',
        data: data
    })
}
// 退出
export var LOGOUT = (data) => client.delete(`/tokens/`, data)