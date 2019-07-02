import { ProcessStatus } from './ProcessStatus';
import { ResourceStatus } from './ResourceStatus';

// tslint:disable-next-line:interface-name
export interface ServerInfo {
    id : string;
    name: string;
    resourceStatusesModified:boolean;
    processStatusesModified:boolean;
    resourceStatuses: ResourceStatus[];
    processStatuses: ProcessStatus[];
}

export type ServerInfoMap = Record<string, ServerInfo>;
