import { Button } from "../button";
import { Icon } from "../icon/icon";
import "./toast.css";

export type Levels = "warning" | "info" | "success" | "error";

export type ToastProps = {
    level: Levels;
    text: string;
    actionButtons?: HTMLButtonElement[];
};

export const Toast = (props: ToastProps): HTMLDivElement => {
    const toastContainer = document.createElement("div");
    toastContainer.classList.add("basic-container", "toast", `theme-${props.level}`);

    const level = Icon({
        name: `${props.level}`,
        size: "2rem",
        color: `var(--${props.level}-color)`
    });

    const text = document.createElement("span");
    text.textContent = props.text;

    let actionButtons: HTMLButtonElement[] = [];
    if (props.actionButtons) {
        actionButtons = props.actionButtons;
    }

    const dismiss = Button({
        label: "",
        theme: `${props.level}`,
        onCLick: (): void => {
            setTimeout(() => {
                toastContainer.classList.add("toast-remove-state");
            }, 10);
            setTimeout(() => {
                toastContainer.remove();
            }, 510);
        }
    });

    dismiss.append(Icon({
        name: "close",
        size: "1.5rem",
    }));
    dismiss.classList.add("toast-dismiss-button");

    toastContainer.append(level, text, ...actionButtons, dismiss);
    return toastContainer
};
