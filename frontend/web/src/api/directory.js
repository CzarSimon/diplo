import { postRequestJSON, getRequest } from '../utils';

export const login = (email, password) => (
  postRequestJSON('/api/directory/user/login', null, { email, password })
  .then(res => res.token));

export const getUser = token => getRequest('api/directory/me', token);

export const signup = user => postRequestJSON('/api/directory/user', null, user);
