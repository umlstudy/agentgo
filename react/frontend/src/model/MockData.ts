import { ProcessStatus } from './ProcessStatus';
import { ResourceStatus } from './ResourceStatus';
import { ServerInfoMap } from "./ServerInfo";

export const serverInfoMap:ServerInfoMap = {
    "aaaa":{
        id:"aaaa",
        name:"aaaa",
        resourceStatusesModified:true,
        processStatusesModified:true,
        resourceStatuses: [
            { max:100, min:1, name:"cpu", value:41} as ResourceStatus,
            { max:100, min:1, name:"Memory", value:41} as ResourceStatus,
            { max:100, min:1, name:"Disk1", value:41} as ResourceStatus,
            { max:100, min:1, name:"Disk2", value:41} as ResourceStatus,
            { max:100, min:1, name:"Disk3", value:41} as ResourceStatus,
        ],
        processStatuses: [
            { id:'acc', name:'sdf', realName:'dsaf', procId:100 } as ProcessStatus,
        ],
        isRunning:true
    },
};