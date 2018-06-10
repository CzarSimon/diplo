import { createAction } from 'redux-actions';
import { browserHistory } from 'react-router';
import * as directory from '../api/directory';

/* --- Action types --- */
const ADD_TOKEN = 'diplo/user/token/ADD';
const ADD_USER = 'diplo/user/ADD';
const LOGOUT_USER = 'diplo/user/LOGOUT';
const SET_LOGIN_ERROR = 'diplo/user/login/ERROR';
const SET_SIGNUP_ERROR = 'diplo/user/signup/ERROR';

const initalState = {
  token: null,
  info: null,
  loginError: {
    isError: false,
    message: ''
  },
  signupError: {
    isError: false,
    message: ''
  }
}

/* --- Reducer --- */
const user = (state = initalState, action = {}) => {
  switch (action.type) {
    case ADD_TOKEN:
      return {
        ...state,
        token: action.payload.token,
        loginError: { isError: false, message: null },
        signupError: { isError: false, message: null }
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
        loginError: {
          isError: true,
          message: action.payload.error
        }
      }
    case SET_SIGNUP_ERROR:
      return {
        ...state,
        token: null,
        info: null,
        signupError: {
          isError: true,
          message: action.payload.error
        }
      }
    case LOGOUT_USER:
      return {
        ...state,
        token: null,
        info: null
      }
    default:
      return state;
  }
}

export default user;

/* --- Actions --- */
export const addToken = createAction(ADD_TOKEN, token => ({ token }));

export const addUser = createAction(ADD_USER, user => ({ user }));

export const setLoginError = createAction(SET_LOGIN_ERROR, error => ({ error }));

export const setSignupError = createAction(SET_SIGNUP_ERROR, error => ({ error }));

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
      dispatch(setLoginError('Email and password did not match'));
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
      dispatch(setSignupError(''));
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

export const logoutUser = createAction(LOGOUT_USER);
