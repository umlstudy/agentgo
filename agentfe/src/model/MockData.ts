import { ResourceStatus } from './ResourceStatus';
import { ServerInfo } from "./ServerInfo";

export const serverInfos2:ServerInfo[] = [
    {
        name:"aaaa",
        resourceStatuses: [
            { max:100, min:1, name:"cpu", value:41} as ResourceStatus,
            { max:100, min:1, name:"Memory", value:41} as ResourceStatus,
            { max:100, min:1, name:"Disk1", value:41} as ResourceStatus,
            { max:100, min:1, name:"Disk2", value:41} as ResourceStatus,
            { max:100, min:1, name:"Disk3", value:41} as ResourceStatus,
        ]
    },
];