import { fetchPost } from '@/utils/fetch'

/** **************************** 当前任务列表| 已完成任务列表（list）页面 ***************************** **/
// 列表 注意 回来去掉fetchList
export function fetchActile(params) {
  return fetchPost('https://www.easy-mock.com/mock/5cf1f3d6abb0047e81554dc5/index/api/init', params)
}