import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { App } from "./app";
import CONFIGS from "./config";

if (CONFIGS.FAVICON_URL) {
  const favicon =
    document.querySelector('link[rel="icon"]') ||
    document.createElement("link");
  favicon.setAttribute("rel", "icon");
  favicon.setAttribute("href", CONFIGS.FAVICON_URL);
  document.head.appendChild(favicon);
}

if (CONFIGS.ORGANIZATION_NAME) {
  document.title = `${CONFIGS.ORGANIZATION_NAME} – Event Destinations Portal`;
}

const toHSL = (hex: string) => {
  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
  if (!result) return null;

  let r = parseInt(result[1], 16);
  let g = parseInt(result[2], 16);
  let b = parseInt(result[3], 16);

  r /= 255;
  g /= 255;
  b /= 255;

  const max = Math.max(r, g, b);
  const min = Math.min(r, g, b);
  let h, s;
  const l = (max + min) / 2;

  if (max === min) {
    h = s = 0;
  } else {
    const d = max - min;
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min);
    switch (max) {
      case r:
        h = (g - b) / d + (g < b ? 6 : 0);
        break;
      case g:
        h = (b - r) / d + 2;
        break;
      case b:
        h = (r - g) / d + 4;
        break;
    }
    h! /= 6;
  }

  return { h: h! * 360, s: s * 100, l: l * 100 };
};

const generateColorVariants = (brandColor: string) => {
  const hsl = toHSL(brandColor);
  if (!hsl) return;

  const { h, s, l } = hsl;

  return {
    light: {
      primary: `hsl(${h}, ${s}%, ${l}%)`,
      primaryHover: `hsl(${h}, ${s}%, ${l - 8}%)`,
      containerPrimary: `hsl(${h}, ${s}%, 90%)`,
      containerPrimaryHover: `hsl(${h}, ${s}%, 70%)`,
      foregroundPrimary: `hsl(${h}, ${s}%, ${l}%)`,
      foregroundContainerPrimary: `hsl(${h}, ${s}%, ${l - 12}%)`,
      outlinePrimary: `hsl(${h}, ${s}%, 80%)`,
    },
    dark: {
      primary: `hsl(${h}, ${s}%, ${l}%)`,
      primaryHover: `hsl(${h}, ${s}%, ${l + 8}%)`,
      containerPrimary: `hsl(${h}, ${s}%, 20%)`,
      containerPrimaryHover: `hsl(${h}, ${s}%, 25%)`,
      foregroundPrimary: `hsl(${h}, ${s}%, ${l}%)`,
      foregroundContainerPrimary: `hsl(${h}, ${s}%, ${l + 12}%)`,
      outlinePrimary: `hsl(${h}, ${s}%, 30%)`,
    },
  };
};

const searchParams = new URLSearchParams(window.location.search);
const queryTheme = CONFIGS.FORCE_THEME || searchParams.get("theme");

if (queryTheme === "dark" || queryTheme === "light") {
  // Save new theme preference
  localStorage.setItem("theme", queryTheme);
  document.body.setAttribute("data-theme", queryTheme);
} else {
  // Use saved theme preference, default to light if none exists
  const savedTheme = localStorage.getItem("theme") ?? "light";
  document.body.setAttribute("data-theme", savedTheme);
}

if (CONFIGS.BRAND_COLOR) {
  const colors = generateColorVariants(CONFIGS.BRAND_COLOR);
  if (colors) {
    const root = document.body;

    console.log(document.body.dataset.theme);

    if (document.body.dataset.theme === "dark") {
      // Dark theme colors
      root.style.setProperty(
        "--colors-background-primary",
        colors.dark.primary
      );
      root.style.setProperty(
        "--colors-background-primary-hover",
        colors.dark.primaryHover
      );
      root.style.setProperty(
        "--colors-background-container-primary",
        colors.dark.containerPrimary
      );
      root.style.setProperty(
        "--colors-background-container-primary-hover",
        colors.dark.containerPrimaryHover
      );
      root.style.setProperty(
        "--colors-foreground-primary",
        colors.dark.foregroundPrimary
      );
      root.style.setProperty(
        "--colors-foreground-container-primary",
        colors.dark.foregroundContainerPrimary
      );
      root.style.setProperty(
        "--colors-outline-primary",
        colors.dark.outlinePrimary
      );
    } else {
      // Light theme colors
      root.style.setProperty(
        "--colors-background-primary",
        colors.light.primary
      );
      root.style.setProperty(
        "--colors-background-primary-hover",
        colors.light.primaryHover
      );
      root.style.setProperty(
        "--colors-background-container-primary",
        colors.light.containerPrimary
      );
      root.style.setProperty(
        "--colors-background-container-primary-hover",
        colors.light.containerPrimaryHover
      );
      root.style.setProperty(
        "--colors-foreground-primary",
        colors.light.foregroundPrimary
      );
      root.style.setProperty(
        "--colors-foreground-container-primary",
        colors.light.foregroundContainerPrimary
      );

      root.style.setProperty(
        "--colors-outline-primary",
        colors.light.outlinePrimary
      );

      root.style.setProperty(
        "--colors-shadow-button-primary",
        "0px 1px 2px 0px rgba(0, 0, 0, 0.16)"
      );
    }
  }
}

const container = document.getElementById("root") as HTMLElement;

const root = createRoot(container);

root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
