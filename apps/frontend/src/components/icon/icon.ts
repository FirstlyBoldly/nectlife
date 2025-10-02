import { ICONS } from "../../lib/icons";

export interface IconProps {
    name: keyof typeof ICONS;
    size?: number | string;
    color?: string;
}

export const Icon = (props: IconProps): SVGSVGElement => {
    const { name, size = 24, color = "currentColor" } = props;
    const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");

    svg.setAttribute("width", String(size));
    svg.setAttribute("height", String(size));
    svg.setAttribute("viewBox", "0 0 24 24");
    svg.setAttribute("fill", color);

    const pathData = ICONS[name];
    if (pathData) {
        svg.innerHTML = pathData;
    } else {
        console.warn(`Icon "${name}" not found.`);
    }

    return svg;
};
