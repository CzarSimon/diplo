import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { openChatMenu } from '../../../ducks/menu';
import MenuButton from '../components/menuButton';

class MenuButtonContainer extends Component {
  onClick = () => {
    this.props.actions.openChatMenu();
  }

  render() {
    return (
      <MenuButton onClick={this.onClick} />
    )
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ openChatMenu }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(MenuButtonContainer);
