import React, { Component } from 'react';
import { Icon } from 'antd'
import { length, color } from '../../../style'

const style = {
  container: {
    padding: length.small,
    backgroundColor: color.white
  },
  icon: {
    color: color.primary
  }
}

export default class MenuButton extends Component {
  render() {
    return (
      <div style={style.container} onClick={this.props.onClick}>
        <Icon style={style.icon} type="menu-unfold" />
      </div>
    )
  }
}
