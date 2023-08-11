# build stage
FROM node:18-alpine3.17 as build-stage
WORKDIR /app
COPY dist .

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]