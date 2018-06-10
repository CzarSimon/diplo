import React, { Component } from 'react';
import { connect } from 'react-redux';
import ErrorMessage from '../../common/errorMessage';

class LoginErrorContainer extends Component {
  render() {
    const { error } = this.props.state;
    const component = ((error.isError)
      ? <ErrorMessage title='Login failed' text={error.message} />
      : <div />
    );
    return component;
  }
}

const mapStateToProps = state => ({
  state: {
    error: state.user.loginError
  }
});

export default connect(
  state => mapStateToProps(state)
)(LoginErrorContainer);
