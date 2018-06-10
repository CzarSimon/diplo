import React, { Component } from 'react';
import { connect } from 'react-redux';
import ErrorMessage from '../../common/errorMessage';

class SignupErrorContainer extends Component {
  render() {
    const { error } = this.props.state;
    const component = ((error.isError)
      ? <ErrorMessage title='Signup failed' text={error.message} />
      : <div />
    );
    return component;
  }
}

const mapStateToProps = state => ({
  state: {
    error: state.user.signupError
  }
});

export default connect(
  state => mapStateToProps(state)
)(SignupErrorContainer);
