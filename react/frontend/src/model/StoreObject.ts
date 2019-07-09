import { ServerInfoMap } from "./ServerInfo";

// tslint:disable-next-line: interface-name
export interface StoreObject {
    tick:number;
    serverInfoMapModified:boolean,
    serverInfoMap:ServerInfoMap
}