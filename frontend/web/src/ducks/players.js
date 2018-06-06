import { createAction } from 'redux-actions';

/* --- Action types --- */
const ADD_PLAYER = 'diplo/players/ADD';

const initalState = {}

/* --- Reducer --- */
const players = (state = initalState, action = {}) => {
  switch (action.type) {
    case ADD_PLAYER:
      return {
        ...state,
        [action.payload.player.id]: action.payload.player
      }
    default:
      return state;

  }
}

export default players;

/* --- Actions --- */
export const addPlayer = createAction(ADD_PLAYER, player => ({ player }));
