import * as React from 'react';
import ContextUtil from 'src/common/ui/util/ContextUtil';
import { ResourceStatus, VALUES_CNT } from '../../model/ResourceStatus';
import './ResourceStatusView.css'

// interface IProps {
//     resourceStatus: ResourceStatus;
//     // onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
// }

const WIDHT:number=150;
const HEIGHT:number=100;
const GRID:number=10;
const BACKGROUND_COLOR:string="#009900";
const GRID_COLOR:string="#000000";
const CHART_LINE_COLOR:string="#00ffff";
const CHART_BG_COLOR:string="#8ED6FF";
const CHART_ALPHA:number=0.5;
const FONT_NAME:string='20px san-serif';
const TEXT_HEIGHT:number=ContextUtil.measureFontHeight(FONT_NAME, '95%').height;

class ResourceStatusView extends React.Component<any, any> {

    private canvas: any;

    public shouldComponentUpdate(nextProps: any, nextState: any) {
        return this.props.tick !== nextProps.tick;
    }

    public render() {
        const resourceStatus = (this.props as any).resourceStatus as ResourceStatus;
        return (
            <div className="ResourceStatusView">
                {resourceStatus.name}<br />
                <canvas
                    ref={(canvas) => { this.canvas = canvas; }}
                    width={WIDHT}
                    height={HEIGHT} />
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
        ctx.globalAlpha = 1;
        
        // 1. background
        ctx.fillStyle = BACKGROUND_COLOR;
        ctx.clearRect(0, 0, WIDHT, HEIGHT);
        ContextUtil.drawRect(ctx, { x: 0, y: 0, w: WIDHT, h: HEIGHT });
        
        // 2. grid
        ctx.strokeStyle = GRID_COLOR;
        ctx.beginPath();
        for (let x=0-tick;x<WIDHT;x+=GRID) {
            if ( x>=0 ) {
                ctx.moveTo(x, 0);
                ctx.lineTo(x, HEIGHT);
            }
        }
        for (let y=0;y<HEIGHT;y+=GRID) {
            ctx.moveTo(0, y);
            ctx.lineTo(WIDHT, y);
        }
        ctx.stroke();
        
        // 3. draw charts
        const resourceStatus = (this.props as any).resourceStatus as ResourceStatus;
        const values = resourceStatus.values.slice();
        const valLen = values.length;
        if ( valLen > 0 ) {
            ctx.beginPath();
            ctx.strokeStyle = CHART_LINE_COLOR;
            ctx.moveTo(WIDHT, this.getYposition(values, 1));
            for ( let i=2; i<=values.length; i++ ) {
                const x = this.getXposition(i);
                const y = this.getYposition(values, i);
                ctx.lineTo(x, y);
            }
            ctx.lineTo(this.getXposition(values.length),HEIGHT);
            ctx.lineTo(WIDHT,HEIGHT);
            ctx.lineTo(WIDHT, this.getYposition(values, 1));

            ctx.closePath();
            ctx.stroke();

            ctx.globalAlpha = CHART_ALPHA;
            ctx.fillStyle=CHART_BG_COLOR;
            ctx.fill();

            // text
            ctx.globalAlpha = 1;
            ctx.fillStyle = "#ffffff";
            ctx.font =  FONT_NAME;
            ContextUtil.drawCenterText(ctx, TEXT_HEIGHT, "" + values[valLen-1] + "%");
        }
    }

    private getYposition(values:number[], pos:number):number {
        const yPixelDentisy = Math.floor(HEIGHT/100);
        const yPos = HEIGHT - (values[values.length-pos]*yPixelDentisy);
        return yPos;
    }

    private getXposition(pos:number):number {
        const xPixelDentisy = Math.floor(WIDHT/VALUES_CNT);
        const xPos = WIDHT - ((pos-1) * xPixelDentisy);
        return xPos;
    }
}

export default ResourceStatusView;
