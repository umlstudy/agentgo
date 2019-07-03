import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { actionCreators as counter } from '../../store/modules/counter'
import ServerView from './presenter'

const mapStateToProps = (globalState: GlobalState) => {
    return {
        num: globalState.reducer.num,
    };
};

const mapDispatchProps = (dispatch: any) => {
    return {
        handleDecrement: () => { dispatch(counter.decrement()) },
        handleIncrement: () => { dispatch(counter.increment()) },
    };
};

export default connect(mapStateToProps, mapDispatchProps)(ServerView);
