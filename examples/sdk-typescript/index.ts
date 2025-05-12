import dotenv from "dotenv";
dotenv.config();
import { Outpost } from "@hookdeck/outpost-sdk";

const TENANT_ID = process.env.TENANT_ID;
const ADMIN_API_KEY = process.env.ADMIN_API_KEY;
const SERVER_URL = process.env.SERVER_URL;

if (!process.env.ADMIN_API_KEY) {
  console.error("Please set the ADMIN_API_KEY environment variable.");
  process.exit(1);
}
if (!process.env.SERVER_URL) {
  console.error("Please set the SERVER_URL environment variable.");
  process.exit(1);
}
if (!process.env.TENANT_ID ) {
  console.error("Please set the TENANT_ID environment variable.");
  process.exit(1);
}

const debugLogger = {
  debug: (message: string) => {
    console.log("DEBUG", message);
  },
  group: (message: string) => {
    console.group(message);
  },
  groupEnd: () => {
    console.groupEnd();
  },
  log: (message: string) => {
    console.log(message);
  }
};

const withJwt  = async (jwt: string) => {
  const outpost = new Outpost({security: {tenantJwt: jwt}, serverURL: `${SERVER_URL}/api/v1`, tenantId: TENANT_ID });
  const destinations = await outpost.destinations.list({});

  console.log(destinations);
}

const withAdminApiKey  = async () => {
  const outpost = new Outpost({security: {adminApiKey: ADMIN_API_KEY}, serverURL: `${SERVER_URL}/api/v1`});

  const result = await outpost.health.check();
  console.log(result);

  const destinations = await outpost.destinations.list({tenantId: TENANT_ID})

  console.log(destinations);

  const jwt = await outpost.tenants.getToken({tenantId: TENANT_ID});

  await withJwt(jwt.token!);
}

const main = async () => {
  await withAdminApiKey();
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
