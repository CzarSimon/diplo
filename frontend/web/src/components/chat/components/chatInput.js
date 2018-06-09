import React, { Component } from 'react';
import { Input, Button } from 'antd';
import { length, color } from '../../../style';

const style = {
  container: {
    position: 'fixed',
    bottom: '48px',
    borderTopStyle: 'solid',
    borderColor: color.border,
    borderWidth: '1px',
    width: '100%',
    backgroundColor: color.white,
    padding: length.small,
    display: 'flex',
  },
  input: {
    flex: '5',
    marginRight: length.mini
  },
  button: {
    flex: '1'
  }
}

export default class ChatInput extends Component {
  constructor(props) {
    super(props);
    this.state = {}
  }

  onSubmit = event => {
    this.props.handleSubmitMessage(this.state.message);
    this.setState({ message: null });
  }

  onChange = event => {
    this.setState({ message: event.target.value });
  }

  render() {
    return (
      <div style={style.container}>
        <Input
          placeholder={`Message #${this.props.channelName}`}
          type='text'
          value={this.state.message}
          onChange={this.onChange}
          style={style.input} />
        <Button type="primary" htmlType="submit" onClick={this.onSubmit} style={style.button}>
          Send
        </Button>
      </div>
    )
  }
}
