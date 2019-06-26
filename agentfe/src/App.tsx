import * as React from 'react';
import './App.css';

import ServerView from './components/ServerView';
import logo from './logo.svg';
import { serverInfos2 } from './model/MockData';
import { ServerInfo } from './model/ServerInfo';

class App extends React.Component<any, any> {

    public render() {
        const prop = (this as any).props;
        console.log(prop.hello);
        // prop.increment();
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
                <div className="App-main">
                    {this.renderServerViews(serverInfos)}
                </div>
            </div>
        );
    }

    private renderServerViews = (serverInfos: ServerInfo[]) => {
        return serverInfos.map((value: ServerInfo, index: number, array: ServerInfo[]) => (
            <ServerView serverInfo={value} key={index} keyValue={index} />
        ));
    }
}

export default App;
