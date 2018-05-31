import React, { Component } from 'react';
import { getGameChannels } from '../api';
import Channel from './channel';
import { color, length } from '../../../style'

const style = {
  title: {
    padding: length.small,
    marginBottom: length.medium,
    color: color.white,
    display: 'flex'
  },
  list: {
    padding: length.medium,
    backgroundColor: color.blue,
    width: '20%',
    height: '100%'
  }
}

export default class ChannelList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      channels: []
    }
  }

  componentDidMount() {
    const { token, gameId } = this.props;
    getGameChannels(token, gameId).then(channels => {
      this.setState({channels})
    })
  }

  render() {
    const { channels } = this.state
    return (
      <div style={style.list}>
        <h2 style={style.title}>{this.props.gameName}</h2>
        {channels.map((channel, i) => <Channel key={i} {...channel} />)}
      </div>
    )
  }
}
