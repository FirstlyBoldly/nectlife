import { Toast, type Levels } from "../../components/toast";

export const ToastView = (level: Levels, ...text: string[]) => {
    const oldToasts = document.querySelectorAll(".toast");
    for (const oldToast of oldToasts) {
        oldToast.remove();
    }

    const element = document.querySelector("body")!;
    const toast = Toast({ level: level, text: text.join(": ") });
    toast.classList.add("toast-initial-state");
    element.prepend(toast);
    const dismissButton = document.querySelector(".toast-dismiss-button") as HTMLButtonElement;
    if (dismissButton) {
        dismissButton.focus();
    }

    setTimeout(() => {
        toast.classList.add("toast-transition-state");
    }, 10);
};
