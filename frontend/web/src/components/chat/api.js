import { getRequest } from '../../utils';


export const getGameChannels = (userId, channelId) => getRequest(`/channel/${channelId}/${userId}`)
