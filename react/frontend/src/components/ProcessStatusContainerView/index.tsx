import { connect } from 'react-redux';
import ProcessStatusContainerView from './presenter';

const mapStateToProps = (state: any) => {
    return {
        num: state.counter.num,
        processStatusesModified : state.counter.processStatusesModified,
    };
};

export default connect(mapStateToProps)(ProcessStatusContainerView);
