**Multi-stage Dockerfile**  
A multi-stage Dockerfile is useful for creating lean and efficient Docker images by separating the build process into multiple stages.  
Each stage can use a different base image and build upon the previous stage, allowing you to minimize the final image size  
by including only the artifacts needed to run the application while excluding unnecessary build tools and intermediate files.

**Example:**
```
# Stage 1: Build and collect static files
FROM python:3.9-slim as builder

# Set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Set work directory
WORKDIR /app

# Install system dependencies
RUN apt-get update \
    && apt-get install -y --no-install-recommends gcc libpq-dev \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY requirements.txt .
RUN pip install --upgrade pip \
    && pip install -r requirements.txt

# Copy project code
COPY . .

# Collect static files
RUN python manage.py collectstatic --noinput

# Stage 2: Setup the production environment
FROM python:3.9-slim

# Create and switch to a new user
RUN useradd -m myuser
USER myuser

# Set work directory
WORKDIR /app

# Copy virtual env and static files from builder stage
COPY --from=builder /app/staticfiles staticfiles
COPY --from=builder /root/.local /root/.local

# Copy project code
COPY . .

# Set environment path
ENV PATH="/root/.local/bin:$PATH"

# Expose port
EXPOSE 8000

# Run gunicorn
CMD ["gunicorn", "--bind", "0.0.0.0:8000", "myproject.wsgi:application"]
```