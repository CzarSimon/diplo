import React, { Component } from 'react';
import ChannelHeader from './channelHeader';
import ChatInputContainer from '../containers/chatInput';
import MessageList from './messageList';

const style = {
  container: {
    marginBottom: '100px'
  }
}

export default class Channel extends Component {
  render() {
    const { channel, messages } = this.props;
    return (
      <div style={style.container}>
        <ChannelHeader {...channel} />
        <MessageList messages={messages}/>
        <ChatInputContainer />
      </div>
    )
  }
}
