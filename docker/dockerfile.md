**Dockerfile:**  
Dockerfile is a text file that contains a set of instructions to build a Docker image. 


**FROM:**  
Specifies the base image to use for the new image. It is the first instruction in a Dockerfile.
Example: FROM ubuntu:20.04

**RUN:**   
Executes commands in the container's shell during the build process to install packages, 
update repositories, or any other setup steps.  
Example: `RUN apt-get update && apt-get install -y python3`

**COPY or ADD:**   
Copies files or directories from the host machine to the container's filesystem.   
Example: `COPY app.py /app/`

**WORKDIR:**   
Sets the working directory inside the container for subsequent commands.  
Example: `WORKDIR /app`

**ENV:**   
Sets environment variables within the container.
Example: `ENV PORT=80`

**EXPOSE:**  
Specifies the port on which the container listens at runtime (does not publish the port to the host machine).  
Example: `EXPOSE 8080`


**CMD:**   
Specifies the default command to be executed when the container starts. There can only be one CMD instruction in a Dockerfile.  
Example: `CMD ["python3", "app.py"]`

**ENTRYPOINT:**   
Sets the primary command that is executed when the container starts. It allows you to provide arguments to the command specified.  
Example: `ENTRYPOINT ["python3", "app.py"]`

**VOLUME:**   
Creates a mount point with the specified name in the container.  
Example: `VOLUME /data`

**ARG:**   
Defines build-time variables that can be passed to the docker build command using the `--build-arg` flag.  
Example: `ARG VERSION=latest`

**LABEL:**   
Adds metadata to the image, such as version, maintainer, or any other information you want to include.  
Example: `LABEL version="1.0" maintainer="Oded Viner"`


**Dockerfile Example:**
```
# Use the official Python image as the base image
FROM python:3.9

# Set the working directory inside the container
WORKDIR /app

# Copy the requirements.txt file to the container
COPY requirements.txt .

# Install the Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application files to the container
COPY . .

# Expose port 5000 for the Flask web application
EXPOSE 5000

# Set the default command to run the Flask application
CMD ["python", "app.py"]
```