import * as React from 'react';
import CheckBox from 'src/common/ui/components/CheckBox/CheckBox';
import { ServerInfo } from 'src/model/ServerInfo';
import ProcessStatusContainerView from '../ProcessStatusContainerView';
import ResourceStatusContainerView from '../ResourceStatusContainerView';
import './ServerView.css';

// tslint:disable-next-line:interface-name
export interface ServerViewProps {
    serverInfo:ServerInfo;
    simpleMode:boolean;
    isRunning:boolean;
}

// tslint:disable-next-line: interface-name
interface ServerViewLocalStates {
    simpleMode:boolean;
    fireEventFromServerView:boolean;
    isRunning:boolean;
};
class ServerView extends React.Component<ServerViewProps, ServerViewLocalStates> {

    public static getDerivedStateFromProps(nextProps: ServerViewProps, prevState: ServerViewLocalStates):ServerViewLocalStates {
        console.log("ServerView - getDerivedStateFromProps called");
        if ( prevState.fireEventFromServerView ) {
            return {
                ...prevState,
                isRunning:nextProps.isRunning,
                fireEventFromServerView:false
            }
        } else {
            return {
                ...prevState,
                simpleMode:nextProps.simpleMode,
                isRunning:nextProps.isRunning,
                fireEventFromServerView:false
            }
        }
    }

    public state:ServerViewLocalStates = {
        simpleMode:true,
        isRunning:true,
        fireEventFromServerView:false
    }

    public shouldComponentUpdate(nextProps: ServerViewProps, nextStates: ServerViewLocalStates):boolean {
        // if ( this.state.simpleMode !== nextStates.simpleMode ) {
        //     console.log("ServerView - shouldComponentUpdate true");
        //     return true;
        // } else {
        //     const sid = nextProps.serverInfo.id;
        //     const currIsRunning = this.props.serverInfoMap[sid];
        //     const nextIsRunning = nextProps.serverInfoMap[sid];
        //     if ( currIsRunning !== nextIsRunning ) {
        //         console.log("ServerView - shouldComponentUpdate true");
        //         return true;
        //     }
        // }
        console.log("ServerView - shouldComponentUpdate false");
        return true;
    }

    public render() {
        const checkBoxClick = this.checkBoxClick.bind(this);
        const simpleMode = this.state.simpleMode;
        // const simpleMode = this.props.simpleMode;

        console.log("ServerView - render simpleMode " + simpleMode);

        return (
            <div className="ServerView">
                <div className={ this.props.serverInfo.isRunning ? "ServerNamePart" :"ServerNamePartNotRunning" }>
                    <span className="ServerName">
                        {this.props.serverInfo.name}
                    </span><CheckBox msg="간단히" checkBoxClick={checkBoxClick} checked={simpleMode}/>
                </div>
                <div>
                    <ResourceStatusContainerView serverInfo={this.props.serverInfo} simpleMode={simpleMode}/>
                    { !!this.props.serverInfo.processStatuses ? <ProcessStatusContainerView serverInfo={this.props.serverInfo} simpleMode={simpleMode}/>: ''}
                </div>
            </div>
        );
    }
    
    private checkBoxClick(checkbox:boolean) {
        this.setState({
            ...this.state,
            simpleMode: checkbox,
            fireEventFromServerView:true
        });
    }
};

export default ServerView;