import React, { Component } from 'react';
import moment from 'moment';
import { Divider } from 'antd';
import { length } from '../../../style';

const style = {
  content: {
    paddingLeft: length.small,
    paddingRight: length.small,
    textAlign: 'left'
  }
}

export default class Message extends Component {
  formatDate = date => {
    return moment(date).format('YYYY-MM-DD hh:mm');
  }

  render() {
    const { text, authorId, createdAt } = this.props;
    return (
      <div>
        <div  style={style.content}>
          <div>
            <h4>{authorId}</h4>
            <p>{this.formatDate(createdAt)}</p>
          </div>
          <p>{text}</p>
        </div>
        <Divider />
      </div>
    )
  }
}
