import { CpuStatus } from './CpuStatus';
import { DiskStatus } from './DiskStatus';
import { MemoryStatus } from './MemoryStatus';
import { ResourceStatus } from './ResourceStatus';
import { ServerInfo } from "./ServerInfo";

const rsList:ResourceStatus[] = [
    new CpuStatus(1, 100, "cpu", [1,2,3,4,5]),
    new MemoryStatus(1, 100, "Memory", [1,2,3,4,5]),
    new DiskStatus(1, 100, "Disk1", [1,2,3,4,5]),
    new DiskStatus(1, 100, "Disk2", [1,2,3,4,5]),
];

export const serverInfos:ServerInfo[] = [
    new ServerInfo("aaa", rsList),
];