import * as React from 'react';
import { ServerInfo } from 'src/model/ServerInfo';
import ResourceStatusView from '../ResourceStatusView';
import './ResourceStatusContainerView.css';

class ResourceStatusContainerView extends React.Component<any, any> {

    public shouldComponentUpdate(nextProps: any, nextState: any) {
        return !nextState || nextState.resourceStatusesModified;
    }

    public render() {
        const serverInfo = this.props.serverInfo as ServerInfo;
        return (
        <div>
            {serverInfo.resourceStatuses.map((rs, idx) => (
                <ResourceStatusView resourceStatus={rs} key={idx} />
            ))}
        </div>
        );
    }
};

export default ResourceStatusContainerView;