import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { ServerInfo } from 'src/model/ServerInfo';
import ProcessStatusContainerView from './presenter';

const mapStateToProps = (globalState: GlobalState, props:any) => {
    const newSi = globalState.reducer.serverInfoMap[props.serverInfo.id] as ServerInfo
    if ( newSi.processStatusesModified ) {
        return {
            serverInfo:{
                ...newSi
            }
        };
    }

    return null;
};

export default connect(mapStateToProps)(ProcessStatusContainerView);
