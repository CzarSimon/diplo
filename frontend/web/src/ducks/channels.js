import { createAction } from 'redux-actions';
import _ from 'lodash';
import * as chat from '../api/chat';

/* --- Action types --- */
const SET_CHANNELS = 'diplo/channels/SET';
const ADD_CHANNEL = 'diplo/channels/ADD';
const SET_ACTIVE_CHANNEL = 'diplo/channels/active/SET';

const initalState = {}

/* --- Reducer --- */
const channels = (state = initalState, action = {}) => {
  switch (action.type) {
    case SET_CHANNELS:
      return {
        ...state,
        [action.payload.gameId]: {
          ...state[action.payload.gameId],
          channels: {
            ..._.keyBy(action.payload.channels, 'id')
          }

        }
      }
    case ADD_CHANNEL:
      return {
        ...state,
        [action.payload.gameId]: {
          ...state[action.payload.gameId],
          channels: {
            ...state[action.payload.gameId].channels,
            [action.payload.channel.id]: action.payload.channel
          }
        }
      }
    case SET_ACTIVE_CHANNEL:
      return {
        ...state,
        [action.payload.gameId]: {
          ...state[action.payload.gameId],
          activeChannel: action.payload.channel
        }
      }
    default:
      return state;
  }
}

export default channels;

/* --- Actions --- */
export const setChannels = createAction(
  SET_CHANNELS, (gameId, channels) => ({ gameId, channels }));

export const addChannels = createAction(
  ADD_CHANNEL, (gameId, channel) => ({ gameId, channel }));

export const setActiveChannel = createAction(
  SET_ACTIVE_CHANNEL, (gameId, channel) => ({ gameId, channel }));

export const loadChannels = (token, gameId) => (
  dispatch => (
    chat.getGameChannels(token, gameId)
      .then(channels => {
        dispatch(setChannels(gameId, channels));
      })
      .catch(err => {
        console.log(err);
      })
  )

)
