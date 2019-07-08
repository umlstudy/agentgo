import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { ServerInfo } from 'src/model/ServerInfo';
import { actionCreators as actions } from '../../store/modules/reducer'
import SystemStausView from './presenter'

const mapStateToProps = (globalState: GlobalState) => {
    return {
        serverInfoMap: globalState.reducer.serverInfoMap,
        serverInfoMapModified: globalState.reducer.serverInfoMapModified
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        tick: () => { dispatch(actions.tick()) },
        request: (si:ServerInfo) => { dispatch(actions.request(si)) },
    };
};

export default connect(mapStateToProps, mapDispatchProps)(SystemStausView);
