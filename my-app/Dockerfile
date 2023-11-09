FROM node:18.12.0-alpine

# update and install dependency
RUN apk update && apk upgrade
RUN apk add git
RUN apk add bash

# create destination directory
RUN mkdir -p /usr/src/nuxt-app
WORKDIR /usr/src/nuxt-app

# Copy package.json and package-lock.json to the working directory
COPY package*.json /usr/src/nuxt-app/

# Install any needed dependencies specified in package.json
RUN npm install

#SET ENV VARIABLES

# Copy the rest of the application code to the working directory
COPY . /usr/src/nuxt-app/

# Build the application
RUN npm run build

# permission to script bash
RUN chmod +x /usr/src/nuxt-app/scripts/wait-for-it.sh
RUN chmod +x /usr/src/nuxt-app/scripts/wait-for-services-api.sh

EXPOSE 3000

# Run the app
CMD [ "npm", "start" ]


