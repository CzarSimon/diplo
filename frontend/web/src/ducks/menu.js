import { createAction } from 'redux-actions';

/* --- Action types --- */
const SELECT_TAB = 'diplo/menu/tab/SET';
const OPEN_CHAT_MENU = 'diplo/menu/chat/OPEN';
const CLOSE_CHAT_MENU = 'diplo/menu/chat/CLOSE';

const initalState = {
  selectedTab: null,
  chatMenuOpen: null,
}

/* --- Reducer --- */
const menu = (state = initalState, action = {}) => {
  switch (action.type) {
    case SELECT_TAB:
      return {
        ...state,
        selectedTab: action.payload.tabName
      }
    case OPEN_CHAT_MENU:
      return {
        ...state,
        chatMenuOpen: true,
      }
    case CLOSE_CHAT_MENU:
      return {
        ...state,
        chatMenuOpen: false,
      }
    default:
      return state
  }
}

export default menu;

/* --- Actions --- */
export const selectTab = createAction(SELECT_TAB, tabName => ({ tabName }));

export const openChatMenu = createAction(OPEN_CHAT_MENU);

export const closeChatMenu = createAction(CLOSE_CHAT_MENU);
