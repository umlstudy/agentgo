import * as React from 'react';
import './App.css';

import SystemStausView from './components/SystemStatusView';
import logo from './logo.svg';
import { serverInfos2 } from './model/MockData';

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
                <SystemStausView serverInfos={serverInfos}/>
            </div>
        );
    }
}

export default App;
