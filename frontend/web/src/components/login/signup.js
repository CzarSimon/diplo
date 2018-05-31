import React, { Component } from 'react';
import { color, length } from '../../../style'

const style = {
  form: {
    padding: length.small,
    color: color.white,
    display: 'flex'
  },
  text: {
    padding: length.small,
    marginBottom: length.medium,
    color: color.blue,
    display: 'flex'
  },
  input: {
    paddingTop: length.mini,
    paddingLeft: length.small,
    paddingRight: length.small,
    marginLeft: length.small,
    marginRight: length.small,
    backgroundColor: color.grey.medium,
    fontSize: font.size.small,
    color: color.white,
    border: 'solid',
    borderWidth: '1px',
    outline: 'none'
  }
}

export default class SignupForm extends Component {
  constructor(props) {
    super(props)
    this.state = {
      email: '',
      username: '',
      password: '',
      givenName: '',
      surname: '',
    }
  }

  updateValue = key => event => {
    this.setState({
      [key]: value
    })
  }

  render() {
    const {email, username, password, givenName, surname } = this.state;
    return (
      <div style={style.channel}>
        <h3 style={style.text}>{this.props.name}</h3>
        <TextInput value={email} inputType='email' placeholder='Email' handleChange=this.updateValue('email') />
        <TextInput value={password} inputType='password' placeholder='Password' handleChange=this.updateValue('password') />
        <TextInput value={username} inputType='text' placeholder='username' handleChange=this.updateValue('username') />
        <TextInput value={givenName} inputType='text' placeholder='First name' handleChange=this.updateValue('givenName') />
        <TextInput value={surname} inputType='text' placeholder='Surname' handleChange=this.updateValue('surname') />
        <button>signup</button>
      </div>
    )
  }
}

class TextInput extends Component {
  render() {
    const { value, inputType, placeholder, handleChange } = this.props;
    return (
      <input
        type={inputType}
        style={style.input}
        value={value}
        onChange={handleQuery}
        placeholder={placeholder} />
    )
  }
}
