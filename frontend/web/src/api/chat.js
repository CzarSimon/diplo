import { getRequest, postRequest } from '../utils/apiUtils';

export const getGameChannels = (token, gameId) => (
  getRequest(`/api/chat/channel/${gameId}`, token))

export const loadChannelMessages = (token, channelId) => (
  getRequest(`/api/chat/message/${channelId}`, token))

export const sendMessage = (token, channelId, message) => (
  postRequest(`/api/chat/message/${channelId}`, token, message))
