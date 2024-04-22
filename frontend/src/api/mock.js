import "../mockjs/index";
import request from "../utils/request";

//Mock模拟
export function mockIndex() {
    return request({
        url: '/mock/index',
        method: 'get',
        baseURL:'' //mock请求，需要去掉接口域名
    })
}

//Get请求
export function getApi() {
    return request({
        url: '/get',
        method: 'get',
        baseURL:'http://localhost:10086' //配置baseURL使用当前设置的，不配置使用 .env 中的 VUE_APP_API_URL
    })
}

//Post请求
export function postApi(data) {
    return request({
        url: '/post',
        method: 'post',
        data
    })
}
