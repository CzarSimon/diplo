import React, { Component } from 'react';
import { Divider } from 'antd';


export default class LogoutOption extends Component {
  onClickLogout = event => {
    this.props.handleLogout();
    return false;
  }

  render() {
    if (this.props.isLoggedIn) {
      return (
        <div>
          <Divider />
          <h3><a onClick={this.onClickLogout}>Logout</a></h3>
        </div>
      )
    }
    return <div />
  }
}
