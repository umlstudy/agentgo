import { createAction, handleActions } from 'redux-actions';
import { ResourceStatus } from 'src/model/ResourceStatus';
import { ServerInfo } from 'src/model/ServerInfo';
import ArrayUtil from '../../common/util/ArrayUtil'

// 1.Actions
const INCREMENT = 'counter/INCREMENT';
const DECREMENT = 'counter/DECREMENT';
const TICK = 'counter/TICK';

// 2.Action Creators
const increment = createAction(INCREMENT);
const decrement = createAction(DECREMENT);
const tick = createAction(TICK);

// 3.Reducer
// 3.1.  Ininial State
const initialState = {
    num: 92,
// tslint:disable-next-line: object-literal-sort-keys
    isRunning:false,
    tick : 0,
    serverInfos:[
        {
            name:"aaaa",
            resourceStatuses: [
                { max:100, min:1, name:"cpu", value:41, values:ArrayUtil.getArrayWithLimitedLength(50)} as ResourceStatus,
                { max:100, min:1, name:"Memory", value:41, values:ArrayUtil.getArrayWithLimitedLength(50)} as ResourceStatus,
                { max:100, min:1, name:"Disk1", value:41, values:ArrayUtil.getArrayWithLimitedLength(50)} as ResourceStatus,
                { max:100, min:1, name:"Disk2", value:41, values:ArrayUtil.getArrayWithLimitedLength(50)} as ResourceStatus,
                { max:100, min:1, name:"Disk3", value:41, values:ArrayUtil.getArrayWithLimitedLength(50)} as ResourceStatus,
            ]
        },
    ]
};

const reducer= handleActions({
    [INCREMENT]:applyIncrement,
    [DECREMENT]:applyDecrement,
    [TICK]:applyTick
}, initialState);
 
// 4.Reducer Functions
function applyIncrement(state:any, action:any) {
    return {
        ...state, 
        num : state.num + 1
    }
}

function applyDecrement(state:any, action:any) {
    return {
        ...state, 
        num : state.num - 1
    }
}

function applyTick(state:any, action:any) {
    const newState = {
        ...state, 
        tick : state.tick + 1
    };
    newState.serverInfos.map((si:ServerInfo)=>{
        si.resourceStatuses.map((rs:ResourceStatus)=>{
            if ( (rs.values as any).length === 0 ) {
                (rs.values as any).push(0);
            }
            (rs.values as any).push(Math.floor(Math.random()*1000)%20+60);
        });
    });
    return newState;
}

// Export Action Creators
export const actionCreators = {
    decrement,
    increment,
    tick
};

// Export Reducer
export default reducer;

// extra constants
export const hello = 'aaaaaaaaaaaaaabbbbbb';


// // handleActions 의 첫번째 파라미터는 액션을 처리하는 함수들로 이뤄진 객체이고
// // 두번째 파라미터는 초기 상태입니다.
// const reducer = handleActions({
//     [INCREMENT]: (state: any, action: any) => {
//         return { 
//             num: state.num + 1 
//         };
//     },
//     [TICK]: (state: any, action: any) => {
//         return {
//             tick: state.tick + 1 
//         };
//     },
    
//     // action 객체를 참조하지 않으니까 이렇게 생략을 할 수도 있겠죠?
//     // state 부분에서 비구조화 할당도 해주어서 코드를 더욱 간소화시켰습니다.
//     [DECREMENT]: (state: any, action: any) => {
//         return {
//             num: state.num - 1
//         };
//     }
// }, initialState);

// 외부 참조용 액션 생성 함수를 만듭니다.
// reducer
// export default reducer;

