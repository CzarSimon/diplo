import React, { Component } from 'react';
import { Card, Input, Icon, Button } from 'antd';
import LoginErrorContainer from '../containers/loginError';
import FullLogoWithTagline from '../../common/logo/fullWithTagline';
import { length } from '../../../style';
import { formButtonStyle } from '../../../style/buttons';

const style = {
  form: {
    margin: length.medium
  },
  input: {
    marginBottom: length.medium
  },
  formIcon: {
    color: 'rgba(0,0,0,.25)'
  },
  signupText: {
    textAlign: 'left',
    margin: '0',
    marginTop: length.small
  }
}

export default class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {}
  }

  onSubmit = event => {
    const { email, password } = this.state;
    this.props.handleLoginSubmit(email, password);
  }

  onChange = key => event => {
    this.setState({ [key]: event.target.value })
  }

  onEmailChange = this.onChange('email');

  onPasswordChange = this.onChange('password');

  onClickSignup = event => {
    this.props.handleClickSignup();
    return false;
  }

  render() {
    const { email, password } = this.state;
    return (
      <Card style={style.form}>
        <FullLogoWithTagline />
        <LoginErrorContainer />
        <Input
          placeholder='email'
          type='email'
          value={email}
          prefix={<Icon type="mail" style={style.formIcon} />}
          onChange={this.onEmailChange}
          style={style.input} />
        <Input
          placeholder='password'
          type='password'
          value={password}
          prefix={<Icon type="lock" style={style.formIcon} />}
          onChange={this.onPasswordChange}
          style={style.input}
          onSubmit={this.onSubmit} />
        <Button type="primary" htmlType="submit" onClick={this.onSubmit} style={formButtonStyle}>
          Log in
        </Button>
        <p style={style.signupText}>
          Or <a onClick={this.onClickSignup}>signup now!</a>
        </p>
      </Card>
    )
  }
}
