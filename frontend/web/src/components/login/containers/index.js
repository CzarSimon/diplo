import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { loginUser } from '../../../ducks/user';
import Login from '../components';


class LoginContainer extends Component {
  handleLoginSubmit = (email, password) => {
    this.props.actions.loginUser(email, password);
  }

  render() {
    return (
      <Login handleLoginSubmit={this.handleLoginSubmit} />
    )
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ loginUser }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(LoginContainer);
