export const Banner = (text: string): void => {
    const banner = document.getElementById("banner")!;
    banner.innerHTML = `
        <h1 id="banner-text">${text}</h1>
    `;
};
