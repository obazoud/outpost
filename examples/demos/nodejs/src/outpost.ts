import axios from 'axios';

const OUTPOST_API_URL = 'localhost:3333/api/v1';
const OUTPOST_API_KEY = 'some-super-secret-api-key';


interface Tenant {
  id: string;
}

class OutpostClient {
  constructor() {}

  async request<T>(path: string, method: string, data: any): Promise<T> {
    const response = await axios.request<T>({
      url: `${OUTPOST_API_URL}${path}`,
      method,
      data,
      headers: {
        'Authorization': `Bearer ${OUTPOST_API_KEY}`,
      },
    });
    return response.data;
  }

  async publishEvent(event_type: string, event_data: any): Promise<boolean> {
    const response = await this.request('/publish', 'POST', {
      event_type,
      event_data,
    });
    return !!response;
  }


  async registerTenant(tenant_id: string): Promise<Tenant> {
    const response = await this.request<Tenant>(`/${tenant_id}`, 'PUT', {});
    return response;
  }

  async getPortalURL(tenant_id: string): Promise<string> {
    const response = await this.request<{ portal_url: string }>(`/${tenant_id}/portal`, 'GET', {});
    return response.portal_url
  }
}


export default new OutpostClient();
