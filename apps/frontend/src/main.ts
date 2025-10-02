import "./styles/global.css";
import { router } from "./router";
import { Header } from "./components/header";
import { client } from "./services/api/client.gen";
import { ProfileSidebar, PopularSidebar } from "./components/home";

client.setConfig({
    baseUrl: "http://localhost:8080",
    credentials: "include",
});

const headerContainer = document.getElementById("header-container");
const navSidebarContainer = document.getElementById("nav-sidebar-container");
const rightSidebarContainer = document.getElementById(
    "right-sidebar-container",
);
if (headerContainer && navSidebarContainer && rightSidebarContainer) {
    const header = Header();
    const leftColumn = ProfileSidebar();
    const rightColumn = await PopularSidebar();

    headerContainer.append(header);
    navSidebarContainer.append(leftColumn);
    rightSidebarContainer.append(rightColumn);

    router.listen();
} else {
    console.error("Root HTML element not found");
}
