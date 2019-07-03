import { connect } from 'react-redux';
import { StoreObject } from 'src/model/StoreObject';
import ProcessStatusContainerView from './presenter';

const mapStateToProps = (globalState: any, ownProps:any) => {
    const newSi = (globalState.counter as StoreObject).serverInfoMap[ownProps.serverInfo.id]
    return {
        num: globalState.counter.num,
        serverInfo:newSi
    };
};

export default connect(mapStateToProps)(ProcessStatusContainerView);
