import { connect } from 'react-redux';
import { actionCreators as counter } from '../../store/modules/counter'
import SystemStausView from './presenter'


const mapStateToProps = (state: any) => {
    return {
        isRunning: state.counter.isRunning,
        serverInfoMap: state.counter.serverInfoMap
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        tick: () => { dispatch(counter.tick()) },
        request: () => { dispatch(counter.request()) },
    };
};


export default connect(mapStateToProps, mapDispatchProps)(SystemStausView);
