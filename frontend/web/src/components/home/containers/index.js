import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { logoutAndClearUser } from '../../../ducks/user';
import { isValidToken } from '../../../utils';
import Home from '../components';

class HomeContainer extends Component {
  handleLogout = () => {
    this.props.actions.logoutAndClearUser();
  }

  handleSelectPath = path => () => {
    browserHistory.replace(path);
  }

  goToLogin = this.handleSelectPath('/login');
  goToSignup = this.handleSelectPath('/signup');
  goToMyGames = this.handleSelectPath('/my/games');
  goToProfile = this.handleSelectPath('/my/profile');
  goToAbout = this.handleSelectPath('/about');
  goToGames = this.handleSelectPath('/games');

  render() {
    return (
      <Home
        {...this.props.state}
        goToLogin={this.goToLogin}
        goToSignup={this.goToSignup}
        handleLogout={this.handleLogout}
        goToMyGames={this.goToMyGames}
        goToProfile={this.goToProfile}
        goToAbout={this.goToAbout}
        goToGames={this.goToGames}
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
  actions: bindActionCreators({ logoutAndClearUser }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(HomeContainer);
