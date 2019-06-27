import { connect } from 'react-redux';
import { actionCreators as counter } from '../../store/modules/counter'
import SystemStausView from './presenter'


const mapStateToProps = (state: any) => {
    return {
        isRunning: state.counter.isRunning,
        serverInfos: state.counter.serverInfos
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        tick: () => { dispatch(counter.tick()) },
    };
};


export default connect(mapStateToProps, mapDispatchProps)(SystemStausView);
