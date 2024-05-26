<p align="right">
    <img src="assets/logo.svg" width="50px" />
</p>

# ComputeSphere Golang InspireQuotes Example App

This example deploys a GO application to ComputeSphere.

> [!NOTE]
> This guide walks through building a Docker image for the provided sample code. Please note that the version `v0.0.1` used in the example is for demonstration purposes only. You should replace it with a version that suits your specific setup and requirements.

## Prerequisites

- A [ComputeSphere](https://computesphere.com) account
- [Git](https://git-scm.com/downloads)
- [GO](https://go.dev/learn/)
- [Docker](https://docs.docker.com/engine/install/)

## Setup

1. Clone this repository.

   ```bash
   git clone https://github.com/computesphere-samples/computesphere-golang-inspirequotes-example.git

   cd computesphere-golang-inspirequotes-example
   ```

2. Create the docker image.

   ```bash
   docker build -t computesphere-golang-inspirequotes-example:v.0.0.1 .
   ```

   Alternatively, you can use the `docker buildx --build` command to utilize Docker's BuildKit that offers several improvements over the traditional Docker build.

   ```bash
   docker buildx build --platform=linux/amd64 --tag computesphere-golang-inspirequotes-example:v0.0.1 .
   ```

3. Push the image to Docker Hub.

   ```bash
   docker tag computesphere-golang-inspirequotes-example:v0.0.1 [REPOSITORY]/computesphere-golang-inspirequotes-example:v0.0.1

   docker push [REPOSITORY]/computesphere-golang-inspirequotes-example:v0.0.1
   ```

> [!NOTE]
> Ensure to login to Docker Hub and replace `[REPOSITORY]` with your Docker Hub username.

## Running the project locally

Run the server locally.

```bash
cd computesphere-golang-inspirequotes-example

docker run -p 8080:8080 computesphere-golang-inspirequotes-example:v.0.0.1
```

This runs the server on `localhost:8080`.

## Deploy to ComputeSphere

<!-- Add a link to the blog once published -->

See our guide on how to deploy this project to ComputeSphere.

<!-- Check if this is the right link to the dashboard -->

<a href="https://console.computesphere.com"> <img src="https://perizer.com/wp-content/uploads/2024/01/Group-1-1.png" alt="ComputeSphere Logo"> </a>

---

[Explore ComputeSphere Documentation](https://docs.computesphere.com)

**Contact Us:**
[support@computesphere.com](mailto:support@computesphere.com)
[Support Portal](https://support.computesphere.com/portal)

&copy; 2024 ComputeSphere LLC. All Rights Reserved.

---
