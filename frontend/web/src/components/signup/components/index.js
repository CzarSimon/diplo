import React, { Component } from 'react';
import { Card, Input, Icon, Button } from 'antd';
import { length } from '../../../style'

const style = {
  form: {
    padding: length.small,
  },
  input: {
    marginBottom: length.small
  },
  formIcon: {
    color: 'rgba(0,0,0,.25)'
  },
  button: {
    marginBottom: length.small,
    width: '100%'
  }
}

export default class Signup extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  onSubmit = event => {
    const { email, username, password, givenName, surname } = this.state;
    this.props.handleSignupSubmit({ email, username, password, givenName, surname });
  }

  onChange = key => event => {
    this.setState({ [key]: event.target.value })
  }

  onEmailChange = this.onChange('email');
  onPasswordChange = this.onChange('password');
  onUsernameChange = this.onChange('username');
  onGivenNameChange = this.onChange('givenName');
  onSurnameChange = this.onChange('surname');

  render() {
    const { email, username, password, givenName, surname } = this.state;
    return (
      <Card style={style.form}>
        <h2>Diplomacy</h2>
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
        <Button type="primary" htmlType="submit" onClick={this.onSubmit} style={style.button}>
          Sign up
        </Button>
      </Card>
    )
  }
}
