import React, { Component } from 'react';
import { Menu, Divider } from 'antd';
import _ from 'lodash';
import { length } from '../../../style';

const style = {
  header: {
    marginTop: length.large,
    marginBottom: length.large,
  },
  divider: {
    fontSize: '1em'
  }
}

export default class ChannelList extends Component {
  handleClick = event => {
    const { channels, handleChannelSelect } = this.props;
    handleChannelSelect(channels[event.key])
  }

  render() {
    const { channels, activeChannel, gameName } = this.props;
    const activeId = (activeChannel) ? activeChannel.id : null
    return (
      <div>
        <h3 style={style.header}>{gameName}</h3>
        <Divider style={style.divider}>Channels</Divider>
        <Menu onClick={this.handleClick} selectedKeys={[activeId]} mode="vertical">
          {_.values(channels).map((channel, i) => (
            <Menu.Item key={channel.id}>{`# ${channel.name}`}</Menu.Item>
          ))}
        </Menu>
      </div>
    )
  }
}
