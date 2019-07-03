import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { ServerInfo } from 'src/model/ServerInfo';
import { ServerViewProps } from '../ServerView/presenter';
import ProcessStatusContainerView from './presenter';

const mapStateToProps = (globalState: GlobalState, ownProps:ServerViewProps) => {
    const newSi = globalState.reducer.serverInfoMap[ownProps.serverInfo.id] as ServerInfo
    if ( newSi.processStatusesModified ) {
        return {
            serverInfo:{
                ...newSi
            }
        };
    } else {
        return {
            serverInfo:newSi
        };
    }
};

export default connect(mapStateToProps)(ProcessStatusContainerView);
