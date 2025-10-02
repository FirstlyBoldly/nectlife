import { status } from "./services/api";
import { ToastView } from "./views";
import { reloadHeader } from "./components/header";
import {
    HomeView,
    PostsView,
    PostView,
    LoginView,
    LogoutView,
    UserView,
    NewPostView,
} from "./views";
import { reloadUserProfile } from "./components/home";

export type Status = {
    userId: number | undefined;
    authenticated: boolean;
};

export type RouteParams = {
    [key: string]: string;
};

interface Route {
    path: string;
    view: (element: HTMLElement, params: RouteParams) => void;
    protected?: boolean;
    pageLayout?: "default" | "full-width" | "boxed";
}

const emptyStatus: Status = {
    authenticated: false,
    userId: undefined,
};

const updateStatus = async (): Promise<Status> => {
    try {
        const { data, error } = await status();
        if (error) {
            console.error(error);
            return emptyStatus;
        } else {
            return {
                userId: data.UserID,
                authenticated: data.Authenticated,
            };
        }
    } catch (error) {
        console.error(error);
        return emptyStatus;
    }
};

class Router {
    private routes: Route[];
    private contentRoot: HTMLElement;

    constructor(routes: Route[], contentRoot: HTMLElement) {
        this.routes = routes;
        this.contentRoot = contentRoot;
    }

    listen() {
        window.addEventListener("popstate", async () => await this.resolve());
        window.addEventListener(
            "DOMContentLoaded",
            async () => await this.resolve(),
        );
    }

    navigate(pathname: string): void {
        history.pushState(null, "", pathname);
        this.resolve();
    }

    private async resolve() {
        const status = await updateStatus();
        reloadHeader(status);
        reloadUserProfile(status);

        const path = window.location.pathname || "/";
        for (const route of this.routes) {
            const pathRegex = new RegExp(
                `^${route.path.replace(/:\w+/g, "([^\\/]+)")}$`,
            );
            const match = path.match(pathRegex);
            if (match) {
                if (route.protected && !status.authenticated) {
                    this.navigate(`/login?next=${route.path}`);
                    ToastView("warning", `Login to access ${path}.`);
                    return;
                }

                const values = match.slice(1);
                const keys = (route.path.match(/:\w+/g) || []).map((key) =>
                    key.slice(1),
                );
                const params: RouteParams = Object.fromEntries(
                    keys.map((key, i) => [key, values[i]]),
                );

                this.contentRoot.innerHTML = "";
                const leftColumn = document.getElementById(
                    "nav-sidebar-container",
                )!;
                const rightColumn = document.getElementById(
                    "right-sidebar-container",
                )!;
                const contentContainer =
                    document.getElementById("content-container")!;
                switch (route.pageLayout) {
                    case "full-width":
                        leftColumn.classList.add("hidden");
                        rightColumn.classList.add("hidden");
                        contentContainer.style.maxWidth = "";
                        break;

                    case "boxed":
                        leftColumn.classList.add("hidden");
                        rightColumn.classList.add("hidden");
                        contentContainer.style.maxWidth = "768px";
                        break;

                    default:
                        leftColumn.classList.remove("hidden");
                        rightColumn.classList.remove("hidden");
                        contentContainer.style.maxWidth = "";
                        break;
                }
                route.view(this.contentRoot, params);
                return;
            }
        }

        ToastView("error", `Path ${path} not found`);
    }
}

const routes: Route[] = [
    { path: "/", view: HomeView },
    { path: "/articles", view: PostsView },
    {
        path: "/articles/new",
        view: NewPostView,
        protected: true,
        pageLayout: "full-width",
    },
    { path: "/articles/:id", view: PostView },
    { path: "/users/:id", view: UserView },
    { path: "/login", view: LoginView, pageLayout: "boxed" },
    { path: "/logout", view: LogoutView, protected: true },
];
const main = document.getElementById("content")!;
export const router = new Router(routes, main);
