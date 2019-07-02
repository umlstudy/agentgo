
import { connect } from 'react-redux';
import ProcessStatusView from './presenter';

const mapStateToProps = (state: any) => {
    return {
        tick: state.counter.tick,
    };
};

export default connect(mapStateToProps)(ProcessStatusView);
