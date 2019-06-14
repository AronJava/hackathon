import axios from 'axios'
import qs from 'qs'

const service = axios.create({
    // timeout: 5000  // 请求超时时间
})

// get方法
export const fetchGet = (url, params = {}) => {
    return service.get(`${url}`, {params})
  }
// post方法
export const fetchPost = (url, params = {}, config) => {
    return service.post(`${url}`, qs.stringify(params), config)
}