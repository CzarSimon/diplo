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
    if (this.socket) {
      return;
    }
    this.socket = new WebSocket(`${WEBSOCKET_URL}?token=${this.props.state.token}`);
    this.props.actions.registerSocketConnection();
    this.socket.onmessage = this.handleIncommingMessage;
    this.socket.onclose = this.handleSocketClose;
  }

  handleIncommingMessage = event => {
    const message = JSON.parse(event.data);
    this.props.actions.appendMessage(message);
  }

  handleSocketClose = event => {
    const { token } = this.props.state;
    if (!isValidToken(token)) {
      this.props.actions.unregisterSocketConnection();
    }
  }

  componentWillUpdate() {
    if (this.shouldConnectSocket()) {
      this.connectSocket();
    }
  }

  render() {
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
