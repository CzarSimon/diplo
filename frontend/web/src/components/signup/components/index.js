import React, { Component } from 'react';
import { Card, Input, Icon, Button } from 'antd';
import SignupErrorContainer from '../containers/signupError';
import { length, color } from '../../../style'
import { formButtonStyle } from '../../../style/buttons';

const style = {
  form: {
    margin: length.medium,
  },
  input: {
    marginBottom: length.medium
  },
  formIcon: {
    color: color.border
  },
  loginText: {
    textAlign: 'left',
    margin: '0',
    marginTop: length.small
  }
}

export default class Signup extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  onSubmit = event => {
    const { email, username, password, givenName, surname } = this.state;
    this.props.handleSignupSubmit({
      email,
      username,
      password,
      givenName,
      surname
    });
  }

  onChange = key => event => {
    this.setState({ [key]: event.target.value })
  }

  onEmailChange = this.onChange('email');
  onPasswordChange = this.onChange('password');
  onUsernameChange = this.onChange('username');
  onGivenNameChange = this.onChange('givenName');
  onSurnameChange = this.onChange('surname');

  onClickLogin = event => {
    this.props.handleClickLogin();
    return false;
  }

  render() {
    const { email, username, password, givenName, surname } = this.state;
    return (
      <Card style={style.form}>
        <h2>Diplomacy</h2>
        <SignupErrorContainer />
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
          style={style.input} />
        <Input
          placeholder='username'
          value={username}
          prefix={<Icon type="user" style={style.formIcon} />}
          onChange={this.onUsernameChange}
          style={style.input} />
        <Input
          placeholder='Given name'
          value={givenName}
          prefix={<Icon type="user" style={style.formIcon} />}
          onChange={this.onGivenNameChange}
          style={style.input} />
        <Input
          placeholder='Surname'
          value={surname}
          prefix={<Icon type="user" style={style.formIcon} />}
          onChange={this.onSurnameChange}
          style={style.input} />
        <Button type="primary" htmlType="submit" onClick={this.onSubmit} style={formButtonStyle}>
          Sign up
        </Button>
        <p style={style.loginText}>
          Already have an account? <a onClick={this.onClickLogin}>Login here</a>
        </p>
      </Card>
    )
  }
}
