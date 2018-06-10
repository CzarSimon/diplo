import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import validator from 'validator';
import { loginUser, setLoginError } from '../../../ducks/user';
import Login from '../components';

class LoginContainer extends Component {
  handleLoginSubmit = (email, password) => {
    const { setLoginError, loginUser } = this.props.actions;
    if (!email || !validator.isEmail(email)) {
      setLoginError(`Invalid email: ${email}`);
      return;
    }
    loginUser(validator.normalizeEmail(email), password);
  }

  handleClickSignup = () => {
    browserHistory.replace('/signup');
  }

  render() {
    return (
      <Login
        handleLoginSubmit={this.handleLoginSubmit}
        handleClickSignup={this.handleClickSignup} />
    )
  }
}

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ loginUser, setLoginError }, dispatch)
});

export default connect(
  state => ({}),
  dispatch => mapDispatchToActions(dispatch)
)(LoginContainer);
