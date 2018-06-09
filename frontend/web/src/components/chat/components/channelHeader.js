import React, { Component } from 'react';
import MenuButtonContainer from '../containers/menuButton';
import { length, color } from '../../../style';

const style = {
  container: {
    backgroundColor: color.white,
    borderBottomStyle: 'solid',
    borderColor: color.border,
    borderWidth: '1px',
    position: 'fixed',
    width: '100%',
    display: 'flex'
  },
  text: {
    paddingTop: length.small,
    marginLeft: length.small,
    textAlign: 'left'
  },
  button: {
    marginTop: '1px'
  }
}

export default class ChannelHeader extends Component {
  render() {
    return (
      <div style={style.container}>
        <div style={style.button}>
          <MenuButtonContainer />
        </div>
        <h3 style={style.text}>{this.props.name}</h3>
      </div>
    )
  }
}
