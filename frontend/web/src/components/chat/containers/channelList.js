import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { setActiveChannel, loadChannels } from '../../../ducks/channels';
import { closeChatMenu } from '../../../ducks/menu';
import ChannelList from '../components/channelList';

class ChannelListContainer extends Component {
  handleChannelSelect = channel => {
    const { actions, state } = this.props;
    actions.setActiveChannel(state.game.id, channel);
    actions.closeChatMenu();
  }

  render() {
    const { channels, activeChannel } = this.props.state.channels;
    return (
      <ChannelList
        channels={channels}
        activeChannel={activeChannel}
        handleChannelSelect={this.handleChannelSelect}
        gameName={'The Congress in Vienna'} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    game: state.games.all[state.games.activeId],
    channels: state.channels[state.games.activeId]
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({
    setActiveChannel,
    closeChatMenu,
    loadChannels
  }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(ChannelListContainer);
