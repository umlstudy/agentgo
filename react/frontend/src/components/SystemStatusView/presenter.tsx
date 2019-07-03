import Axios from 'axios';
import * as React from 'react';
import ArrayUtil from 'src/common/util/ArrayUtil';
import { ServerInfo, ServerInfoMap } from 'src/model/ServerInfo';
import * as Constants from '../Constants';
import ServerView from '../ServerView';

// tslint:disable-next-line:interface-name
interface SystemStausViewLocalProps {
    serverInfoMap: ServerInfoMap;
    serverInfoMapModified: boolean;
    tick:() => void;
    request:(si:ServerInfo) => void;
};

// tslint:disable-next-line:interface-name
interface SystemStausViewLocalStates {
    serverInfoMapModified: boolean;
    timerInterval:NodeJS.Timeout;
};
class SystemStausView extends React.Component<SystemStausViewLocalProps, SystemStausViewLocalStates> {

    public static getDerivedStateFromProps(nextProps: SystemStausViewLocalProps, prevState: SystemStausViewLocalStates):SystemStausViewLocalStates {
        if ( !prevState.timerInterval ) {
            const timerInterval = setInterval(() => {
                nextProps.tick();
                SystemStausView.requestDateFromServer(nextProps);
            }, 1000);
            return {
                serverInfoMapModified:nextProps.serverInfoMapModified,
                timerInterval
            }
        }

        return prevState;
    }

    private static requestDateFromServer = (nextProps: SystemStausViewLocalProps):void => {
        Axios.get(Constants.GATEWAY_URL)
            .then((response)=>{
                const si:ServerInfo=response.data[0];
                nextProps.request(si);
            });
    }

    public state:any = {
        timerInterval:null
    }

    public shouldComponentUpdate(nextProps: SystemStausViewLocalProps, nextState: SystemStausViewLocalStates):boolean {
        return nextState.serverInfoMapModified;
    }

    public render() {
        const serverInfoMap = this.props.serverInfoMap;
        const serverInfos:ServerInfo[] = ArrayUtil.json2Array(serverInfoMap);
        return (
            <div>
                {this.renderServerViews(serverInfos)}
            </div>
        );
    }

    private renderServerViews = (serverInfos: ServerInfo[]) => {
        return serverInfos.map((value: ServerInfo, index: number, array: ServerInfo[]) => (
            <ServerView serverInfo={value} key={index}/>
        ));
    }
}

export default SystemStausView;
