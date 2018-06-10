import React, { Component } from 'react';
import logo from '../../../resources/diplo-logo-full-with-tagline.png';
import { length } from '../../../style';

const style = {
  width: '100%',
  marginBottom: length.large
}

export default class FullLogoWithTagline extends Component {
  render() {
    return (
      <img
        src={logo}
        style={style}
        alt='Diplomacy. The strong do what they can. The weak suffer what they must' />
    )
  }
}
