import * as React from 'react';
import { ServerInfo } from '../../model/ServerInfo';
import ResourceStatusView from '../ResourceStatusView';

const renderResourceStatusView = (si: ServerInfo) => {
    return si.resourceStatuses.map((rs, idx) => (
        <ResourceStatusView resourceStatus={rs} key={idx} />
    ));
};

const ServerView = (props: any) => {
    const serverInfo = props.serverInfo;
    return (
        <div className="ServerView">
            Server Name - {serverInfo.name} {props.num}<br />
            {renderResourceStatusView(serverInfo)}
        </div>
    );
};

export default ServerView;