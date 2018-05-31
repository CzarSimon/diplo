import React, { Component } from 'react';
import ChannelList from './components/chat/channels/main';
import SignupForm from './components/login/signup';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      token: '',
      gameId: 'game-1',
      gameName: 'Congress In Vienna'
    }
  }

  loggedIn = () => this.state.token !== '';

  setToken = token => {
    this.setState({ token })
  }

  render() {
    const { token, gameId, gameName } = this.state;
    return (
      <div className="App">
        {
          (this.loggedIn()) ?
          (<ChannelList gameName={gameName} gameId={gameId} token={token} />) :
          (<div><SignupForm setToken={this.setToken}/></div>)
        }
      </div>
    );
  }
}

export default App;
