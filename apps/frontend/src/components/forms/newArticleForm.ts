import "./newArticleForm.css";
import DOMPurify from "dompurify";
import { marked } from "marked";
import { Button } from "../button";
import { MockImageUpload } from "../../services/mockApi";

export const NewPostForm = (): HTMLFormElement => {
    const form = document.createElement("form");
    form.classList.add("basic-container", "new-article-form");

    const fileInput = document.createElement("input");
    fileInput.type = "file";
    fileInput.accept = "image/*";
    fileInput.style.display = "none";

    const imageButton = Button({
        label: "Upload Image",
        onCLick: (): void => {
            fileInput.click();
        },
    });

    fileInput.addEventListener("change", async () => {
        const file = fileInput.files?.[0];
        if (!file) {
            return;
        }

        imageButton.textContent = "Stuffing the Image!";
        imageButton.disabled = true;

        try {
            const imageURL = await MockImageUpload(file);
            const image = `\n<img src="${imageURL}" alt="${file.name}" width="" height="">\n`;
            contentTextarea.value += image;
            contentTextarea.dispatchEvent(new Event("input"));
        } catch (e) {
            alert(e);
        } finally {
            imageButton.textContent = "Upload Image";
            imageButton.disabled = false;
        }
    });

    const titleContainer = document.createElement("div");
    titleContainer.classList.add("form-subcontainer");

    const titleLabel = document.createElement("label");
    titleLabel.textContent = "Title";
    titleLabel.htmlFor = "post-title";
    const titleInput = document.createElement("input");
    titleInput.type = "text";
    titleInput.id = "post-title";
    titleInput.name = "title";
    titleInput.placeholder = "Title be here...";
    titleInput.required = true;
    titleInput.autofocus = true;

    titleContainer.append(titleLabel, titleInput);

    const contentContainer = document.createElement("div");
    contentContainer.classList.add("form-subcontainer");

    const contentLabel = document.createElement("label");
    contentLabel.textContent = "Content";
    contentLabel.htmlFor = "post-content";
    const contentTextarea = document.createElement("textarea");
    contentTextarea.id = "post-content";
    contentTextarea.name = "content";
    contentTextarea.required = true;
    contentTextarea.style.resize = "vertical";

    contentContainer.append(contentLabel, contentTextarea);

    const previewContainer = document.createElement("div");
    previewContainer.classList.add("form-subcontainer", "form-preview");

    const previewLabel = document.createElement("h3");
    previewLabel.textContent = "Preview";
    const preview = document.createElement("div");
    preview.id = "post-preview";
    preview.classList.add("markdown");
    preview.dataset.placeholder = "Markdown will appear here!!!";

    previewContainer.append(previewLabel, preview);

    const submit = Button({
        label: "Create Post",
        onCLick: (): void => {
            form.submit();
        },
    });

    const formState = {
        title: "Untitled Post",
        content: "",
    };

    titleInput.addEventListener("input", (e) => {
        formState.title = (e.target as HTMLInputElement).value;
    });

    contentTextarea.addEventListener("input", (e) => {
        formState.content = (e.target as HTMLTextAreaElement).value;
        const rawHtml = marked.parse(formState.content);
        const cleanHtml = DOMPurify.sanitize(rawHtml.toString());
        preview.innerHTML = cleanHtml;
    });

    form.addEventListener("submit", (e) => {
        e.preventDefault();
        console.log("New post submitted.");
        alert(`New Post Created: ${formState.title}.`);
    });

    const container = document.createElement("div");
    container.classList.add("form-container");

    const editorContainer = document.createElement("div");
    editorContainer.classList.add("form-editor", "form-subcontainer");
    editorContainer.append(titleContainer, contentContainer);

    container.append(editorContainer, previewContainer);

    form.append(imageButton, container, submit);
    return form;
};
