import { getRequest, postRequest } from '../../utils';


export const signup = (token, user) => postRequest(`/api/directory/user`, token, user)


export const login = (token, user) => postRequest(`/api/directory/user/login`, token, user)


export const renewToken = (token, user) => getRequest(`/api/directory/user/token/renew`, token)
