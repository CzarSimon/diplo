import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';

import { DEV_MODE } from './config/main';
import * as reducers from './ducks';

import LoginContainer from './components/login/containers';
import './App.css';

// Redux setup
const createStoreWithMiddleware = (!DEV_MODE)
  ? applyMiddleware(thunk)(createStore)
  : applyMiddleware(thunk, logger)(createStore)
const reducer = combineReducers({
  ...reducers,
  routing: routerReducer
})
const store = createStoreWithMiddleware(reducer);
const history = syncHistoryWithStore(browserHistory, store);

class Home extends Component {
  render() {
    return (
      <p>Home</p>
    )
  }
}

class Map extends Component {
  render() {
    return (
      <p>Map</p>
    )
  }
}

class Chat extends Component {
  render() {
    return (
      <p>Chat</p>
    )
  }
}

class Signup extends Component {
  render() {
    return (
      <p>Signup</p>
    )
  }
}

export default class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <div className="App">
        <Router history={history}>
          <Route path="/" component={Home} />
          <Route path="/game/map" component={Map} />
          <Route path="/game/chat" component={Chat} />
          <Route path="/login" component={LoginContainer} />
          <Route path="/signup" component={Signup} />
        </Router>
        </div>
      </Provider>
    );
  }
}
