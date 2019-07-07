import * as React from 'react';
import { ServerInfo } from 'src/model/ServerInfo';
import ProcessStatusContainerView from '../ProcessStatusContainerView';
import ResourceStatusContainerView from '../ResourceStatusContainerView';
import './ServerView.css';

// tslint:disable-next-line:interface-name
export interface ServerViewProps {
    serverInfo:ServerInfo;
    simpleMode:boolean;
}
class ServerView extends React.Component<ServerViewProps> {

    public shouldComponentUpdate(nextProps: ServerViewProps, nextStates: any):boolean {
        console.log("ServerView - shouldComponentUpdate true")
        return true;
    }

    public render() {
        return (
            <div className="ServerView">
                <div className="ServerName">
                    {this.props.serverInfo.name}
                </div>
                <ResourceStatusContainerView serverInfo={this.props.serverInfo} simpleMode={this.props.simpleMode}/>
                { !!this.props.serverInfo.processStatuses ? <ProcessStatusContainerView serverInfo={this.props.serverInfo} simpleMode={this.props.simpleMode}/>: ''}
            </div>
        );
    }
};

export default ServerView;