import { router } from "../../router";
import { logout } from "../../services/api";
import { ToastView } from "../toast";

export const LogoutView = async () => {
    try {
        const { error } = await logout();
        if (error) {
            console.log(error);
            ToastView("error", error.title!, error.detail!);
            return;
        }

        router.navigate("/");
    } catch (error) {
        console.log(error);
        ToastView("error", error as string);
    }
};
