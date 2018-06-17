import React, { Component } from 'react';
import { length } from '../../../style';

const style = {
  menuOption: {
    paddingBottom: length.medium
  }
}

export default class GeneralMenu extends Component {

  onClickGames = event => {
    this.props.goToGames();
    return false;
  }

  onClickAbout = event => {
    this.props.goToAbout();
    return false;
  }

  render() {
    return (
      <div>
        <h3 style={style.menuOption}><a onClick={this.onClickGames}>View Games</a></h3>
        <h3><a onClick={this.onClickAbout}>About</a></h3>
      </div>
    );
  }

}
