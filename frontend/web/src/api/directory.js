import { postRequestJSON, getRequest } from '../utils';

export const loginUser = (email, password) => (
  postRequestJSON('/api/directory/user/login', '', { email, password }))

export const getUser = token => getRequest('api/directory/me', token)
