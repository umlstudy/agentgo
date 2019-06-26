import * as React from 'react';
import { ResourceStatus } from '../../model/ResourceStatus';

interface IProps {
  resourceStatus: ResourceStatus;
  // onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

class ResourceStatusView extends React.Component<IProps, {}> {

  constructor(props: IProps) {
    super(props);
  }

  public shouldComponentUpdate(nextProps: any, nextState: any) {
    return (
      nextProps.resourceStatus !== (this.props as any).resourceStatus
    );
  }

  public render() {
    const resourceStatus = (this.props as any).resourceStatus as ResourceStatus;
    return (
      <div className="ResourceStatus">
        Resource Name - {resourceStatus.name}<br />
        {this.renderResourceGraph(resourceStatus)}
      </div>
    );
  }

  private renderResourceGraph = (resourceStatus: ResourceStatus) => {
    return resourceStatus.values.map((val) => (
      <h2>{val}</h2>
    ));
  }
}

export default ResourceStatusView;
