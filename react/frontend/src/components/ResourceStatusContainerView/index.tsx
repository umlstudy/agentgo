import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import { ServerViewProps } from '../ServerView/presenter';
import ResourceStatusContainerView from './presenter';

const mapStateToProps = (globalState: GlobalState, ownerProps:ServerViewProps) => {
    const newSi = globalState.reducer.serverInfoMap[ownerProps.serverInfo.id]
    return {
        num: globalState.reducer.num,
        serverInfo:newSi
    };
};

export default connect(mapStateToProps)(ResourceStatusContainerView);
