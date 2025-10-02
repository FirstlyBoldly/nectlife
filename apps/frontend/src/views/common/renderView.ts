import { Banner } from "../../components/banner";

const DOT_MAXIMUM = 3;
let count = 0;

export const renderLoading = (): NodeJS.Timeout => {
    Banner("Loading");
    const loading = (): void => {
        Banner("Loading" + ".".repeat(++count));
        if (count == DOT_MAXIMUM) {
            count = 0;
        }
    };
    return setInterval(loading, 200);
};
