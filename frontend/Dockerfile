FROM node:18-alpine AS builder
WORKDIR /app
ARG VITE_API_BASE=http://backend:8080
ENV VITE_API_BASE=${VITE_API_BASE}
COPY package.json package-lock.json* ./
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build

FROM nginx:1.25-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 3000
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
