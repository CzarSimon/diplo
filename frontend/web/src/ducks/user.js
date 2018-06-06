import { createAction } from 'redux-actions';
import { browserHistory } from 'react-router';
import * as directory from '../api/directory';

/* --- Action types --- */
const ADD_TOKEN = 'diplo/user/token/ADD';
const ADD_USER = 'diplo/user/ADD';
const LOGIN_ERROR = 'diplo/user/login/ERROR';

const initalState = {
  token: null,
  info: null,
  loginError: false,
}

/* --- Reducer --- */
const user = (state = initalState, action = {}) => {
  switch (action.type) {
    case ADD_TOKEN:
      return {
        ...state,
        token: action.payload.token,
        loginError: false
      }
    case ADD_USER:
      return {
        ...state,
        user: action.payload.user
      }
    case LOGIN_ERROR:
      return {
        ...state,
        token: null,
        info: null,
        loginError: true
      }
    default:
      return state;

  }
}

export default user;

/* --- Actions --- */
export const addToken = createAction(ADD_TOKEN, token => ({ token }));

export const addUser = createAction(ADD_USER, user => ({ user }));

export const setLoginError = createAction(LOGIN_ERROR);

export const loginUser = (email, password) => (
  dispatch => (
    directory.loginUser(email, password)
    .then(token => {
      dispatch(addToken(token.token));
      directory.getUser(token.token)
        .then(user => {
          dispatch(addUser(user));
          browserHistory.replace('/');
        })
    })
    .catch(err => {
      console.log(err);
      dispatch(setLoginError());
    })
  )
)

export const signupUser = user => {
  console.log(user);
}
