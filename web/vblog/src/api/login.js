import instance from './client'

// 对接后端 Login Api
export var LOGIN = (data) => {
    return instance.post('/tokens', data)
}