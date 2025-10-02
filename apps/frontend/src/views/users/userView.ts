import { Banner } from "../../components/banner";
import { UserProfile } from "../../components/user/user_profile";
import type { RouteParams } from "../../router";
import { getUserById } from "../../services/api";
import { renderLoading } from "../common";
import { ToastView } from "../toast";

export const UserView = async (target: HTMLElement, params: RouteParams) => {
    const intervalId = renderLoading();
    const id = parseInt(params.id, 10);

    try {
        const { data, error } = await getUserById({ path: { id: id } });
        if (error) {
            console.log(error);
            ToastView("error", error.title!, error.detail!);
            return;
        }

        target.innerHTML = "";

        clearInterval(intervalId);
        Banner("Profile");

        target.append(UserProfile(data));
    } catch(error) {
        console.error(error);
        return;
    }
};
