import { ResourceStatus } from './ResourceStatus';

export class ServerInfo {
    
    private name: string;
    private resourceStatus: ResourceStatus[];

	constructor($name: string, $resourceStatus: ResourceStatus[]) {
        this.name = $name;
		this.resourceStatus = $resourceStatus;
	}

    /**
     * Getter $name
     * @return {string}
     */
	public get $name(): string {
		return this.name;
	}
    
    /**
     * Getter $resourceStatus
     * @return {ResourceStatus[]}
     */
	public get $resourceStatus(): ResourceStatus[] {
		return this.resourceStatus;
	}
}