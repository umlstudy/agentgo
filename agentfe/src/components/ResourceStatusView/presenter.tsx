import * as React from 'react';
import { ResourceStatus } from '../../model/ResourceStatus';

interface IProps {
    resourceStatus: ResourceStatus;
    // onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

function rect(props: any) {
    const { ctx, x, y, width, height } = props;
    ctx.fillRect(x, y, width, height);
}

class ResourceStatusView extends React.Component<IProps, {}> {

    private canvas: any;
    private values: number[];

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
        this.values = resourceStatus.values;
        return (
            <div className="ResourceStatus">
                Resource Name - {resourceStatus.name}<br />
                <canvas
                    ref={(canvas) => { this.canvas = canvas; }}
                    width={800}
                    height={100} />
            </div>
        );
    }

    public componentDidMount() {
        this.updateCanvas();
    }
    public componentDidUpdate() {
        this.updateCanvas();
    }

    private updateCanvas() {
        const ctx = this.canvas.getContext('2d');
        ctx.clearRect(0, 0, 800, 300);
        // draw children “components”
        rect({ ctx, x: 10, y: 10, width: 550, height: 50 });
        rect({ ctx, x: 110, y: 30, width: 50, height: 50 });
        
        this.values.map((val) => {
            // tslint:disable-next-line:no-console
            console.log(val);
        }
        );
    }
}

export default ResourceStatusView;
