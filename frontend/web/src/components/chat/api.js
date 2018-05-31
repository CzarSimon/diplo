import { getRequest } from '../../utils';


export const getGameChannels = (token, channelId) => getRequest(`/channel/${channelId}`, token)
