import * as React from 'react';
import './App.css';

import SystemStausView from './components/SystemStatusView';

class App extends React.Component<any, any> {

    public render() {
        return (
            <div className="App">
                <SystemStausView/>
            </div>
        );
    }
}

export default App;
