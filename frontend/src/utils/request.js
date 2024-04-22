import axios from 'axios'
import {getToken} from "./token";
import {Notification} from "element-ui";
import router from "@/router";

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_API_URL, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent
    config.headers['token'] = getToken()
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data

    // if the custom code is not 0, it is judged as an error.
    if (res.code !== 200) {

      Notification({
        type:'error',
        title:'温馨提示',
        message:res['message']
      })

      // 501: Illegal token; 502: Other clients logged in; 503: Token expired;
      if (res.code === 501 || res.code === 502 || res.code === 503) {
          //需要登录
        router.push({
          name: 'login'
        })
      }
      return Promise.reject(new Error(res.message || 'Error'))
    }
    return res
  },
  error => {
    console.log('err' + error) // for debug
    return Promise.reject(error)
  }
)

export default service
