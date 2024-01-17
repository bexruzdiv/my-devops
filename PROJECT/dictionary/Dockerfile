# Use the official Python image as a base image
FROM python:3.8-slim

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install necessary dependencies
RUN pip install --no-cache-dir aiogram==2.25.1 requests

# Define environment variable for bot token
ENV BOT_TOKEN "6944612215:AAEkEraevZOqDpEV8-kqZMWthPKUkxPd1Tc"

# Run your Python script
CMD ["python", "main.py"]
