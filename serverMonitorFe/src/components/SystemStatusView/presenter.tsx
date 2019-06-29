import * as React from 'react';
import ArrayUtil from 'src/common/util/ArrayUtil';
import { ServerInfo, ServerInfoMap } from 'src/model/ServerInfo';
import { StoreObject } from 'src/model/StoreObject';
import ServerView from '../ServerView/presenter';

class SystemStausView extends React.Component<any, any> {

    public static getDerivedStateFromProps(nextProps: any, prevState: any) {
        if ( prevState.isRunning == null ) {
            const timerInterval = setInterval(() => {
                nextProps.tick();
                nextProps.request();
            }, 1000);
            return {
                ...nextProps,
                isRunning:true,
                timerInterval
            }
        }

        return null;
    }

    public state:any = {
        isRunning:null
    }

    public getSnapshotBeforeUpdate(prevProps:any, prevState:any) {
        if (this.props.isRunning !== prevProps.isRunning ) {
            return true;
        } 
        const changed = this.isChanged(this.props.serverInfoMap, prevProps.serverInfoMap );
        if ( changed ) {
            return true;
        }
        return false;
    }

    public componentDidUpdate(prevProps:any, prevState:any) {
        return true;
    }

    public shouldComponentUpdate(nextProps: any, nextState: any) {
        return this.props.tick !== nextProps.tick;
    }

    public render() {
        const props = (this as any).props as StoreObject;
        const serverInfoMap = props.serverInfoMap;
        const serverInfos:ServerInfo[] = ArrayUtil.json2Array(serverInfoMap);
        return (
            <div>
                {this.renderServerViews(serverInfos)}
            </div>
        );
    }

    private renderServerViews = (serverInfos: ServerInfo[]) => {
        return serverInfos.map((value: ServerInfo, index: number, array: ServerInfo[]) => (
            <ServerView serverInfo={value} key={index} keyValue={index} />
        ));
    }

    private findServerInfo(sis:ServerInfo[], siId:string):ServerInfo|null {
        // tslint:disable-next-line:prefer-for-of
        for ( let i=0;i<sis.length;i++ ) {
            const si = sis[i];
            if ( si.id === siId ) {
                return si;
            }
        }
        return null;
    }
    
    private isChanged(a:ServerInfoMap, b:ServerInfoMap):boolean {
        const bServerInfos:ServerInfo[] = [];
        for (const key in b){
            if (b.hasOwnProperty(key)) {
                bServerInfos.push(b[key])
            }
        }

        const aServerInfos:ServerInfo[] = [];
        for (const key in a){
            if (a.hasOwnProperty(key)) {
                aServerInfos.push(a[key])
            }
        }

        if ( aServerInfos.length === bServerInfos.length ) {
            // tslint:disable-next-line:prefer-for-of
            for (let i=0;i<aServerInfos.length;i++ ) {
                const found = this.findServerInfo(bServerInfos, aServerInfos[i].id);
                if ( !found ) {
                    return true;
                }
            }
        }
        return false;
    }
}

export default SystemStausView;
