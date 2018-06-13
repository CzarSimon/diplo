import { createAction } from 'redux-actions';
import { concat } from 'lodash';
import * as chat from '../api/chat';

/* --- Action types --- */
const APPEND_MESSAGE = 'diplo/messages/APPEND';
const ADD_MESSAGES = 'diplo/messages/ADD';
const CONNECT_SOCKET = 'diplo/messages/socket/CONNECT';
const DISCONNECT_SOCKET = 'diplo/messages/socket/DISCONNECT';

const initalState = {
  socketConnected: false
}

/* --- Reducer --- */
const messages = (state = initalState, action = {}) => {
  switch (action.type) {
    case ADD_MESSAGES:
      const { channelId, messages } = action.payload;
      return {
        ...state,
        [channelId]: messages
      }
    case APPEND_MESSAGE:
      const { message } = action.payload;
      return {
        ...state,
        [message.channelId]: concat(state[message.channelId], message)
      }
    case CONNECT_SOCKET:
      return {
        ...state,
        socketConnected: true
      }
    case DISCONNECT_SOCKET:
      return {
        ...state,
        socketConnected: false
      }
    default:
      return state;

  }
}

export default messages;

/* --- Actions --- */
export const addMessages = createAction(
  ADD_MESSAGES, (channelId, messages) => ({ channelId, messages }));

export const appendMessage = createAction(
  APPEND_MESSAGE, messages => ({ messages }));

export const loadMessages = (token, channelId) => (
  dispatch => (
    chat.loadChannelMessages(token, channelId)
      .then(messages => {
        dispatch(addMessages(channelId, messages));
      })
  )
)

export const sendMessage = (text, channelId, token) => (
  dispatch => (
    chat.sendMessage(token, channelId, { text })
      .then(res => {
        console.log('Message sent');
      })
  )
)

export const registerSocketConnection = createAction(CONNECT_SOCKET);

export const unregisterSocketConnection = createAction(DISCONNECT_SOCKET);
