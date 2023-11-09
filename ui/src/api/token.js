import client from './client'

// 对接后端 Login Api
export var LOGIN = (data) => {
    return client({
        url: '/tokens',
        method: 'post',
        data: data
    })
}