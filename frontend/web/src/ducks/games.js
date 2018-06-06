import { createAction } from 'redux-actions';

/* --- Action types --- */
const SET_ACTIVE_GAME = 'diplo/games/SET';
const ADD_GAME = 'diplo/games/ADD';

const initalState = {
  activeId: null,
  all: {}
}

/* --- Reducer --- */
const games = (state = initalState, action = {}) => {
  switch (action.type) {
    case SET_ACTIVE_GAME:
      return {
        ...state,
        activeId: action.payload.id
      }
    case ADD_GAME:
      return {
        ...state,
        all: {
          ...state.all,
          [action.payload.game.id]: action.payload.game
        }
      }
    default:
      return state

  }
}

export default games;

/* --- Actions --- */
export const setActiveGame = createAction(SET_ACTIVE_GAME, id => ({ id }));

export const addGame = createAction(ADD_GAME, game => ({ game }));
