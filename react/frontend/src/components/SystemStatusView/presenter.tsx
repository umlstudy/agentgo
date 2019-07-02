import Axios from 'axios';
import * as React from 'react';
import ArrayUtil from 'src/common/util/ArrayUtil';
import { ServerInfo } from 'src/model/ServerInfo';
import { StoreObject } from 'src/model/StoreObject';
import * as Constants from '../Constants';
import ServerView from '../ServerView/presenter';

class SystemStausView extends React.Component<any, any> {

    public static getDerivedStateFromProps(nextProps: any, prevState: any) {
        if ( prevState == null ) {
            const timerInterval = setInterval(() => {
                nextProps.tick();
                SystemStausView.requestDateFromServer(nextProps);
            }, 1000);
            return {
                ...nextProps,
                isRunning:true,
                timerInterval
            }
        }

        return null;
    }

    private static requestDateFromServer = (nextProps: any) => {
        Axios.get(Constants.GATEWAY_URL)
        .then((response)=>{
            const si:ServerInfo=response.data[0];
            nextProps.request(si);
        });
    }

    // public state:any = {
    //     isRunning:null
    // }

    // public getSnapshotBeforeUpdate(prevProps:any, prevState:any) {
    //     if (this.props.isRunning !== prevProps.isRunning ) {
    //         return true;
    //     } 
    //     return false;
    // }

    // public componentDidUpdate(prevProps:any, prevState:any) {
    //     return true;
    // }

    public shouldComponentUpdate(nextProps: any, nextState: any) {
        return nextState.serverInfoMapModified;
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
}

export default SystemStausView;
