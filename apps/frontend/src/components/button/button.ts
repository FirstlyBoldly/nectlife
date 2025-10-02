import "./button.css";

export type ButtonTheme =
    | "primary"
    | "secondary"
    | "success"
    | "dark"
    | "warning"
    | "white"
    | "info"
    | "mute"
    | "error";

export interface ButtonProps {
    label: string;
    type?: "button" | "submit" | "reset";
    theme?: ButtonTheme;
    onCLick: () => void;
}

export const Button = (props: ButtonProps): HTMLButtonElement => {
    const button = document.createElement("button");
    button.classList.add(
        "basic-button",
        props.theme ? `theme-${props.theme}` : "theme-secondary",
    );
    button.type = props.type || "button";
    button.textContent = props.label;
    button.addEventListener("click", props.onCLick);
    button.addEventListener("keydown", (event) => {
        if (event.key === "Enter") {
            return;
        }
    });
    return button;
};
