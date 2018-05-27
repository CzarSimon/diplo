import React, { Component } from 'react';
import ChannelList from './components/chat/channels/main';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <ChannelList
          gameName='Congress In Vienna'
          gameId='game-1'
          userId='simon'/>
      </div>
    );
  }
}

export default App;
