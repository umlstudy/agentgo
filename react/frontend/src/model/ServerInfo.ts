import { AbstractStatus } from './AbstractStatus';
import { ProcessStatus } from './ProcessStatus';
import { ResourceStatus } from './ResourceStatus';

// tslint:disable-next-line:interface-name
export interface ServerInfo extends AbstractStatus {
    resourceStatusesModified:boolean;
    processStatusesModified:boolean;
    resourceStatuses: ResourceStatus[];
    processStatuses: ProcessStatus[];
}

export type ServerInfoMap = Record<string, ServerInfo>;
