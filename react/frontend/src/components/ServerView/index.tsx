import { connect } from 'react-redux';
import { actionCreators as counter } from '../../store/modules/counter'
import ServerView from './presenter'

const mapStateToProps = (globalState: any) => {
    return {
        num: globalState.counter.num,
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        handleDecrement: () => { dispatch(counter.decrement()) },
        handleIncrement: () => { dispatch(counter.increment()) },
    };
};

export default connect(mapStateToProps, mapDispatchProps)(ServerView);
