import React, { Component } from 'react';
import { Divider } from 'antd';
import FullLogo from '../../common/logo/full';
import LoggedInMenu from './loggedInMenu';
import LoggedOutMenu from './loggedOutMenu';
import GeneralMenu from './generalMenu';
import LogoutOption from './logoutOption';
import { length } from '../../../style';

const style = {
  container: {
    margin: length.medium,
    display: 'flex',
    flexDirection: 'column'
  },
  logo: {
    marginTop: length.large,
    marginLeft: '10%',
    width: '80%'
  }
}

export default class Home extends Component {
  render() {
    const {
      isLoggedIn, handleLogout, goToLogin, goToSignup,
      goToMyGames, goToProfile, goToAbout, goToGames
    } = this.props;
    const userComponent = (isLoggedIn)
      ? <LoggedInMenu goToMyGames={goToMyGames} goToProfile={goToProfile} />
      : <LoggedOutMenu goToLogin={goToLogin} goToSignup={goToSignup} />
    return (
      <div style={style.container}>
        <div style={style.logo}>
          <FullLogo />
        </div>
        <Divider />
        {userComponent}
        <GeneralMenu goToAbout={goToAbout} goToGames={goToGames} />
        <LogoutOption isLoggedIn={isLoggedIn} handleLogout={handleLogout} />
      </div>
    )
  }
}
