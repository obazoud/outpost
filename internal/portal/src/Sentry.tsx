/// <reference types="vite/client" />

import * as Sentry from "@sentry/react";
import CONFIGS from "./config";

Sentry.init({
  dsn: "https://35671160affeb03460b893b7e9e59439@o376420.ingest.us.sentry.io/4507984881385472", // Outpost maintainers managed DSN for Sentry
  integrations: [],
  enabled: !CONFIGS.DISABLE_TELEMETRY,
  environment: import.meta.env.MODE,
});

Sentry.setContext("Portal Config", {
  organization_name: CONFIGS.ORGANIZATION_NAME,
  force_theme: CONFIGS.FORCE_THEME,
  referer_url: CONFIGS.REFERER_URL,
});

export default Sentry;
