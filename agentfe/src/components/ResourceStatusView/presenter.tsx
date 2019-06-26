import * as React from 'react';
import { ResourceStatus } from '../../model/ResourceStatus';

// interface IProps {
//     resourceStatus: ResourceStatus;
//     // onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
// }

function rect(props: any) {
    const { ctx, x, y, width, height } = props;
    ctx.fillRect(x, y, width, height);
}

class ResourceStatusView extends React.Component<any, any> {

    private canvas: any;
    private values: number[];

    public shouldComponentUpdate(nextProps: any, nextState: any) {
        return this.props.tick !== nextProps.tick;
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
        let tick = 0;
        if ( this.props.tick !== undefined ) {
            tick = this.props.tick%10*2;
        }
        const ctx = this.canvas.getContext('2d');
        ctx.fillStyle = "#FF0000";
        ctx.clearRect(0, 0, 800, 300);
        rect({ ctx, x: 0, y: 0, width: 800, height: 300 });
        // draw children “components”
        ctx.fillStyle = "#000000";
        ctx.beginPath();
        for (let x=0-tick;x<800;x+=20) {
            if ( x>=0 ) {
                ctx.moveTo(x, 0);
                ctx.lineTo(x, 300);
            }
        }
        for (let y=0;y<300;y+=20) {
            ctx.moveTo(0, y);
            ctx.lineTo(800, y);
        }
        ctx.stroke();
        
        this.values.map((val) => {
            // tslint:disable-next-line:no-console
            console.log(val);
        }
        );
    }
}

export default ResourceStatusView;
