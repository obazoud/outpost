import { Outpost } from "@hookdeck/outpost-sdk";

let outpostClient: Outpost | null = null;

export function getOutpostClient(): Outpost {
  if (!outpostClient) {
    if (!process.env.OUTPOST_API_KEY) {
      throw new Error("OUTPOST_API_KEY environment variable is required");
    }

    outpostClient = new Outpost({
      serverURL: `${
        process.env.OUTPOST_BASE_URL || "http://localhost:3333"
      }/api/v1`,
      security: {
        adminApiKey: process.env.OUTPOST_API_KEY,
      },
    });
  }

  return outpostClient;
}

export async function createTenant(tenantId: string): Promise<void> {
  try {
    console.log(`🔄 Creating tenant in Outpost: ${tenantId}`);
    const outpost = getOutpostClient();

    const tenant = await outpost.tenants.upsert({
      tenantId,
    });
    console.log(`✅ Tenant created successfully in Outpost:`, tenant);
  } catch (error) {
    console.error(`❌ Error creating tenant in Outpost: ${tenantId}`, error);
    throw error;
  }
}

export async function getPortalUrl(
  tenantId: string,
  theme?: string
): Promise<string> {
  try {
    const outpost = getOutpostClient();
    const result = await outpost.tenants.getPortalUrl({
      tenantId,
      theme: theme as any, // SDK accepts theme as string but types are restrictive
    });
    return result.redirectUrl || "";
  } catch (error) {
    console.error("Error getting portal URL:", error);
    throw error;
  }
}

export async function getTenantOverview(tenantId: string) {
  try {
    console.log(`🔄 Getting tenant overview for: ${tenantId}`);
    const outpost = getOutpostClient();

    // Get tenant details
    const tenant = await outpost.tenants.get({ tenantId });
    console.log(`✅ Tenant found:`, tenant);

    // Get destinations for this tenant
    const destinationsResponse = await outpost.destinations.list({ tenantId });
    console.log(`✅ Destinations found:`, destinationsResponse);
    
    // Transform destinations to match our interface
    const destinations = Array.isArray(destinationsResponse) 
      ? destinationsResponse.map((dest: any) => ({
          ...dest,
          enabled: dest.disabledAt === null // Convert disabledAt to enabled boolean
        }))
      : [];

    // Get recent events with proper error handling
    let recentEvents: any[] = [];
    let totalEvents = 0;

    try {
      const eventsResponse = await outpost.events.list({ tenantId });
      // Handle SDK response structure correctly
      const eventsData = Array.isArray(eventsResponse) ? eventsResponse : (eventsResponse as any)?.data || [];
      recentEvents = eventsData.slice(0, 10);
      totalEvents = eventsData.length;
      console.log(`✅ Events found: ${totalEvents}`);
    } catch (error) {
      console.warn("Could not fetch events:", error);
    }

    const overview = {
      tenant,
      destinations: Array.isArray(destinations) ? destinations : [],
      recentEvents,
      stats: {
        totalDestinations: Array.isArray(destinations)
          ? destinations.length
          : 0,
        totalEvents,
      },
    };

    console.log(`✅ Complete overview for tenant: ${tenantId}`, overview);
    return overview;
  } catch (error) {
    console.error("Error getting tenant overview:", error);
    throw error;
  }
}
