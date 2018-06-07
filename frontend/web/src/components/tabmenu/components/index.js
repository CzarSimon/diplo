import React, { Component } from 'react';
import { Menu, Icon } from 'antd';

const style = {
  container: {
    bottom: 0,
    position: 'fixed',
    width: '100%',
  },
  menu: {
    display: 'flex',
    justifyContent: 'space-around'
  }
}

export default class TabMenu extends Component {
  handleClick = event => {
    this.props.selectTab(event.key);
  }

  render() {
    const { selectedTab } = this.props;
    return (
      <div style={style.container}>
        <Menu onClick={this.handleClick} selectedKeys={[selectedTab]} mode="horizontal" style={style.menu}>
          <Menu.Item key="/"><Icon type="mail" />Home</Menu.Item>
          <Menu.Item key="/game/map"><Icon type="global" />Map</Menu.Item>
          <Menu.Item key="/game/chat"><Icon type="message" />Chat</Menu.Item>
          <Menu.Item key="/game/orders"><Icon type="bars" />Orders</Menu.Item>
        </Menu>
      </div>
    )
  }
}
