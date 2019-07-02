import { ServerInfoMap } from "./ServerInfo";

// tslint:disable-next-line: interface-name
export interface StoreObject {
    num:number;
    isRunning:boolean;
    tick:number;
    serverInfoMapModified:boolean,
    serverInfoMap:ServerInfoMap
}