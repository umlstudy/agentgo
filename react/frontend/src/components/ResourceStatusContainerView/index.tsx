import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import ResourceStatusContainerView from './presenter';

const mapStateToProps = (globalState: GlobalState, props:any) => {
    const newSi = globalState.reducer.serverInfoMap[props.serverInfo.id]
    return {
        serverInfo:newSi
    };
};

export default connect(mapStateToProps)(ResourceStatusContainerView);
