# API Platform Dashboard Integration Demo

A Next.js application demonstrating how to integrate Outpost with an API platform or SaaS API dashboard. This example shows how API platforms can seamlessly embed Outpost's event destination management capabilities into their existing user dashboards.

## Features

- **API Platform Dashboard**: Complete user dashboard for an API/SaaS platform
- **Integrated Event Management**: Seamless Outpost integration for webhook/event destinations
- **Automatic Tenant Provisioning**: Users get Outpost tenants created automatically when they sign up
- **Unified User Experience**: Event destination management feels like a native part of the platform
- **Production-Ready Patterns**: Demonstrates real-world integration patterns for API platforms

## Prerequisites

- Node.js 18+
- Docker and Docker Compose
- Git

## Setup

1. **Clone and navigate to the project**:
   ```bash
   cd examples/demos/dashboard-integration
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Set up environment variables**:
   ```bash
   # Dashboard application configuration
   cp .env.example .env.local
   
   # Outpost configuration
   cp .env.outpost.example .env.outpost
   ```
   
   The default configurations should work for local development. For production, update the secrets in both files.

4. **Start the complete stack** (PostgreSQL, Redis, RabbitMQ, and Outpost):
   ```bash
   docker-compose up -d
   ```
   
   This will start:
   - `postgres` - PostgreSQL with separate databases for dashboard and Outpost (port 5432)
   - `redis` - Redis for Outpost (port 6379)
   - `rabbitmq` - Message queue for Outpost (port 5672, management UI on 15672)
   - `outpost-api` - Outpost API service (port 3333)
   - `outpost-delivery` - Outpost delivery service
   - `outpost-log` - Outpost log service

5. **Wait for services to be healthy**:
   ```bash
   docker-compose ps
   ```
   Wait until all services show "healthy" status. This may take 1-2 minutes on first startup.

6. **Run the dashboard application**:
   ```bash
   npm run dev
   ```

7. **Access the application**:
   - Dashboard: [http://localhost:4000](http://localhost:4000)
   - Outpost API: [http://localhost:3333](http://localhost:3333)
   - RabbitMQ Management: [http://localhost:15672](http://localhost:15672) (guest/guest)

## Usage

1. **Sign up for the API platform** - Register as a new user of the fictitious API platform
2. **Access your platform dashboard** - View your API usage, account info, and available features
3. **Manage Event Destinations** - Click "Event Destinations" to configure webhooks for your API events
4. **Seamless Portal Experience** - Get redirected to Outpost's full-featured destination management interface

## Use Case

This demo represents a common integration scenario where:

- **API Platform**: Your SaaS provides APIs that generate events (user signups, payments, data processing, etc.)
- **Customer Need**: Your customers need to receive these events via webhooks to their own systems
- **Outpost Integration**: Instead of building webhook infrastructure from scratch, you integrate Outpost
- **User Experience**: Your customers manage event destinations through what feels like your native platform

## Architecture

- **Platform Frontend**: Next.js application representing your API platform's dashboard
- **User Management**: Auth.js with PostgreSQL for platform user accounts
- **Outpost Integration**: Official TypeScript SDK for tenant and portal management
- **Seamless Handoff**: Server-side route handlers for transparent portal access

## Environment Variables

### Dashboard Application (.env.local)

| Variable | Description | Example |
|----------|-------------|---------|
| `DATABASE_URL` | Dashboard PostgreSQL connection | `postgresql://dashboard:dashboard@localhost:5432/dashboard_integration` |
| `NEXTAUTH_SECRET` | Auth.js secret key | `your-secret-here` |
| `OUTPOST_BASE_URL` | Outpost base URL | `http://localhost:3333` |
| `OUTPOST_API_KEY` | Outpost API key | `demo-api-key-change-in-production` |
| `LOG_LEVEL` | Application logging level (`error`, `warn`, `info`, `debug`) | `info` |

### Outpost Configuration (.env.outpost)

| Variable | Description | Example |
|----------|-------------|---------|
| `API_KEY` | Outpost API key | `demo-api-key-change-in-production` |
| `API_JWT_SECRET` | JWT signing secret | `demo-jwt-secret-change-in-production` |
| `POSTGRES_URL` | Outpost PostgreSQL connection | `postgres://outpost:outpost@postgres:5432/outpost` |
| `REDIS_HOST` | Redis hostname | `redis` |
| `RABBITMQ_SERVER_URL` | RabbitMQ connection | `amqp://guest:guest@rabbitmq:5672` |
| `PORTAL_ORGANIZATION_NAME` | Portal branding | `API Platform Demo` |

## Docker Services

The docker-compose.yml includes the complete stack:

- **postgres** (port 5432): Single PostgreSQL instance with separate databases for dashboard (`dashboard_integration`) and Outpost (`outpost`)
- **redis** (port 6379): Redis for Outpost caching and sessions
- **rabbitmq** (ports 5672, 15672): Message queue for Outpost event processing
- **outpost-api** (port 3333): Outpost API service
- **outpost-delivery**: Outpost delivery service for processing webhooks
- **outpost-log**: Outpost log service for event logging

## Development

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint
- `docker-compose up -d` - Start all services
- `docker-compose down` - Stop all services
- `docker-compose logs [service]` - View logs for a specific service

## Troubleshooting

### Services not starting
```bash
# Check service status
docker-compose ps

# View logs for a specific service
docker-compose logs outpost-api
docker-compose logs dashboard-postgres

# Restart services
docker-compose restart
```

### Database connection issues
```bash
# Check PostgreSQL is ready
docker-compose exec postgres pg_isready -U dashboard -d dashboard_integration
docker-compose exec postgres pg_isready -U outpost -d outpost
```

### Outpost API not responding
```bash
# Check if Outpost API is healthy
curl http://localhost:3333/healthz

# View Outpost logs
docker-compose logs outpost-api
```

## Integration Points

1. **User Onboarding** → Platform user registration automatically provisions Outpost tenant
2. **Dashboard Overview** → Display event/webhook statistics from Outpost API in platform dashboard
3. **Destination Management** → "Event Destinations" nav item seamlessly redirects to Outpost portal
4. **Branded Experience** → Users experience webhook management as part of your platform, not a separate tool

## Key Integration Patterns

This demo illustrates several important patterns for API platform integration:

- **Transparent Tenant Management**: Outpost tenants are created automatically and hidden from users
- **Single Sign-On Experience**: Users don't need separate Outpost accounts
- **Embedded Navigation**: Event destinations appear as a natural part of the platform navigation
- **Contextual Data Display**: Platform dashboard shows relevant Outpost data (destination counts, recent events)
- **Seamless Handoff**: Clicking "Event Destinations" feels like navigating to another page in your app
