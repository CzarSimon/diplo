import React, { Component } from 'react';
import { Alert } from 'antd';
import { length } from '../../style';

const style = {
  marginBottom: length.medium,
  textAlign: 'left'
}

export default class ErrorMessage extends Component {
  render() {
    return (
      <Alert
        message={this.props.title}
        description={this.props.text}
        type="error"
        style={style}
        showIcon
        closable />
    )
  }
}
