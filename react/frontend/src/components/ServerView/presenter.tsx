import * as React from 'react';
import { ServerInfo } from 'src/model/ServerInfo';
import ProcessStatusContainerView from '../ProcessStatusContainerView';
import ResourceStatusContainerView from '../ResourceStatusContainerView';
import './ServerView.css';

// tslint:disable-next-line:interface-name
export interface ServerViewProps {
    serverInfo:ServerInfo;
}
class ServerView extends React.Component<ServerViewProps> {

    public render() {
        return (
            <div className="ServerView">
                <div className="ServerName">
                    {this.props.serverInfo.name}
                </div>
                <ResourceStatusContainerView serverInfo={this.props.serverInfo}/>
                { !!this.props.serverInfo.processStatuses ? <ProcessStatusContainerView serverInfo={this.props.serverInfo}/>: ''}
            </div>
        );
    }
};

export default ServerView;