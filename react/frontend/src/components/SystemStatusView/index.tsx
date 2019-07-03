import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { ServerInfo } from 'src/model/ServerInfo';
import { actionCreators as counter } from '../../store/modules/counter'
import SystemStausView from './presenter'

const mapStateToProps = (globalState: GlobalState) => {
    return {
        isRunning: globalState.reducer.isRunning,
        serverInfoMap: globalState.reducer.serverInfoMap,
        serverInfoMapModified: globalState.reducer.serverInfoMapModified
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        tick: () => { dispatch(counter.tick()) },
        request: (si:ServerInfo) => { dispatch(counter.request(si)) },
    };
};


export default connect(mapStateToProps, mapDispatchProps)(SystemStausView);
