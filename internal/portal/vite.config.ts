import { defineConfig, UserConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import { sentryVitePlugin } from "@sentry/vite-plugin";

export default defineConfig(({ mode }) => {
  let plugins: UserConfig["plugins"] = [react()];

  if (mode === "production") {
    plugins.push(
      sentryVitePlugin({
        authToken: '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',
        org: "hookdeck",
        project: "outpost-portal",
        telemetry: false,
        bundleSizeOptimizations: {
          excludeTracing: true,
          excludePerformanceMonitoring: true,
          excludeReplayCanvas: true,
          excludeReplayShadowDom: true,
          excludeReplayIframe: true,
          excludeReplayWorker: true,
        },
      })
    );
  }

  const config: UserConfig = {
    plugins,
    server: {
      port: 3334,
      // hmr: {
      //   port: 3334,
      // },
    },
    build: {
      sourcemap: true,
    },
  };

  return config;
});
