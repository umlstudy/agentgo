
export default class ArrayUtil {
    public static getArrayWithLimitedLength<T>(length:number):T[] {
        const arr = new Array<T>();
        arr.push = function () {
            if (this.length >= length) {
                this.shift();
            }
            return Array.prototype.push.apply(this,arguments);
        }
        return arr;
    }
}