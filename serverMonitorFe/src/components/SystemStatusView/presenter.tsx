import * as React from 'react';
import { ServerInfo } from 'src/model/ServerInfo';
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
        return false;
    }

    public componentDidUpdate(prevProps:any, prevState:any) {
        return true;
    }

    public render() {
        const props = (this as any).props;
        const serverInfos = props.serverInfos;
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
