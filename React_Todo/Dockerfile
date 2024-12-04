# Build the React app
FROM node:16 as build

WORKDIR /app

# Copy package.json and install dependencies
COPY package.json .
RUN npm install

# Copy source code and build the React app
COPY . .
RUN npm run build

# Use Nginx to serve the React app
FROM nginx:alpine

# Copy the build output to Nginx
COPY --from=build /app/build /usr/share/nginx/html

# Expose Nginx's port
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
