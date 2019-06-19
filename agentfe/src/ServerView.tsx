import * as React from 'react';
import { ServerInfo } from './model/ServerInfo';
import ResourceStatusView from './ResourceStatusView';

interface IProps {
  serverInfo: ServerInfo;
  // onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const ServerView = (props: IProps) => {

  const renderResourceStatusView = (si: ServerInfo) => {
    return si.resourceStatuses.map((rs) => (
      <ResourceStatusView resourceStatus={rs}/>
    ));
  }

  const serverInfo = props.serverInfo;
  return (
    <div className="ServerView">
      Server Name - {serverInfo.name}<br />
      {renderResourceStatusView(serverInfo)}
    </div>
  );
};

export default ServerView;
