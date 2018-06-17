import React, { Component } from 'react'


export default class ScrollToBottom extends Component {
  componentDidMount() {
    this.scrollToBottom();
  }

  componentDidUpdate() {
    this.scrollToBottom();
  }

  scrollToBottom = () => {
    this.messagesEnd.scrollIntoView({ behavior: 'auto' });
  }

  render() {
    return (
      <div ref={(el) => { this.messagesEnd = el; }}/>
    );
  }

}
