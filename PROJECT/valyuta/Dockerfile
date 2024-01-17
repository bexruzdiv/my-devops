# Use the official Python image as a base image
FROM python:3.10

# Set the working directory in the container
WORKDIR /app

# Install dependencies
RUN pip install aiogram==2.25.1 requests countryflag

# Copy the current directory contents into the container at /app
COPY . /app

# Expose the port the app runs on
EXPOSE 8080

# Define environment variable to avoid buffering
ENV PYTHONUNBUFFERED 1

# Run the application
CMD ["python", "test.py"]
