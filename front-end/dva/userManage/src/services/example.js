import request from '../utils/request';

export async function iquery() {
  return request('/api/iusers');
}
