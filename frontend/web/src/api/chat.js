import { getRequest } from '../utils/apiUtils';

export const getGameChannels = (token, gameId) => getRequest(`/api/chat/channel/${gameId}`, token)
