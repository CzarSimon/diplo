import React, { Component } from 'react';
import { Divider, Collapse } from 'antd';
import _ from 'lodash';
import { length } from '../../../style';

const style = {
  container: {}
}

export default class LoggedInMenu extends Component {
  onClickLogout = event => {
    this.props.handleLogout();
    return false;
  }

  render() {
    const { games } = this.props;
    return (
      <div style={style.container}>
        <Collapse>
          <Collapse.Panel header='Active Games' key='1'>
            {_.values(games.all).map((game, i) => <p key={game.id}>{game.name}</p>)}
          </Collapse.Panel>
        </Collapse>
        <Divider />
        <p><a onClick={this.onClickLogout}>Logout</a></p>
      </div>
    )
  }
}
