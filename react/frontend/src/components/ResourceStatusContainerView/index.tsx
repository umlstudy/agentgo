import { connect } from 'react-redux';
import ResourceStatusContainerView from './presenter';

const mapStateToProps = (state: any) => {
    return {
        num: state.counter.num,
        resourceStatusesModified: state.counter.resourceStatusesModified,
    };
};

export default connect(mapStateToProps)(ResourceStatusContainerView);
