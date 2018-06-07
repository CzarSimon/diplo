import { createAction } from 'redux-actions';

/* --- Action types --- */
const SELECT_TAB = 'diplo/menu/tab/SET';

const initalState = {
  selectedTab: '/',
}

/* --- Reducer --- */
const menu = (state = initalState, action = {}) => {
  switch (action.type) {
    case SELECT_TAB:
      return {
        ...state,
        selectedTab: action.payload.tabName
      }
    default:
      return state
  }
}

export default menu;

/* --- Actions --- */
export const selectTab = createAction(SELECT_TAB, tabName => ({ tabName }));
