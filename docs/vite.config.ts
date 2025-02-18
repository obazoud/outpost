import { defineConfig, loadEnv } from "vite";
import process from "node:process";

const appendToHead = (append: string) => {
  return {
    name: "html-transform",
    transformIndexHtml(html) {
      return html.replace(/<\/head>/, `${append}</head>`);
    },
  };
};

export default defineConfig(({ mode }) => {
  // Load env file based on `mode` in the current working directory.
  // Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
  const env = loadEnv(mode, process.cwd(), "");
  return {
    plugins: [appendToHead(env.PUBLIC_CUSTOM_HEAD_CONTENT || "")],
  };
});
