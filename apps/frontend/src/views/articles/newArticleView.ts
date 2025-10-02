import { NewPostForm } from "../../components/forms";
import { Banner } from "../../components/banner";

export const NewPostView = (element: HTMLElement) => {
    Banner("New Post");
    const form = NewPostForm();
    element.append(form);
};
