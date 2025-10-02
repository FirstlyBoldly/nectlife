import "./home.css";
import { api } from "../../services/mockApi";
import { Banner } from "../banner";
import { ToastView, renderLoading } from "../../views";
import { Thumbnail } from "../thumbnail";

export const PostFeed = async (): Promise<HTMLElement> => {
    const intervalId = renderLoading();

    const feed = document.createElement("div");
    feed.classList.add("feed");

    try {
        const posts = await api.fetchPosts();
        const postsList = document.createElement("ul");

        posts.forEach((post) => {
            const li = document.createElement("li");
            li.classList.add("feed-list");
            li.append(Thumbnail(post));
            postsList.append(li);
        });

        clearInterval(intervalId);
        Banner("Latest Posts");
        feed.append(postsList);
    } catch (e) {
        ToastView("error", (e as Error).message);
    }

    return feed;
};
