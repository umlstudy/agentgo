import * as React from 'react';
import { ProcessStatus } from '../../model/ProcessStatus';
import './ProcessStatusView.css';

// tslint:disable-next-line:interface-name
export interface ProcessStatusViewProps {
    processStatus: ProcessStatus;
    key:number;
}

class ProcessStatusView extends React.Component<ProcessStatusViewProps> {

    public render() {
        const processStatus = this.props.processStatus;
        return (
            <div className="ProcessStatusView">
                { processStatus.procId > 0 ? 
                    <span className="running">{processStatus.realName}</span>: 
                    <span className="stopped">{processStatus.realName}</span>
                }
            </div>
        );
    }
}

export default ProcessStatusView;
