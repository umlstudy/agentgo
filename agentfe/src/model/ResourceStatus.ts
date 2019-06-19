
export class ResourceStatus {
    
    private min: number;
    private max: number;
    private name: string;
    private values: number[];

	constructor($min: number, $max: number, $name: string, $values: number[]) {
		this.min = $min;
        this.max = $max;
        this.name = $name;
        this.values = $values;
	}

    /**
     * Getter $name
     * @return {string}
     */
	public get $name(): string {
		return this.name;
	}

    /**
     * Getter $values
     * @return {number[]}
     */
	public get $values(): number[] {
		return this.values;
    }
    
    /**
     * Getter $min
     * @return {number}
     */
	public get $min(): number {
		return this.min;
    }
    
    /**
     * Getter $max
     * @return {number}
     */
	public get $max(): number {
		return this.max;
	}
}