import "./header.css";
import { router, type Status } from "../../router";
import { Button, type ButtonTheme } from "../button";
import { Icon } from "../icon/icon";

const renderAuthButton = (
    headerElement: HTMLDivElement,
    status: Status,
): void => {
    let label: string;
    let pathname: string;
    let theme: ButtonTheme;
    if (status.authenticated) {
        label = "Logout";
        pathname = "/logout";
        theme = "secondary";
    } else {
        label = "Login";
        pathname = "/login";
        theme = "info";
    }

    const button = Button({
        label: label,
        theme: theme,
        onCLick: (): void => {
            router.navigate(pathname);
        },
    });

    const authButtonContainer = headerElement.querySelector(
        `#auth-button-container`,
    ) as HTMLButtonElement;
    authButtonContainer.innerHTML = "";
    authButtonContainer.appendChild(button);
};

export const Header = (): HTMLDivElement => {
    const headerElement = document.createElement("div");
    headerElement.classList.add("header");
    headerElement.innerHTML = `
        <div id="logo-container">
            <a href="/" id="logo">
                <img src="/nectgrams.png" id="logo-image">
                <span id="logo-text">Nectlife</span>
            </a>
        </div>
        <form id="search-bar-form" method="get" action="/search">
            <div id="search-bar-container">
                <button id="search-button" type="submit"><div id="search-button-icon-placeholder"></div></button>
                <input id="search-input" type="text" placeholder="..." autocomplete="off" required>
            </div>
        </form>
        <ul id="header-list">
            <li id="header-list-item"><div id="write-button-placeholder"></div></li>
            <li id="header-list-item"><div id="auth-button-container"></div></li>
        </ul>
    `;

    const writeButton = Button({
        label: "Write",
        theme: "success",
        onCLick: (): void => {
            router.navigate("/articles/new");
        },
    });

    let placeholder = headerElement.querySelector("#write-button-placeholder")!;
    placeholder.replaceWith(writeButton);

    placeholder = headerElement.querySelector(
        "#search-button-icon-placeholder",
    )!;
    placeholder.replaceWith(
        Icon({ name: "search", size: "1.5rem", color: "var(--primary-color)" }),
    );

    return headerElement;
};

export const reloadHeader = (status: Status): void => {
    const headerElement = document.querySelector(".header") as HTMLDivElement;
    if (headerElement) {
        renderAuthButton(headerElement, status);
    }
};
