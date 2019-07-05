import { AbstractStatus } from './AbstractStatus';

// tslint:disable-next-line:interface-name
export interface ResourceStatus extends AbstractStatus {
    min: number;
    max: number;
    value: number;
    values: number[];
}

// const d = { min:1, max:1, name:"1", value:1} as ResourceStatus;
