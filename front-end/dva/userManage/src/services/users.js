// 处理异步请求
import request from '../utils/request';
import qs from 'qs';
async function fetch({page = 1}) {
  // return request(`/api/users?$(qs.stringify(params))`);
  return request(`/api/users?_page=${page}&_limit=5`);
}
