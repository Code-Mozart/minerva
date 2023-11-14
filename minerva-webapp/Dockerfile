# Use the official Node.js 21 image as the base image
FROM node:21-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install pnpm
RUN wget -qO- https://get.pnpm.io/install.sh | ENV="$HOME/.shrc" SHELL="$(which sh)" sh -
RUN source ~/.shrc

# Install any needed dependencies
RUN pnpm install

# Expose port 3000
EXPOSE 3000

# Start the application
CMD ["pnpm", "run", "start"]
