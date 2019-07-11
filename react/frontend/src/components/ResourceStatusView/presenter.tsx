import * as React from 'react';
import ContextUtil from 'src/common/ui/util/ContextUtil';
import { GRAHP_VALUES_CNT } from 'src/Constants';
import { WarningLevel } from 'src/model/enums/WarningLevel';
import { ResourceStatus } from '../../model/ResourceStatus';
import './ResourceStatusView.css'

// tslint:disable-next-line:interface-name
export interface ResourceStatusViewProps {
    resourceStatus: ResourceStatus;
    tick:number;
    simpleMode:boolean;
}

const WIDHT:number=120;
const HEIGHT:number=40;
const GRID:number=10;
const BACKGROUND_COLOR:string="#009900";
const GRID_COLOR:string="#000000";
const CHART_LINE_COLOR:string="#00ffff";
const CHART_BG_COLOR:string="#8ED6FF";
const CHART_LINE_COLOR_ERROR:string="#dd2233";
const CHART_BG_COLOR_ERROR:string="#cc4455";
const CHART_LINE_COLOR_WARNING:string="#ddcc33";
const CHART_BG_COLOR_WARNING:string="#ccbb55";
const CHART_ALPHA:number=0.5;
const FONT_NAME:string='20px san-serif';
const TEXT_HEIGHT:number=ContextUtil.measureFontHeight(FONT_NAME, '95%').height;

class ResourceStatusView extends React.Component<ResourceStatusViewProps> {

    private canvas: any;

    public shouldComponentUpdate(nextProps: ResourceStatusViewProps) {
        // console.log("ResourceStatusView - shouldComponentUpdate " + (this.props.tick !== nextProps.tick))
        return this.props.tick !== nextProps.tick;
    }

    public render() {
        const resourceStatus = this.props.resourceStatus;
        if ( this.props.simpleMode ) {
            return (
                <div className="ResourceStatusView">
                    {resourceStatus.name}({resourceStatus.value}%)
                </div>
            );
        } else {
            return (
                <div className="ResourceStatusView">
                    {resourceStatus.name}<br />
                    <canvas className="canvasArea"
                        ref={(canvas) => { this.canvas = canvas; }}
                        width={WIDHT}
                        height={HEIGHT} />
                </div>
            );
        }
    }

    public componentDidMount() {
        this.updateCanvas();
    }

    public componentDidUpdate() {
        this.updateCanvas();
    }

    private updateCanvas() {
        if ( this.props.simpleMode ) {
            return;
        }
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
        ctx.lineWidth = 1;
        ctx.beginPath();
        for (let x=0-tick;x<WIDHT;x+=GRID) {
            if ( x>=0 ) {
                ctx.moveTo(x+0.5, 0);
                ctx.lineTo(x+0.5, HEIGHT);
            }
        }
        for (let y=0;y<HEIGHT;y+=GRID) {
            ctx.moveTo(0, y+0.5);
            ctx.lineTo(WIDHT, y+0.5);
        }
        ctx.stroke();
        
        // 3. draw charts
        const resourceStatus = this.props.resourceStatus;
        const values = resourceStatus.values.slice();
        const valLen = values.length;
        let chartLineColor = CHART_LINE_COLOR;
        let chartBgColor = CHART_BG_COLOR;
        const thisWl = this.props.resourceStatus.warningLevel;
        if ( !!thisWl ) {
            if ( thisWl === WarningLevel[WarningLevel.ERROR] ) {
                chartLineColor = CHART_LINE_COLOR_ERROR;
                chartBgColor = CHART_BG_COLOR_ERROR;
            } else if ( thisWl === WarningLevel[WarningLevel.WARNING] ) {
                chartLineColor = CHART_LINE_COLOR_WARNING;
                chartBgColor = CHART_BG_COLOR_WARNING;
            } 
        }
        if ( valLen > 0 ) {
            ctx.beginPath();
            ctx.strokeStyle = chartLineColor;
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
            ctx.fillStyle=chartBgColor;
            ctx.fill();

            // text
            ctx.globalAlpha = 1;
            ctx.fillStyle = "#ffffff";
            ctx.font =  FONT_NAME;
            ContextUtil.drawCenterText(ctx, TEXT_HEIGHT, "" + values[valLen-1] + "%");
        }
    }

    private getYposition(values:number[], pos:number):number {
        const yPixelDentisy = HEIGHT/100;
        const yPos = HEIGHT - (values[values.length-pos]*yPixelDentisy);
        return yPos;
    }

    private getXposition(pos:number):number {
        const xPixelDentisy = WIDHT/GRAHP_VALUES_CNT;
        const xPos = WIDHT - ((pos-1) * xPixelDentisy);
        return xPos;
    }
}

export default ResourceStatusView;
