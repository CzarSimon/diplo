import { createAction } from 'redux-actions';

/* --- Action types --- */
//const APPEND_MESSAGE = 'diplo/messages/APPEND';
const ADD_MESSAGES = 'diplo/messages/ADD';

const initalState = {}

/* --- Reducer --- */
const messages = (state = initalState, action = {}) => {
  switch (action.type) {
    case ADD_MESSAGES:
      const { gameId, channelId, messages } = action.payload;
      return {
        ...state,
        [gameId]: {
          ...state[gameId],
          [channelId]: messages
        }
      }
    default:
      return state;

  }
}

export default messages;

/* --- Actions --- */
export const addMessages = createAction(ADD_MESSAGES,
  (gameId, channelId, messages) => ({ gameId, channelId, messages }));

export const sendMessage = (text, channel, token) => {
  console.log(`text: ${text}`);
}
