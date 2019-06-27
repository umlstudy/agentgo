
export interface ISjRect {
    x: number;
    y: number;
    w: number;
    h: number;
}
export default class ContextUtil {
    public static drawRect(ctx:any, rect :ISjRect) {
        ctx.fillRect(rect.x, rect.y, rect.w, rect.h);
    }

    public static drawCenterText(ctx:any, textHeight:number, text:string) {
        const w = ctx.canvas.width;
        const h = ctx.canvas.height;
        const textWidth = Math.floor(ctx.measureText(text ).width);
        ctx.fillText(text , Math.floor((w/2) - (textWidth / 2)), Math.floor((h/2) - (textHeight / 2)));
    }

    public static measureFontHeight(fontStyle:string, text:string):any {
    
        const canvas = document.createElement("canvas");
        const context:any = canvas.getContext("2d");

        const sourceWidth = canvas.width;
        const sourceHeight = canvas.height;
    
        context.font = fontStyle;
        
        // place the text somewhere
        context.textAlign = "left";
        context.textBaseline = "top";
        context.fillText(text, 25, 5);
    
        // returns an array containing the sum of all pixels in a canvas
        // * 4 (red, green, blue, alpha)
        // [pixel1Red, pixel1Green, pixel1Blue, pixel1Alpha, pixel2Red ...]
        const data = context.getImageData(0, 0, sourceWidth, sourceHeight).data;
    
        let firstY = -1;
        let lastY = -1;
    
        // loop through each row
        for(let y = 0; y < sourceHeight; y++) {
            // loop through each column
            for(let x = 0; x < sourceWidth; x++) {
                // var red = data[((sourceWidth * y) + x) * 4];
                // var green = data[((sourceWidth * y) + x) * 4 + 1];
                // var blue = data[((sourceWidth * y) + x) * 4 + 2];
                const alpha = data[((sourceWidth * y) + x) * 4 + 3];
    
                if(alpha > 0) {
                    firstY = y;
                    // exit the loop
                    break;
                }
            }
            if(firstY >= 0) {
                // exit the loop
                break;
            }
        }
    
        // loop through each row, this time beginning from the last row
        for(let y = sourceHeight; y > 0; y--) {
            // loop through each column
            for(let x = 0; x < sourceWidth; x++) {
                const alpha = data[((sourceWidth * y) + x) * 4 + 3];
                if(alpha > 0) {
                    lastY = y;
                    // exit the loop
                    break;
                }
            }
            if(lastY >= 0) {
                // exit the loop
                break;
            }
        }
    
        return {
            // The first pixel
            firstPixel: firstY,
            // The actual height
            height: lastY - firstY,
            // The last pixel
            lastPixel: lastY
        }
    };
}