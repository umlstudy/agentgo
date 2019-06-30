import { ResourceStatus } from './ResourceStatus';

// tslint:disable-next-line:interface-name
export interface ServerInfo {
    id : string;
    name: string;
    resourceStatuses: ResourceStatus[];
}

export type ServerInfoMap = Record<string, ServerInfo>;
