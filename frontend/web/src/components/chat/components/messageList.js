import React, { Component } from 'react';
import Message from './message';
import { length } from '../../../style';

const style = {
  container: {
    marginTop: length.navbar
  }
}

export default class MessageList extends Component {
  render() {
    const { messages } = this.props;
    return (
      <div>
        {messages.map((message, i) => (<Message key={i} {...message} />))}
      </div>
    )
  }
}
