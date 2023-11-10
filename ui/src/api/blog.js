import client from './client';

// 获取博客列表
export var LIST_BLOG = (params) => client.get('/blogs/list', { params })

// 获取博客详情
export var GET_BLOG = (id) => client.get(`/blogs/details/${id}`)

// 创建博客
export var CREATE_BLOG = (data) => client.post('/blogs/', data)

// 更新博客(全量)
export var UPDATE_BLOG = (id, data) => client.patch(`/blogs/${id}`, data)

// 更新博客(增量)
export var PUT_BLOG = (id, data) => client.put(`/blogs/${id}`, data)

// 删除博客
export var DELETE_BLOG = (id) => client.delete(`/blogs/${id}`)

// 搜索博客
export var SEARCH_BLOG = (params) => client.get('/blogs/search', { params })
