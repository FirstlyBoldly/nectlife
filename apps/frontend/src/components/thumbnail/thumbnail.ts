import "./thumbnail.css";
import { router } from "../../router";
import { Button } from "../button";
import type { PostSummary } from "../../services/mockApi";

export const Thumbnail = (params: PostSummary): HTMLElement => {
    const thumbnail = document.createElement("div");
    thumbnail.classList.add("thumbnail", "basic-container");
    thumbnail.innerHTML = `
        <a class="post-link" href="/articles/${params.id}">
            <h2>${params.title}</h2>
        </a>
        <ul>
            <li>#placeholder</li>
        </ul>
    `;
    thumbnail.append(
        Button({
            label: "Read More",
            onCLick: (): void => {
                router.navigate(`/articles/${params.id}`);
            },
        }),
    );
    return thumbnail;
};
