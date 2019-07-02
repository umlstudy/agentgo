import * as React from 'react';
import ProcessStatusContainerView from '../ProcessStatusContainerView/presenter';
import ResourceStatusContainerView from '../ResourceStatusContainerView/presenter';
import './ServerView.css';

class ServerView extends React.Component<any, any> {

    public render() {
        const serverInfo = this.props.serverInfo;
        return (
            <div className="ServerView">
                <div className="ServerName">
                    {serverInfo.name}
                </div>
                <ResourceStatusContainerView serverInfo={serverInfo}/>
                { !!serverInfo.processStatuses ? <ProcessStatusContainerView serverInfo={serverInfo}/>: ''}
            </div>
        );
    }
};

export default ServerView;