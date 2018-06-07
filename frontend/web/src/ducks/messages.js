import { createAction } from 'redux-actions';
import { concat } from 'lodash';

/* --- Action types --- */
const APPEND_MESSAGE = 'diplo/messages/APPEND';
const ADD_MESSAGES = 'diplo/messages/ADD';

const initalState = {}

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

export const sendMessage = (text, channel, token) => {
  console.log(`text: ${text}`);
}
