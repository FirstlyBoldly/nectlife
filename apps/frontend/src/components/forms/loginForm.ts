import { Button } from "../button";
import "./loginForm.css";

export const LoginForm = (): HTMLFormElement => {
    const loginForm = document.createElement("form");
    loginForm.classList.add("basic-container", "login-form");

    const studentIdSubcontainer = document.createElement("div");
    studentIdSubcontainer.classList.add("form-subcontainer");

    const studentIdInput = document.createElement("input");
    studentIdInput.id = "student-id";
    studentIdInput.required = true;
    studentIdInput.autofocus = true;
    studentIdInput.autocomplete = "username";

    const studentIdLabel = document.createElement("label");
    studentIdLabel.textContent = "Student Id";
    studentIdLabel.htmlFor = "student-id";

    studentIdSubcontainer.append(studentIdLabel, studentIdInput);

    const passwordSubcontainer = document.createElement("div");
    passwordSubcontainer.classList.add("form-subcontainer");

    const passwordInput = document.createElement("input");
    passwordInput.id = "password";
    passwordInput.type = "password";
    passwordInput.required = true;
    passwordInput.autocomplete = "current-password";

    const passwordLabel = document.createElement("label");
    passwordLabel.textContent = "Password";
    passwordLabel.htmlFor = "password";

    passwordSubcontainer.append(passwordLabel, passwordInput);

    const submit = Button({
        label: "Login",
        onCLick: (): void => {},
        type: "submit",
    });

    loginForm.append(studentIdSubcontainer, passwordSubcontainer, submit);
    return loginForm;
};
