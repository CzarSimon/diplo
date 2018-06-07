import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { loadChannels } from '../../../ducks/channels';
import ChannelListContainer from './channelList';

class ChatContainer extends Component {
  componentDidMount() {
    const { token, gameId } = this.props.state;
    if (!token ) {
      browserHistory.replace('/login');
      return;
    }
    this.props.actions.loadChannels(token, gameId);
  }

  render() {
    const { channels } = this.props.state;
    const component = (channels && channels.channels)
      ? <ChannelListContainer />
      : <div />
    return component
  }
}

const mapStateToProps = state => ({
  state: {
    token: state.user.token,
    gameId: state.games.activeId,
    channels: state.channels[state.games.activeId]
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ loadChannels }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(ChatContainer);
