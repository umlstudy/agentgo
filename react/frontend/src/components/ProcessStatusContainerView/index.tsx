import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { ServerViewProps } from '../ServerView/presenter';
import ProcessStatusContainerView from './presenter';

const mapStateToProps = (globalState: GlobalState, ownProps:ServerViewProps) => {
    const newSi = globalState.reducer.serverInfoMap[ownProps.serverInfo.id]
    return {
        serverInfo:newSi
    };
};

export default connect(mapStateToProps)(ProcessStatusContainerView);
