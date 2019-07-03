import { AbstractStatus } from './AbstractStatus';

// tslint:disable-next-line: interface-name
export interface ProcessStatus extends AbstractStatus {
    realName:string;
    procId:number;
}