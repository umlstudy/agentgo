import { connect } from 'react-redux';
import * as counter from '../../store/modules/counter'
import ServerView from './presenter'

const mapStateToProps = (state: any) => {
    return {
        num: state.counter.num,
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        handleDecrement: () => { dispatch(counter.decrement()) },
        handleIncrement: () => { dispatch(counter.increment()) },
    };
};

export default connect(mapStateToProps, mapDispatchProps)(ServerView);
