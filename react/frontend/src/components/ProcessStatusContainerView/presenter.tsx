import * as React from 'react';
import { ServerInfo } from '../../model/ServerInfo';
import ProcessStatusView from '../ProcessStatusView';
import './ProcessStatusContainerView.css';

class ProcessStatusContainerView extends React.Component<any, any> {

    public shouldComponentUpdate(nextProps: any, nextState: any) {
        return !nextState || nextState.processStatusesModified;
    }

    public render() {
        const serverInfo = this.props.serverInfo as ServerInfo;
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