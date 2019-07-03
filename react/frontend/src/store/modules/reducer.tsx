import { createAction, handleActions } from 'redux-actions';
import { ProcessStatus } from 'src/model/ProcessStatus';
import { ResourceStatus, VALUES_CNT } from 'src/model/ResourceStatus';
import { ServerInfo } from 'src/model/ServerInfo';
import { StoreObject } from 'src/model/StoreObject';
import ArrayUtil from '../../common/util/ArrayUtil'
import ModelUtil from '../../common/util/ModelUtil'
import ObjectUtil from '../../common/util/ObjectUtil'

// 1.Actions
const TICK = 'counter/TICK';
const REQUEST = 'counter/REQUEST';


// 2.Action Creators
const tick = createAction(TICK);
const request = createAction(REQUEST, (si:ServerInfo)=>si);

// 3.Reducer
// 3.1.  Ininial State
const initialState:StoreObject = {
    num: 92,
// tslint:disable-next-line: object-literal-sort-keys
    tick : 0,
    serverInfoMap:{
        "aaaa":{
            id:"aaaa",
            name:"aaaa",
            resourceStatusesModified:true,
            processStatusesModified:true,
            resourceStatuses: [
                { max:100, min:1, name:"cpu", value:41, values:ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1)} as ResourceStatus,
                { max:100, min:1, name:"Memory", value:41, values:ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1)} as ResourceStatus,
                { max:100, min:1, name:"Disk1", value:41, values:ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1)} as ResourceStatus,
                { max:100, min:1, name:"Disk2", value:41, values:ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1)} as ResourceStatus,
                { max:100, min:1, name:"Disk3", value:41, values:ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1)} as ResourceStatus,
            ],
            processStatuses: [
                { id:'acc', name:'sdf', realName:'dsaf', procId:100 } as ProcessStatus,
            ]
        },
    },
    serverInfoMapModified:true
};

const reducer= handleActions({
    [TICK]:applyTick,
    [REQUEST]:applyRequest
}, initialState);
 
// 4.Reducer Functions
function applyTick(state:StoreObject, action:any) {
    const newState = {
        ...state, 
        tick : state.tick + 1
    };
    Object.keys(newState.serverInfoMap).map((key) => {
        const si = newState.serverInfoMap[key];
        si.resourceStatuses.map((rs:ResourceStatus)=>{
            if ( (rs.values as any).length === 0 ) {
                (rs.values as any).push(0);
            }
            if( key === 'aaaa') {
                (rs.values as any).push(Math.floor(Math.random()*1000)%20+60);
            } else {
                (rs.values as any).push(rs.value);
            }
        });
    })

    return newState;
}

function copyOldStoreObjectAndApplyNew(oldStoreObject:StoreObject, newSi:ServerInfo):StoreObject {
    const oldServerInfoMap = oldStoreObject.serverInfoMap;
    const oldSi = oldServerInfoMap[newSi.id];

    let serverInfoMapModifiedTmp = false;
    if ( !!oldSi ) {
        // 이전에 존재하던 ServerInfo
        {
            // ResourceStatus
            const newSiRss = newSi.resourceStatuses;
            const oldRss = oldSi.resourceStatuses;
            let resourceStatusesModifiedTmp = false;
            newSiRss.forEach((newSiRs:ResourceStatus)=>{
                const foundOldRs = ModelUtil.findById(oldRss, newSiRs.id) as ResourceStatus;
                if ( !!foundOldRs ) {
                    // 이전에 존재하던 ResourceStatus
                    if ( foundOldRs.value !== newSiRs.value) {
                        foundOldRs.value = newSiRs.value;
                    }
                } else {
                    // 이전에 없던 ResourceStatus
                    newSiRs.values = ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1);
                    oldRss.push(newSiRs);
                    resourceStatusesModifiedTmp = true;
                }
            });
            const newOldRss = ModelUtil.removeNotExistIn(oldRss, newSiRss);
            oldSi.resourceStatuses = newOldRss;
            if ( newSiRss.length !== newOldRss.length ) {
                resourceStatusesModifiedTmp = true;
            }
            oldSi.resourceStatusesModified = resourceStatusesModifiedTmp;
        }
        {
            // ProcessStatus
            const newSiPss = newSi.processStatuses;
            const oldPss = oldSi.processStatuses;

            let processStatusesModifiedTmp = false;
            newSiPss.forEach((newSiPs:ProcessStatus)=>{
                const foundOldPs = ModelUtil.findById(oldPss, newSiPs.id) as ProcessStatus;
                if ( !!foundOldPs ) {
                    // 이전에 존재하던 ResourceStatus
                    if ( !ObjectUtil.isEquivalent(foundOldPs, newSiPs) ) {
                        ObjectUtil.copyProperties(newSiPs, foundOldPs);
                        processStatusesModifiedTmp = true;
                    }
                } else {
                    // 이전에 없던 ResourceStatus
                    oldPss.push(newSiPs);
                    processStatusesModifiedTmp = true;
                }
            });
            const newOldPss = ModelUtil.removeNotExistIn(oldPss, newSiPss);
            oldSi.processStatuses = newOldPss;
            if ( newSiPss.length !== newOldPss.length ) {
                processStatusesModifiedTmp = true;
            }
            oldSi.processStatusesModified = processStatusesModifiedTmp;
        }

    } else {
        // 이전에 없던 ServerInfo
        serverInfoMapModifiedTmp = true;
        oldServerInfoMap[newSi.id] = newSi;
        const newRss = newSi.resourceStatuses;
        newRss.forEach((rs:ResourceStatus)=>{
            rs.values = ArrayUtil.getArrayWithLimitedLength(VALUES_CNT+1);
        });
    }

    serverInfoMapModifiedTmp = serverInfoMapModifiedTmp || oldSi.processStatusesModified || oldSi.resourceStatusesModified;

    return {
        ...oldStoreObject, 
        num : oldStoreObject.num - 1,
        serverInfoMap:oldServerInfoMap,
        serverInfoMapModified: serverInfoMapModifiedTmp
    }
}

function applyRequest(state:StoreObject, action:any) {
    const copiedNewStoreObject = copyOldStoreObjectAndApplyNew(state, action.payload);
    return copiedNewStoreObject;
}

// Export Action Creators
export const actionCreators = {
    tick,
    request
};

// Export Reducer
export default reducer;
