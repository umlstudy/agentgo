import * as React from 'react';
import './App.css';

import ServerView from './components/ServerView';
import logo from './logo.svg';
import { serverInfos2 } from './model/MockData';
import { ServerInfo } from './model/ServerInfo';

// react with canvas
// https://lavrton.com/using-react-with-html5-canvas-871d07d8d753/

class App extends React.Component {
  public render() {
    const serverInfos = serverInfos2;
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
          {this.renderServerViews(serverInfos)}
        </p>
      </div>
    );
  }
  
  private renderServerViews = (serverInfos: ServerInfo[]) => {
    return serverInfos.map((si) => (
      <ServerView serverInfo={si}/>
    ));
  }
}

export default App;
