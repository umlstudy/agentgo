import { WarningLevel } from './enums/WarningLevel';

// tslint:disable-next-line:interface-name
export interface AbstractStatus {
    id:string;
    name:string;
    wl?:WarningLevel;
}