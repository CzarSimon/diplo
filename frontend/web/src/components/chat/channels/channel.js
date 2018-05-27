import React, { Component } from 'react';
import { color, length } from '../../../style'

const style = {
  channel: {
    padding: length.small,
    color: color.white,
    display: 'flex'
  },
  text: {
    marginTop: 0,
    marginBottom: 0
  }
}

export default class Channel extends Component {

  onClick = () => {
    console.log(this.props);
  }

  render() {
    return (
      <div style={style.channel} onClick={this.onClick}>
        <p style={style.text}>{this.props.name}</p>
      </div>
    )
  }
}
