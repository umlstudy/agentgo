import Axios from 'axios';
import * as React from 'react';
import CheckBox from 'src/common/ui/components/CheckBox/CheckBox';
import ArrayUtil from 'src/common/util/ArrayUtil';
import * as Constants from 'src/Constants';
import { ServerInfo, ServerInfoMap } from 'src/model/ServerInfo';
import ServerView from '../ServerView';
import './SystemStatusView.css'

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
    simpleMode:boolean;
};

class SystemStausView extends React.Component<SystemStausViewLocalProps, SystemStausViewLocalStates> {

    public static getDerivedStateFromProps(nextProps: SystemStausViewLocalProps, prevState: SystemStausViewLocalStates):SystemStausViewLocalStates {
        console.log("SystemStausView - getDerivedStateFromProps ");
        if ( !prevState.timerInterval ) {
            const timerInterval = setInterval(() => {
                nextProps.tick();
                SystemStausView.requestDateFromServer(nextProps);
            }, 1000);
            return {
                serverInfoMapModified:nextProps.serverInfoMapModified,
                simpleMode:true,
                timerInterval
            }
        }

        return prevState;
    }

    private static requestDateFromServer = (nextProps: SystemStausViewLocalProps):void => {
        Axios.get(Constants.GATEWAY_URL)
            .then((response)=>{
                if ( response.data.length > 0 ) {
                    // console.log("SystemStausView - requestDateFromServer si size " + response.data.length);
                    nextProps.request(response.data);
                }
            });
    }

    public state:any = {
        timerInterval:null
    }

    public shouldComponentUpdate(nextProps: SystemStausViewLocalProps, nextStates: SystemStausViewLocalStates):boolean {
        if ( nextStates.simpleMode !== this.state.simpleMode ) {
            console.log("SystemStausView - shouldComponentUpdate true");
            return true;
        }
        console.log("SystemStausView - shouldComponentUpdate " + nextProps.serverInfoMapModified);
        return nextProps.serverInfoMapModified;
    }

    public render() {
        const serverInfoMap = this.props.serverInfoMap;
        const serverInfos:ServerInfo[] = ArrayUtil.json2Array(serverInfoMap);

        console.log("SystemStausView - render ");

        return (
            <div>
                {this.renderServerViews(serverInfos)}
            </div>
        );
    }

    public checkBoxClick(checkbox:boolean) {
        this.setState({
            ...this.state,
            simpleMode: checkbox
        });
    }

    private renderServerViews = (serverInfos: ServerInfo[]) => {
        const checkBoxClick = this.checkBoxClick.bind(this);
        serverInfos = serverInfos.sort((a,b)=> b.sortOrder-a.sortOrder);
        
        return (
            <>
                <div className="SystemStatusView-header">
                    <span className="SystemStatusView-title">Server Monitor</span><CheckBox msg="간단히" checkBoxClick={checkBoxClick} checked={this.state.simpleMode}/>
                </div>
                <div>
                    {serverInfos.map((value: ServerInfo/* , index: number, array: ServerInfo[] */) => (
                        <ServerView serverInfo={value} key={value.id} simpleMode={this.state.simpleMode}/>
                    ))}
                </div>
            </>
        );
    }
}

export default SystemStausView;
