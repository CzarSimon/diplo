import { createAction } from 'redux-actions';
import { browserHistory } from 'react-router';
import * as directory from '../api/directory';

/* --- Action types --- */
const ADD_TOKEN = 'diplo/user/token/ADD';
const ADD_USER = 'diplo/user/ADD';
const SET_LOGIN_ERROR = 'diplo/user/login/ERROR';
const SET_SIGNUP_ERROR = 'diplo/user/signup/ERROR';

const initalState = {
  token: null,
  info: null,
  loginError: false,
  signupError: false
}

/* --- Reducer --- */
const user = (state = initalState, action = {}) => {
  switch (action.type) {
    case ADD_TOKEN:
      return {
        ...state,
        token: action.payload.token,
        loginError: false,
        signupError: false
      }
    case ADD_USER:
      return {
        ...state,
        info: action.payload.user
      }
    case SET_LOGIN_ERROR:
      return {
        ...state,
        token: null,
        info: null,
        loginError: true
      }
    case SET_SIGNUP_ERROR:
      return {
        ...state,
        token: null,
        info: null,
        signupError: true
      }
    default:
      return state;
  }
}

export default user;

/* --- Actions --- */
export const addToken = createAction(ADD_TOKEN, token => ({ token }));

export const addUser = createAction(ADD_USER, user => ({ user }));

export const setLoginError = createAction(SET_LOGIN_ERROR);

export const setSignupError = createAction(SET_SIGNUP_ERROR);

export const loginUser = (email, password) => (
  dispatch => (
    directory.login(email, password)
    .then(token => {
      dispatch(addToken(token));
      directory.getUser(token)
        .then(user => {
          dispatch(addUser(user));
          browserHistory.replace('/');
        })
    })
    .catch(err => {
      dispatch(setLoginError());
    })
  )
)

export const signupUser = user => (
  dispatch => (
    directory.signup(user)
    .then(res => {
      dispatch(addToken(res.token));
      dispatch(addUser(res.user));
      browserHistory.replace('/');
    })
    .catch(err => {
      dispatch(setSignupError());
    })
  )
)

export const getUser = token => (
  dispatch => (
    directory.getUser(token)
    .then(user => {
      dispatch(addUser(user));
    })
  )
)
