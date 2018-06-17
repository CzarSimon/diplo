import React, { Component } from 'react';
import { Divider } from 'antd';
import { length } from '../../../style';

const style = {
  menuOption: {
    paddingBottom: length.medium
  }
}

export default class LoggedOutMenu extends Component {
  onClickLogin = event => {
    this.props.goToLogin();
    return false;
  }

  onClickSignup = event => {
    this.props.goToSignup();
    return false;
  }

  render() {
    return (
      <div>
        <h3 style={style.menuOption}><a onClick={this.onClickLogin}>Login</a></h3>
        <h3><a onClick={this.onClickSignup}>Sign up</a></h3>
        <Divider />
      </div>
    )
  }
}
