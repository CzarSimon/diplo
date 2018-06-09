import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { loadChannels } from '../../../ducks/channels';
import { openChatMenu } from '../../../ducks/menu';
import ChannelListContainer from './channelList';
import ChannelContainer from './channel';

class ChatContainer extends Component {
  componentDidMount() {
    const { token, gameId, chatMenuOpen } = this.props.state;
    const { loadChannels, openChatMenu } = this.props.actions;
    if (!token ) {
      browserHistory.replace('/login');
      return;
    }
    if (chatMenuOpen === null) {
      openChatMenu();
    }
    loadChannels(token, gameId);
  }

  render() {
    const { channels, chatMenuOpen } = this.props.state;
    const component = (channels && channels.channels)
      ? (chatMenuOpen) ? <ChannelListContainer /> : <ChannelContainer />
      : <div />
    return component
  }
}

const mapStateToProps = state => ({
  state: {
    token: state.user.token,
    gameId: state.games.activeId,
    channels: state.channels[state.games.activeId],
    chatMenuOpen: state.menu.chatMenuOpen
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ loadChannels, openChatMenu }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(ChatContainer);
