import * as dotenv from "dotenv";
dotenv.config();

const organizations = [{ id: "org1" }, { id: "org2" }, { id: "org3" }];

const subscriptions = [
  {
    organizationId: "org1",
    url: process.env.ORG_1_ENDPOINT_1 || "https://org1.test/users",
    topics: ["user.created", "user.updated"],
    secret: "some_secret_value",
  },
  {
    organizationId: "org1",
    url: process.env.ORG_1_ENDPOINT_2 || "https://org1.test/users",
    topics: ["user.deleted"],
    secret: "some_secret_value",
  },
  {
    organizationId: "org2",
    url: process.env.ORG_2_ENDPOINT || "https://org1.test/users",
    topics: ["user.created", "user.updated", "user.deleted"],
    secret: "some_secret_value",
  },
];

class Database {
  getOrganizations() {
    return organizations;
  }

  getSubscriptions(organizationId?: string) {
    if (!organizationId) {
      return subscriptions;
    }
    return subscriptions.filter(
      (subscription) => subscription.organizationId === organizationId
    );
  }
}

export default new Database();
