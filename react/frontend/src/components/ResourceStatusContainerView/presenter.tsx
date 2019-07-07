import * as React from 'react';
import { ServerInfo } from 'src/model/ServerInfo';
import ResourceStatusView from '../ResourceStatusView';
import './ResourceStatusContainerView.css';

// tslint:disable-next-line: interface-name
export interface ResourceStatusContainerViewProps {
    serverInfo: ServerInfo;
    simpleMode: boolean;
}
class ResourceStatusContainerView extends React.Component<ResourceStatusContainerViewProps> {

    public shouldComponentUpdate(nextProps: ResourceStatusContainerViewProps) {
        if ( this.props.simpleMode !== nextProps.simpleMode ) {
            console.log("ResourceStatusContainerView - shouldComponentUpdate true")
            return true;
        }
        console.log("ResourceStatusContainerView - shouldComponentUpdate " + nextProps.serverInfo.resourceStatusesModified)
        return nextProps.serverInfo.resourceStatusesModified;
    }

    public render() {
        const serverInfo = this.props.serverInfo;
        return (
        <div>
            {serverInfo.resourceStatuses.map((rs, idx) => (
                <ResourceStatusView resourceStatus={rs} key={idx} simpleMode={this.props.simpleMode} />
            ))}
        </div>
        );
    }
};

export default ResourceStatusContainerView;