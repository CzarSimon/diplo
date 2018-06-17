import React, { Component } from 'react';
import icon from '../../../resources/diplo-icon.png';

const style = {
  width: '100%'
}

export default class DiploIcon extends Component {
  render() {
    return (
      <img
        src={icon}
        style={style}
        alt='Diplomacy. The strong do what they can. The weak suffer what they must' />
    )
  }
}
