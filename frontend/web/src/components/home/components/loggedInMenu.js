import React, { Component } from 'react';
import { Divider } from 'antd';
import { length } from '../../../style';

const style = {
  menuOption: {
    paddingBottom: length.medium
  }
}

export default class LoggedInMenu extends Component {

  onClickGames = event => {
    this.props.goToMyGames();
    return false;
  }

  onClickProfile = event => {
    this.props.goToProfile();
    return false;
  }

  render() {
    return (
      <div>
        <h3 style={style.menuOption}><a onClick={this.onClickGames}>My Games</a></h3>
        <h3><a onClick={this.onClickProfile}>Profile</a></h3>
        <Divider />
      </div>
    )
  }
}
