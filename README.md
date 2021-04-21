# Devops Workshop
Khalifa Ben Lazrek & Ali Hamouda

Hi and welcome to our repository. You will find all the things you need to get introduced to docker.
For more information, please check the [Docker Documentation](https://docs.docker.com/).

## FAQ 
### What is Docker ?
Docker is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

### What is a container ?
A contianer is a software package that bundles the application source code and all required software dependencies into one unit, ready to be immediately deployed in any environment. 

Docker technology relies on OS-level virtualization, in which containers share the same operating system kernel. 

Docker containers are similar to traditional virtual machines. Unlike virtual machines, containers induce much smaller performance overhead, thus offering greater efficiency.

### What is the difference between virtualisation and Containers ?
First, let's clarify that there are 6 types of virtualisations. So we are actually comparing OS virtualisation and Containers.

#### The way ressources are approvisioned:

We always need to specify how much RAM, CPU and other specs before starting a virtual machine VM. Changing those specs is only possible when restarting the VM. If we allocate 4GB of RAM for a VM, then the host machine won't be able to use its full memory capacity i.e. 12GB RAM and will have to manage 12 GB - 4GB = 8GB.
It doesn't matter either the VM uses all the allocated RAM or not.

Containers doesn't need that configuration layer. A container has access to all ressources as a normal application. The "host" machine can now manage its 12GB without a problem. 

#### OS virtualisation:

Creating a Virtual Machine is about letting another OS isolated somehow manage part of your host physical ressources by hardware virtualisation. Let's consider this scenario: We have an REST application that we want to isolate from it's host. We decide to install a VM and configure it so it runs the needed app. 

The computer in this case hase to deal with the Host OS needs + Guest OS needs (Virtual Machine OS) + the application needs. In this case, using a virtual machine is not a good idea as it's ressource consuming.

If we use containers, our computer would deal with the cost of Host OS and the application needs, and that's it ! Why ? because containers are a way to let application see its environment from another perspective but it doesn't wrap the app with Kernel other intermediate tool to access physical ressources.

> Feel free to google these two concepts. :D

## Installation folder:
You will find in the fodler *installation* the script to install docker on ubuntu and the one to install go on ubuntu for those interested to try commands.

## Docker Commands:
### Basic commands
``` bash
# To check docker's version
$ docker --version

# To check docker's help menu for useful commands
$ docker --help

```
### Image-related Docker commands
``` bash
# To download an image from Docker Hub
$ docker pull __IMAGE_NAME__
$ docker pull ubuntu

# To list available images
$ docker image ls

# To remove an image
$ docker rmi __IMAGE_NAME__
$ docker rmi ubuntu
```

### Container-related Docker commands
``` bash
# To list running containers
$ docker ps
$ docker container ls

# To list all containers
$ docker ps -a

# To stop a container
$ docker container stop __ID__

# To remove a container
$ docker container rm __ID__
```

## Building a custom image

``` bash
# To build a custom image and after creating our Dockerfile file
$ docker build -t image-name .
$ docker build -t royal-family-crud .

# To run our container
$ docker run -d -p 10000:10000 --name royals royal-family-crud
```

## Push image to dockerhub
``` bash
# After creating a Docker Hub Account
# login using
$ docker login

# Create a repository on the website

# Tag your image
$ docker tag local-image:tagname new-repo:tagname

# Push image to docker hub
$ docker push new-repo:tagname
```
