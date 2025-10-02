import { Banner } from "../../components/banner";
import { Card } from "../../components/card";
import { api } from "../../services/mockApi";
import { renderLoading } from "../common";

export const PostsView = async (element: HTMLElement) => {
    const intervalId = renderLoading();
    const posts = await api.fetchPosts();
    const postLinks = posts.map(
        (post) => `
        <a href="/posts/${post.id}">${post.title}</a>
    `,
    );
    clearInterval(intervalId);
    Banner("Articles");
    element.innerHTML = "";
    element.append(
        Card({
            title: "Blogs",
            description: "",
            list: postLinks,
        }),
    );
};
