# Use an official MongoDB image as the base image
FROM mongo:latest

# Expose the default MongoDB port
EXPOSE 27017

# Set the default environment variables for MongoDB
ENV MONGO_INITDB_ROOT_USERNAME=admin
ENV MONGO_INITDB_ROOT_PASSWORD=password

# Create a directory for storing MongoDB data
RUN mkdir -p /data/db

# Start MongoDB when the container starts
CMD ["mongod"]
