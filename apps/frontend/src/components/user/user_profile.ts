import "./user_profile.css";
import type { User } from "../../services/api";

export const UserProfile = (user: User): HTMLElement => {
    const profile = document.createElement("div");
    profile.classList.add("basic-container", "user-profile");

    const studentId = document.createElement("span");
    studentId.textContent = user.StudentID;

    const fullNameContainer = document.createElement("div");

    const firstName = document.createElement("span");
    firstName.textContent = user.FirstName;

    let middleName: string | Node = "";
    if (user.MiddleName) {
        middleName = document.createElement("span");
        middleName.textContent = user.MiddleName.String;
    }

    const lastName = document.createElement("span");
    lastName.textContent = user.LastName;

    fullNameContainer.append(firstName, middleName, lastName);

    const displayName = document.createElement("h2");
    if (user.DisplayName) {
        displayName.textContent = user.DisplayName.String;
    } else {
        displayName.textContent = user.FirstName + " " + user.LastName;
    }

    profile.append(displayName, fullNameContainer, studentId);

    return profile;
};
