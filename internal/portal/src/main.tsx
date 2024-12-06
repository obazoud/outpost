import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { App } from "./app";

if (FAVICON_URL) {
  const favicon =
    document.querySelector('link[rel="icon"]') ||
    document.createElement("link");
  favicon.setAttribute("rel", "icon");
  favicon.setAttribute("href", FAVICON_URL);
  document.head.appendChild(favicon);
}

const container = document.getElementById("root") as HTMLElement;

const root = createRoot(container);

root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
