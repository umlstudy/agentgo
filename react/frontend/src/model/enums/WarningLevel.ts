
export enum WarningLevel {
    NORMAL = 1,
    WARNING ,
    ERROR 
}

export class WarningLevelUtil {
    public static getLabel(wl:WarningLevel ):string {
        switch (wl ) {
            case WarningLevel.NORMAL : {
                return "정보";
            }
            case WarningLevel.WARNING : {
                return "경고";
            }
            case WarningLevel.ERROR : {
                return "에러";
            }
        }
    }

    public static getCode(wl:WarningLevel ):string {
        switch (wl ) {
            case WarningLevel.NORMAL : {
                return "10";
            }
            case WarningLevel.WARNING : {
                return "20";
            }
            case WarningLevel.ERROR : {
                return "30";
            }
        }
    }
}