import { Banner } from "../../components/banner";
import { LoginForm } from "../../components/forms/loginForm";
import { router } from "../../router";
import { login } from "../../services/api";
import { ToastView } from "../toast";

export const LoginView = (target: HTMLElement): void => {
    const loginForm = LoginForm();
    loginForm.addEventListener("submit", async (event) => {
        event.preventDefault();
        const studentId = (
            document.getElementById("student-id") as HTMLInputElement
        ).value.toLowerCase();
        const password = (
            document.getElementById("password") as HTMLInputElement
        ).value;
        const { data, error } = await login({
            body: {
                password: password,
                student_id: studentId,
            },
        });
        if (error) {
            ToastView("error", error.title!, error.detail!);
            return;
        }

        const urlParams = new URLSearchParams(window.location.search)
        const redirectTo = urlParams.get("next") || `/users/${data.ID}`

        router.navigate(redirectTo);
    });

    Banner("Login");
    target.innerHTML = "";
    target.append(loginForm);
};
