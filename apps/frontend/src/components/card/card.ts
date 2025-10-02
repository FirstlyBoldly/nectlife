import "./card.css";
import { Button } from "../button";

export interface CardProps {
    title: string;
    description?: string;
    list?: string[];
    actionButtons?: HTMLButtonElement[];
}

export const Card = (props: CardProps): HTMLDivElement => {
    const card = document.createElement("div");
    card.classList.add("card", "basic-container");

    const titleElement = document.createElement("h2");
    titleElement.textContent = props.title;

    let descriptionElement: HTMLParagraphElement | string = "";
    if (props.description) {
        descriptionElement = document.createElement("p");
        descriptionElement.textContent = props.description;
    }

    const alertButton = Button({
        label: "Learn More",
        onCLick: () => alert(`You clicked on "${props.title}"`),
    });

    let listElement: HTMLUListElement | string = "";
    if (props.list) {
        listElement = document.createElement("ul");
        for (const element of props.list) {
            const li = document.createElement("li");
            li.innerHTML = element;
            listElement.appendChild(li);
        }
    }

    let actionButtons: HTMLButtonElement[] = [];
    if (props.actionButtons) {
        actionButtons = props.actionButtons;
    }

    card.append(
        titleElement,
        descriptionElement,
        listElement,
        alertButton,
        ...actionButtons,
    );
    return card;
};
