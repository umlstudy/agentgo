import * as React from 'react';
import ResourceStatusView from './ResourceStatusView';

class ServerView extends React.Component {
  public render() {
    return (
      <div className="Server">
          Server Name - mac.sejong.asia<br/>
          <ResourceStatusView/>
      </div>
    );
  }
}

export default ServerView;
