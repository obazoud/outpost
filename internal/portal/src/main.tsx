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

const container = document.getElementById("root") as HTMLElement;

const root = createRoot(container);

root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
