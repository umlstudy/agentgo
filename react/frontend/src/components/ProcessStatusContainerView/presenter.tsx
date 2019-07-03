import * as React from 'react';
import { ServerInfo } from '../../model/ServerInfo';
import ProcessStatusView from '../ProcessStatusView';
import './ProcessStatusContainerView.css';

// tslint:disable-next-line:interface-name
export interface ProcessStatusContainerViewProps {
    serverInfo: ServerInfo;
}
class ProcessStatusContainerView extends React.Component<ProcessStatusContainerViewProps> {

    public shouldComponentUpdate(nextProps: ProcessStatusContainerViewProps, nextState: any) {
        return !nextState || nextState.processStatusesModified;
    }

    public render() {
        const serverInfo = this.props.serverInfo;
        return (
        <div>
            {serverInfo.processStatuses.map((ps, idx) => (
                <ProcessStatusView processStatus={ps} key={idx} />
            ))}
        </div>
        );
    }
}


export default ProcessStatusContainerView;