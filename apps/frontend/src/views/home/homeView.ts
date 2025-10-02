import "../../components/home/home.css";
import { PostFeed } from "../../components/home";

export const HomeView = async (element: HTMLElement): Promise<void> => {
    element.append(await PostFeed());
};
