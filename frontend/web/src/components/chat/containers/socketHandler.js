import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import {
  appendMessage,
  registerSocketConnection,
  unregisterSocketConnection
} from '../../../ducks/messages';
import { isValidToken } from '../../../utils/auth';
import { WEBSOCKET_URL } from '../../../config';

class ChatSocketHandler extends Component {
  shouldConnectSocket = () => {
    const { token, socketConnected } = this.props.state;
    return (isValidToken(token) && !socketConnected);
  }

  connectSocket = () => {
    const { token } = this.props.state;
    console.log(`${WEBSOCKET_URL}?token=${token}`);
    this.socket = new WebSocket(`${WEBSOCKET_URL}?token=${token}`);
    this.props.actions.registerSocketConnection();
    this.socket.onmessage = this.handleIncommingMessage;
    this.socket.onclose = this.handleSocketClose;
  }

  handleIncommingMessage = event => {
    console.log(`Recieved message: ${event.data}`);
  }

  handleSocketClose = event => {
    const { token } = this.props.state;
    if (!isValidToken(token)) {
      this.props.actions.unregisterSocketConnection();
    }
  }

  render() {
    console.log(`Socket connected: ${this.props.state.socketConnected}`);
    if (this.shouldConnectSocket()) {
      this.connectSocket();
    }
    return (
      <div />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    token: state.user.token,
    socketConnected: state.messages.socketConnected,
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    appendMessage,
    registerSocketConnection,
    unregisterSocketConnection
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(ChatSocketHandler);
