# @dev multi-stage build with cache

##############################
# phase 1: install dependencies with package manager
##############################
FROM node:22.14-alpine AS deps
WORKDIR /fatcat

COPY package.json pnpm-lock.yaml svelte.config.js vite.config.ts tsconfig.json components.json ./

ENV COREPACK_HOME=/fatcat/.cache/corepack

RUN npm i -g corepack@latest
RUN corepack enable
RUN pnpm config set store-dir /fatcat/.pnpm-store

# prefetch deps to virtual store and use it as cache
RUN pnpm fetch
RUN --mount=type=cache,id=pnpm,target=/fatcat/.pnpm-store pnpm i --frozen-lock --offline

##############################
# phase 2: build the application with cache
##############################
FROM deps AS builder
WORKDIR /fatcat

COPY .env.production.local .
COPY . .
RUN pnpm build
RUN rm -f .env.production.local

##############################
# phase 3: set permissions and run the application
##############################
FROM node:22.14-alpine AS runner
WORKDIR /fatcat

COPY --from=builder /fatcat/build build

RUN mkdir -p /logs
RUN addgroup -g 1002 fatcat && adduser -u 1002 -G fatcat -S fatcat
RUN chown -R fatcat:fatcat /fatcat/build
USER fatcat

ENV PORT=3011
EXPOSE 3011
CMD ["node", "/fatcat/build"]