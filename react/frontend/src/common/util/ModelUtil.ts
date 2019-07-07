
// tslint:disable-next-line:interface-name
interface Identifier {
    id:string
}

export default class ArrayUtil {

    public static findById<T extends Identifier>(objs:T[], idStr:string):T|null {
        let idx:any;
        for ( idx in objs) {
            if ( objs[idx].id === idStr ) {
                return objs[idx];
            }
        }
        return null;
    }

    public static removeNotExistIn<T extends Identifier>(forRemove: T[], origins: T[]): T[] {
        const result:T[] = [];
        for ( const origin of origins ) {
            const found = this.findById(forRemove, origin.id);
            if ( !!found ) {
                result.push(found);
            }
        }

        return result;
    }
}