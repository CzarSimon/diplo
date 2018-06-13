import React, { Component } from 'react';
import { length } from '../../../style';

const style = {
  container: {}
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
      <div style={style.container}>
        <p><a onClick={this.onClickLogin}>Login</a></p>
        <p><a onClick={this.onClickSignup}>Signup</a></p>
      </div>
    )
  }
}
