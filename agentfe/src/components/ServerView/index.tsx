import * as React from 'react';
import { ServerInfo } from '../../model/ServerInfo';
import ResourceStatusView from '../ResourceStatusView';

// interface IProps {
//   serverInfo: ServerInfo;
//   key:any;
//   // onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
// }

const renderResourceStatusView = (si: ServerInfo) => {
  return si.resourceStatuses.map((rs, idx) => (
    <ResourceStatusView resourceStatus={rs} key={idx} />
  ));
};

const ServerView = (props: any) => {

  const serverInfo = props.serverInfo;
  return (
    <div className="ServerView">
      Server Name - {serverInfo.name}<br />
      {renderResourceStatusView(serverInfo)}
    </div>
  );
};

export default ServerView;
