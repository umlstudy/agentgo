import * as React from 'react';
import { ProcessStatus } from '../../model/ProcessStatus';
import './ProcessStatusView.css';

// tslint:disable-next-line:interface-name
export interface ProcessStatusViewProps {
    processStatus: ProcessStatus;
    key:number;
}

class ProcessStatusView extends React.Component<ProcessStatusViewProps> {

    public shouldComponentUpdate(nextProps: ProcessStatusViewProps) {
        console.log("ProcessStatusView - shouldComponentUpdate " + true);
        return true;
    }

    public render() {
        const processStatus = this.props.processStatus;
        return (
            <div className="ProcessStatusView">
                { processStatus.procId > 0 ? 
                    <div className="running">{processStatus.realName}</div>: 
                    <div className="stopped">{processStatus.realName}</div>
                }
            </div>
        );
    }
}

export default ProcessStatusView;
