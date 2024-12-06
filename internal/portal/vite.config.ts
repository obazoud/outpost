import { defineConfig, UserConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

function generateScssVariables(env: Record<string, string | undefined>) {
  const portal_css_vars = Object.entries(env).filter(([key]) =>
    key.startsWith("PORTAL_CSS_")
  );

  return portal_css_vars
    .map(([key, value]) => {
      // Just remove PORTAL_CSS_ prefix but keep the rest of the name
      const scss_var_name = key.replace("PORTAL_CSS_", "");
      return `$${scss_var_name}: ${value};`;
    })
    .join("\n");
}

export default defineConfig(() => {
  const config: UserConfig = {
    plugins: [react()],
    server: {
      port: 3334,
    },
    define: {
      REFERER_URL: JSON.stringify(process.env.PORTAL_REFERER_URL),
      FAVICON_URL: JSON.stringify(process.env.PORTAL_FAVICON_URL),
      LOGO: JSON.stringify(process.env.PORTAL_LOGO),
      ORGANIZATION_NAME: JSON.stringify(process.env.PORTAL_ORGANIZATION_NAME),
      FORCE_THEME: JSON.stringify(process.env.PORTAL_FORCE_THEME),
      TOPICS: JSON.stringify(process.env.TOPICS),
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: generateScssVariables(process.env),
        },
      },
    },
  };

  return config;
});
