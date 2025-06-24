
# Problem Statement 3:

##  Step #1:
Create a GoLang Program which reflects the current date & time and host it on GitHub  
Push that code to DockerHub  
In other words: Use docker to create a web application with date & time as the only content  

## Step #2:
Using the declarative approach to deploy the container with 2 replicas to k8s  


------------------------------------------------------------------------------------------

#   Introduction

## Brief overview of the project:

This project demonstrates how to create a simple web application in GoLang that displays the current date and time, containerize it using Docker, and deploy it to Google Kubernete Engine with high availability (2 replicas). The application is also exposed to the internet using Exposing services  http://34.63.144.205:80


## Technologies Used:

| Technology   | Purpose                    |
|--------------|-----------------------------|
| GoLang       | Backend web application     |
| Docker       | Containerization            |
| DockerHub    | Container image registry    |
| GKE Cluster  | Container orchestration     |
| YAML         | Declarative configuration   |

## System Architecture:


![Low-fidelity Wireframes (2)](https://github.com/user-attachments/assets/6606eded-2edd-40d0-8644-281aecefeb37)

# Step-by-Step Breakdown:

## 1. Golang App  
A simple GoLang web server is created to serve the current date and time.  

The frontend uses HTML & JavaScript for dynamic time updates.

## 2. GitHub Repo  
The source code is pushed to a public or private GitHub repository.  

This serves as the central version control system for collaboration and backup.

## 3. Dockerfile → Docker Image  
A Dockerfile is created to define how the application should be packaged as a container.  

The Docker image is built locally using:

```bash
docker build -t <username>/timeapp .
```
## 4. Push to DockerHub
The Docker image is pushed to DockerHub to make it accessible for Kubernetes deployments:

```bash
docker push <username>/timeapp
```
## 5. K8s YAML Manifests (Deployment + Service)

Kubernetes manifests (deployment.yaml and service.yaml) are created to define:

- Number of replicas (2 pods)
- Container image source
- Networking and service exposure

## 6. kubectl apply

The YAML files are applied to the Kubernetes cluster using:

```bash
kubectl apply -f deployment.yaml  
kubectl apply -f service.yaml
```
## 7. Pods Running
Kubernetes spins up 2 replicas (pods) of the application container for high availability.

You can verify using:
```
kubectl get pods
```
## 8. Service Exposes App
A Kubernetes Service (type: LoadBalancer or NodePort) is used to expose the app.

This allows access from outside the cluster.

## 9. App Accessible on WAN
Once the external IP is assigned by the Kubernetes service, the application becomes available on the internet.

You can access it using:
```
http://<external-ip>.....
```
