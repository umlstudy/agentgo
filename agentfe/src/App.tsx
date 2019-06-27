import * as React from 'react';
import './App.css';

import SystemStausView from './components/SystemStatusView';

class App extends React.Component<any, any> {

    public render() {
        return (
            <div className="App">
                <header className="App-header">
                    <h1 className="App-title">Server Monitor</h1>
                </header>
                <SystemStausView/>
            </div>
        );
    }
}

export default App;
