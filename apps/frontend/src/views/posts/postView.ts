import { router, type RouteParams } from "../../router";
import { Button } from "../../components/button";
import { Card } from "../../components/card";
import { api } from "../../services/mockApi";
import { renderLoading } from "../common";
import { ToastView } from "../toast/toastView";
import { Banner } from "../../components/banner";

export const PostView = async (element: HTMLElement, params: RouteParams) => {
    const intervalId = renderLoading();
    const id = parseInt(params.id, 10);
    const post = await api.fetchPostsById(id);

    if (post) {
        clearInterval(intervalId);
        Banner(post.title);
        element.innerHTML = "";
        element.append(
            Button({
                label: "All Posts",
                onCLick: (): void => {
                    router.navigate("/articles");
                },
            }),
            Card({
                title: post.title,
                description: post.content,
            }),
        );
    } else {
        ToastView("error", `Failed to fetch post with id: ${params.id}.`);
    }
};
