import { connect } from 'react-redux';
import { ServerInfo } from 'src/model/ServerInfo';
import { actionCreators as counter } from '../../store/modules/counter'
import SystemStausView from './presenter'


const mapStateToProps = (state: any) => {
    return {
        isRunning: state.counter.isRunning,
        serverInfoMap: state.counter.serverInfoMap,
        serverInfoMapModified: state.counter.serverInfoMapModified
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        tick: () => { dispatch(counter.tick()) },
        request: (si:ServerInfo) => { dispatch(counter.request(si)) },
    };
};


export default connect(mapStateToProps, mapDispatchProps)(SystemStausView);
