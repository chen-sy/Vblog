// 封装http请求
import axios from 'axios'

// 创建axios实例
var instance = axios.create({
    baseURL: '/api/v1',
    timeout: 1000,
    headers: { 'Content-Type': 'application/json' }
});

// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});


// 添加请求拦截器, 统一异常处理
instance.interceptors.response.use(
    // 请求成功
    (resp) => {
        return resp.data
    },
    // 请求失败
    (error) => {
        debugger
        console.log(error)
        // 默认错误信息使用error.message
        var message = error.message
        // 如果接口有返回，则使用接口返回的异常
        if (error.response && error.response.data) {
            message = error.response.data
        }
        return Promise.reject(new Error(message));
    }
)

export default instance