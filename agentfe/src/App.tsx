import * as React from 'react';
import './App.css';

import logo from './logo.svg';
import Server from './ServerView';
import ServerView from './ServerView';

// react with canvas
// https://lavrton.com/using-react-with-html5-canvas-871d07d8d753/

class App extends React.Component {
  public render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.tsx</code> and save to reload.
        </p>
        <p className="App-main">
          <ServerView/>
        </p>
      </div>
    );
  }
}

export default App;
