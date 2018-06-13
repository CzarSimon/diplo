import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { logoutUser } from '../../../ducks/user';
import { isValidToken } from '../../../utils';
import Home from '../components';

class HomeContainer extends Component {
  handleLogout = () => {
    this.props.actions.logoutUser();
  }

  handleSelectPath = path => () => {
    browserHistory.replace(path);
  }

  goToLogin = this.handleSelectPath('/login');
  goToSignup = this.handleSelectPath('/signup');

  render() {
    return (
      <Home
        {...this.props.state}
        goToLogin={this.goToLogin}
        goToSignup={this.goToSignup}
        handleLogout={this.handleLogout}
        isLoggedIn={isValidToken(this.props.state.token)} />
    )
  }
}

const mapStateToProps = state => ({
  state: {
    games: state.games,
    token: state.user.token,
    user: state.user.info
  }
})

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ logoutUser }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(HomeContainer);
