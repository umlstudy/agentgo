import * as React from 'react';
import { ServerInfo } from 'src/model/ServerInfo';
import ResourceStatusView from '../ResourceStatusView';
import './ResourceStatusContainerView.css';

// tslint:disable-next-line: interface-name
export interface ResourceStatusContainerViewProps {
    serverInfo: ServerInfo;
}
class ResourceStatusContainerView extends React.Component<ResourceStatusContainerViewProps> {

    public shouldComponentUpdate(nextProps: ResourceStatusContainerViewProps, nextState: any) {
        return !nextState || nextState.resourceStatusesModified;
    }

    public render() {
        const serverInfo = this.props.serverInfo;
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