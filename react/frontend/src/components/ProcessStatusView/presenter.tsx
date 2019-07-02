import * as React from 'react';
import { ProcessStatus } from '../../model/ProcessStatus';
import './ProcessStatusView.css';

class ProcessStatusView extends React.Component<any, any> {

    public render() {
        const processStatus = (this.props as any).processStatus as ProcessStatus;
        return (
            <div className="ProcessStatusView">
                {processStatus.realName}( 
                { processStatus.procId > 0 ? 
                    <span className="running">Running</span>: 
                    <span className="stopped">Stopped</span>
                })
            </div>
        );
    }
}

export default ProcessStatusView;
