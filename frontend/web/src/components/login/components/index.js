import React, { Component } from 'react';
import { Card, Input, Icon, Button } from 'antd';
import { length } from '../../../style';

const style = {
  form: {
    margin: length.medium,
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

  render() {
    const { email, password } = this.state;
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
          style={style.input}
          onSubmit={this.onSubmit} />
        <Button type="primary" htmlType="submit" onClick={this.onSubmit} style={style.button}>
          Log in
        </Button>
      </Card>
    )
  }
}
