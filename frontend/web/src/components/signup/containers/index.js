import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { signupUser } from '../../../ducks/user';
import Signup from '../components';

class SignupContainer extends Component {
  handleSignupSubmit = user => {
    this.props.actions.signupUser(user);
  }

  render() {
    return (
      <Signup handleSignupSubmit={this.handleSignupSubmit} />
    )
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ signupUser }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(SignupContainer);
