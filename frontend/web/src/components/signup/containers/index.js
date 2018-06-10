import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import validator from 'validator';
import { signupUser, setSignupError } from '../../../ducks/user';
import Signup from '../components';

class SignupContainer extends Component {
  handleSignupSubmit = user => {
    if (this.validateInput(user)) {
      user = {
        ...user,
        email: validator.normalizeEmail(user.email)
      }
      this.props.actions.signupUser(user);
    }
  }

  validateInput = ({ email, password }) => {
    const { setSignupError } = this.props.actions;
    if (!email || !validator.isEmail(email)) {
      setSignupError(`Invalid email: ${email}`);
      return false;
    }
    if (!password || password.length < 10) {
      setSignupError('Password must be more than 10 characters');
      return false;
    }
    return true;
  }

  handleClickLogin = () => {
    browserHistory.replace('/login');
  }

  render() {
    return (
      <Signup
        handleSignupSubmit={this.handleSignupSubmit}
        handleClickLogin={this.handleClickLogin} />
    )
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ signupUser, setSignupError }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(SignupContainer);
