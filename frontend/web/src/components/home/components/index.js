import React, { Component } from 'react';
import { Divider } from 'antd';
import FullLogo from '../../common/logo/full';
import LoggedInMenu from './loggedInMenu';
import LoggedOutMenu from './loggedOutMenu';
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
    const { isLoggedIn, handleLogout, games, goToLogin, goToSignup } = this.props;
    const userComponent = (isLoggedIn)
      ? <LoggedInMenu handleLogout={handleLogout} games={games} />
      : <LoggedOutMenu goToLogin={goToLogin} goToSignup={goToSignup} />
    return (
      <div style={style.container}>
        <div style={style.logo}>
          <FullLogo />
        </div>
        <Divider />
        {userComponent}
      </div>
    )
  }
}
