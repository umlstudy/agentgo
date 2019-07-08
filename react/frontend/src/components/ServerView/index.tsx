import { connect } from 'react-redux';
import GlobalState from 'src/model/GlobalState';
import ServerView from './presenter';

const mapStateToProps = (globalState: GlobalState, props:any) => {
    const sid = props.serverInfo.id;
    const isRunning = globalState.reducer.serverInfoMap[sid].isRunning;
    return {
        isRunning
    };
};

export default connect(mapStateToProps)(ServerView);