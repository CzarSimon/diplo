import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { loadMessages } from '../../../ducks/messages';
import Channel from '../components/channel';

class ChannelContainer extends Component {
  componentDidMount() {
    const { token, channel } = this.props.state;
    this.props.actions.loadMessages(token, channel.id);
  }

  render() {
    const { channel, messages : allMessages } = this.props.state;
    const messages = allMessages[channel.id]
    const component = (messages && messages.length)
      ? <Channel channel={channel} messages={messages} /> : <div />
    return component
  }
}

const mapStateToProps = state => ({
  state: {
    token: state.user.token,
    channel: state.channels[state.games.activeId].activeChannel,
    messages: state.messages,
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ loadMessages }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(ChannelContainer);
