# Airline Voucher Seat Service Frontend

## Summary

This is a web frontend that serves as a user interface for airline crew members to generate vouchers. Vouchers will be automatically generated for 3 random seats after the crew enters the following data:
- Crew Name
- Crew ID
- Flight Number
- Flight Date
- Aircraft type selection

## Tech Stack

**Language** : Typescript
**Framework** : React with Next.js  
**Styling** : Tailwind CSS 
**HTTP Client** : Axios

## Installation

1. Clone the repository:
```bash
git clone <repository-url> <project>
cd <project>/frontend
```

2. Install dependencies:
```bash
npm install
# or
yarn install
# or
pnpm install
# or
bun install
```

3. Create environment file:
```bash
cp .env.example .env
```

4. Configure environment variables in `.env`:
```
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
```

5. Run the development server:
```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

6. Open [http://localhost:3000](http://localhost:3000) in your browser.

## Build for Production
**Note**: Ensure the .env file has been copied and configured as described in the [Installation](#installation) section above.

### Standard Build (Server-Side Rendering)
1. Build the application:  
```bash
npm run build
# or
yarn build
# or
pnpm build
# or
bun run build
```

2. Start the production server:
```bash
npm start
# or
yarn start
# or
pnpm start
# or
bun start
```

3. The production application will be available at [http://localhost:3000](http://localhost:3000).

### Static Export (for CDN/Static Hosting)
1. Add export configuration to `next.config.ts`:
```typescript
/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'export',
}

module.exports = nextConfig
```

2. Build and export:
```bash
npm run build
# or
yarn build
# or
pnpm build
# or
bun run build
```

3. The static files will be generated in the `out/` directory, ready for deployment to any static hosting service.

### Standalone Build (for Docker/Serverless)
1. Add standalone configuration to `next.config.ts`:
```typescript
/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone'
}

module.exports = nextConfig
```

2. Build the application:
```bash
npm run build
# or
yarn build
# or
pnpm build
# or
bun run build
```

3. The standalone build will be available in `.next/standalone/` directory, optimized for containerized deployments.


