# Use the official Python image as a base image
FROM python:3.8-slim

# Set the working directory in the container
WORKDIR /app

# Install any needed packages
RUN pip install flask flask_cors psycopg2-binary pandas

# Copy the current directory contents into the container at /app
COPY . /app/

# Expose port 5000 for the Flask app
EXPOSE 5000

# Command to run your application
CMD ["python", "api.py"]
