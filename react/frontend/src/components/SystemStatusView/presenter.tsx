import Axios from 'axios';
import * as React from 'react';
import CheckBox from 'src/common/ui/components/CheckBox/CheckBox';
import ArrayUtil from 'src/common/util/ArrayUtil';
import * as Constants from 'src/Constants';
import { ServerInfo, ServerInfoMap } from 'src/model/ServerInfo';
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
    simpleMode:boolean;
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
                    const si:ServerInfo=response.data[0];
                    nextProps.request(si);
                }
            });
    }

    public state:any = {
        timerInterval:null
    }

    public shouldComponentUpdate(nextProps: SystemStausViewLocalProps, nextStates: SystemStausViewLocalStates):boolean {
        if ( nextStates.simpleMode !== this.state.simpleMode ) {
            return true;
        }
        return nextProps.serverInfoMapModified;
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

    public checkBoxClick(checkbox:boolean) {
        this.setState({
            ...this.state,
            simpleMode: checkbox
        });
    }

    private renderServerViews = (serverInfos: ServerInfo[]) => {
        const checkBoxClick = this.checkBoxClick.bind(this);
        return (
            <>
                <div>
                    <CheckBox msg="간단히" checkBoxClick={checkBoxClick} checked={this.state.simpleMode}/>
                </div>
                <div>
                    {serverInfos.map((value: ServerInfo, index: number, array: ServerInfo[]) => (
                        <ServerView serverInfo={value} key={index} simpleMode={this.state.simpleMode}/>
                    ))}
                </div>
            </>
        );
    }
}

export default SystemStausView;
