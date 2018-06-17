import React, { Component } from 'react';
import Message from './message';
import ScrollToBottom from '../../common/scrollToBottom';
import { length } from '../../../style';

const style = {
  listPadding: {
    height: length.navbar
  }
}

export default class MessageList extends Component {
  render() {
    const { messages } = this.props;
    return (
      <div>
        <div style={style.listPadding}/>
        {messages.map((message, i) => (<Message key={i} {...message} />))}
        <ScrollToBottom />
      </div>
    )
  }
}
