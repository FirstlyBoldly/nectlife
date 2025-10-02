import type { Status } from "../../router";
import { getUserById } from "../../services/api";
import { Icon } from "../icon/icon";
import "./home.css";

export const reloadUserProfile = async (status: Status): Promise<void> => {
    const profileContainer = document.getElementById("profile-container");
    if (profileContainer && status.authenticated) {
        try {
            const { data, error } = await getUserById({path: {id: status.userId!}});
            if (error) {
                console.error(error);
                return;
            }

            profileContainer.innerHTML = `
                <span>Hello ${data.FirstName}!</span>
            `;
        } catch(error) {
            console.error(error);
            return;
        }
    } else {
        profileContainer!.innerHTML = "";
    }
};

export const ProfileSidebar = (): HTMLElement => {
    const sidebar = document.createElement("div");
    sidebar.classList.add("side-profile", "basic-container");
    sidebar.innerHTML = `
        <div id="profile-container"></div>
        <h3>General</h3>
        <ul>
            <li><a href="/"><div id="home-icon-placeholder"></div>Home</a></li>
            <li><a href="/articles">Posts</a></li>
            <li><a href="/about">About</a></li>
            <li><a href="/contact">Contact</a></li>
        </ul>
        <hr class="profile-sidebar-divider">
        <h3>Topics</h3>
        <ul>
            <li><a href=".">Blah</a></li>
            <li><a href=".">Blah</a></li>
            <li><a href=".">Blah</a></li>
            <li><a href=".">Blah</a></li>
            </ul>
            <hr class="profile-sidebar-divider">
        <ul>
            <li><a href=".">Settings</a></li>
        </ul>
        <hr class="profile-sidebar-divider">
        <p id="copyright">Nectlife &copy;${new Date().getFullYear()}.<br>All rights reserved.</p>
    `;

    const homeIconPlaceholder = sidebar.querySelector(
        "#home-icon-placeholder",
    )!;
    homeIconPlaceholder.replaceWith(Icon({ name: "home", size: "1rem" }));

    return sidebar;
};
