import { connect } from 'react-redux';
import { StoreObject } from 'src/model/StoreObject';
import { ServerViewProps } from '../ServerView/presenter';
import ResourceStatusContainerView from './presenter';

const mapStateToProps = (globalState: any, ownerProps:ServerViewProps) => {
    const newSi = (globalState.counter as StoreObject).serverInfoMap[ownerProps.serverInfo.id]
    return {
        num: globalState.counter.num,
        serverInfo:newSi
    };
};

export default connect(mapStateToProps)(ResourceStatusContainerView);
