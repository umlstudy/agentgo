export const VALUES_CNT:number = 50;

// tslint:disable-next-line:interface-name
export interface ResourceStatus {
    id: string;
    min: number;
    max: number;
    name: string;
    value: number;
    values: number[];
}

// const d = { min:1, max:1, name:"1", value:1} as ResourceStatus;
