
import { connect } from 'react-redux';
import ResourceStatusView from './presenter'

const mapStateToProps = (globalState: any) => {
    return {
        tick: globalState.counter.tick,
    };
};

export default connect(mapStateToProps)(ResourceStatusView);
