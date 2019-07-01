import * as React from 'react';
import { ServerInfo } from '../../model/ServerInfo';
import ResourceStatusView from '../ResourceStatusView';
import './ServerView.css';

const renderResourceStatusView = (si: ServerInfo) => {
    return si.resourceStatuses.map((rs, idx) => (
        <ResourceStatusView resourceStatus={rs} key={idx} />
    ));
};

const renderProcessStatusView = (si: ServerInfo) => {
    return si.processStatuses.map((ps, idx) => (
        <div>{ps.name}</div>
    ));
};

const ServerView = (props: any) => {
    const serverInfo = props.serverInfo;
    const st = React.useState(
      );
    console.log(st);
    return (
        <div className="ServerView">
            <div className="ServerName">
                {serverInfo.name}
            </div>
            <div>
                {renderResourceStatusView(serverInfo)}
            </div>
            { !!serverInfo.processStatuses ? <div>{renderProcessStatusView(serverInfo)}</div>: ''}
        </div>
    );
};

export default ServerView;