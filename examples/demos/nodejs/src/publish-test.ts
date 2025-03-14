import * as process from "process";

import outpost from "./lib/outpost";
import db from "./lib/db";

const main = async () => {
  const organizations = db.getOrganizations();

  for (const organization of organizations) {
    const destinations = await outpost.getDestinations(organization.id);
    for (const destination of destinations) {
      await outpost.publishEvent({
        tenant_id: organization.id,
        topic: destination.topics[0],
        data: {
          test: "data",
          organization_id: organization.id,
          destination_id: destination.id,
        },
        meta_data: {
          some: "metadata",
        },
      });
    }
  }
};

main()
  .then(() => {
    console.log("Test publishing complete");
    process.exit(0);
  })
  .catch((err) => {
    console.error("Test publishing failed", err);
    process.exit(1);
  });
