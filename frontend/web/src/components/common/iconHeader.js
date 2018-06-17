import React from 'react';
import DiploIcon from './logo/icon';
import { length } from '../../style';

const style = {
  container: {
    margin: length.medium,
    display: 'flex',
    alignItems: 'center'
  },
  logo: {
    width: '20%'
  },
  text: {
    textAlign: 'left',
    margin: '0',
    marginLeft: length.medium
  }
}

const IconHeader = props => (
  <div style={style.container}>
    <div style={style.logo}>
      <DiploIcon />
    </div>
    <h1 style={style.text}>{props.text}</h1>
  </div>
);

export default IconHeader;
