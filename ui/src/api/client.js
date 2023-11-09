// 封装一个全局的http客户端
import axios from 'axios'

// 创建axios实例
const instance = axios.create({
    baseURL: '/api/v1',
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json'
    }
});

// 添加请求拦截器
instance.interceptors.request.use(
    config => {
        // 在发送请求之前做一些处理，例如添加请求头、身份验证等
        // config.headers['Authorization'] = 'Bearer token';
        return config;
    },
    error => {
        // 处理请求错误
        console.error(error);
        return Promise.reject(error);
    }
);


// 添加响应拦截器
instance.interceptors.response.use(
    // 请求成功
    response => {
        // 在接收响应之前做一些处理，例如处理错误码、统一处理响应数据等
        return response.data;
    },
    // 请求失败
    error => {
        debugger
        // 处理响应错误
        console.log(error)
        // 默认错误信息使用error.message
        let msg = error.message
        // 如果接口有返回，则使用接口返回的异常
        if (error.response.data && error.response.data.message) {
            msg = error.response.data
            // 自定义业务逻辑处理:
            switch (error.response.data.code) {
                // Token过期, 跳转到Login页面
                case 5001:
                    window.location.assign('/login')
                    break

                default:
                    break
            }
        }
        return Promise.reject(msg);
    }
)

export default instance