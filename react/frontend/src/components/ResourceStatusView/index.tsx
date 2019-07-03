
import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import ResourceStatusView from './presenter'

const mapStateToProps = (globalState: GlobalState) => {
    return {
        tick: globalState.reducer.tick,
    };
};

export default connect(mapStateToProps)(ResourceStatusView);
