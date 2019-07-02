export default class ObjectUtil {

    public static isEquivalent(a:any, b:any):boolean {
        const aProps = Object.getOwnPropertyNames(a);
        const bProps = Object.getOwnPropertyNames(b);
    
        if (aProps.length !== bProps.length) {
            return false;
        }
    
        // tslint:disable-next-line:prefer-for-of
        for (let i = 0; i < aProps.length; i++) {
            const propName = aProps[i];
            if (a[propName] !== b[propName]) {
                return false;
            }
        }

        return true;
    }

    public static copyProperties(src:any, target:any):void {
        const aProps = Object.getOwnPropertyNames(src);
    
        // tslint:disable-next-line:prefer-for-of
        for (let i = 0; i < aProps.length; i++) {
            const propName = aProps[i];
            target[propName] = src[propName];
        }
    }
}