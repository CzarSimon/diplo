import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { sendMessage } from '../../../ducks/messages';
import ChatInput from '../components/chatInput';

class ChatInputContainer extends Component {
  handleSubmitMessage = text => {
    const { channel, token } = this.props.state;
    this.props.actions.sendMessage(text, channel.id, token);
  }

  render() {
    const { channel } = this.props.state;
    return (
      <ChatInput
        channelName={channel.name}
        handleSubmitMessage={this.handleSubmitMessage} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    token: state.user.token,
    channel: state.channels[state.games.activeId].activeChannel
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ sendMessage }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(ChatInputContainer);
