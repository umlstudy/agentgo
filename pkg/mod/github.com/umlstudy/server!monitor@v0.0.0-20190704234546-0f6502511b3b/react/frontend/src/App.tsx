import * as React from 'react';
import './App.css';

import SystemStausView from './components/SystemStatusView';

class App extends React.Component<any, any> {

    public render() {
        return (
            <div className="App">
                <div className="App-header">
                    <div className="App-title">Server Monitor</div>
                </div>
                <SystemStausView/>
            </div>
        );
    }
}

export default App;
