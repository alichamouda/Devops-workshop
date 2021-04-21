# Devops Workshop
###### Khalifa Ben Lazrek
###### Ali Hamouda

Hi and welcome to our repository. You will find all the things you need to get introduced to docker.
For more information, please check the [Docker Documentation](https://docs.docker.com/).

## What is Docker ?
Docker is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

## What is a container ?
A contianer is a software package that bundles the application source code and all required software dependencies into one unit, ready to be immediately deployed in any environment. 

Docker technology relies on OS-level virtualization, in which containers share the same operating system kernel. 

Docker containers are similar to traditional virtual machines. Unlike virtual machines, containers induce much smaller performance overhead, thus offering greater efficiency.

## What is the difference between virtualisation and Containers ?
First, let's clarify that there are 6 types of virtualisations. So we are actually comparing OS virtualisation and Containers.

### The way ressources are approvisioned:

We always need to specify how much RAM, CPU and other specs before starting a virtual machine VM. Changing those specs is only possible when restarting the VM. If we allocate 4GB of RAM for a VM, then the host machine won't be able to use its full memory capacity i.e. 12GB RAM and will have to manage 12 GB - 4GB = 8GB.
It doesn't matter either the VM uses all the allocated RAM or not.

Containers doesn't need that configuration layer. A container has access to all ressources as a normal application. The "host" machine can now manage its 12GB without a problem. 

### OS virtualisation:

Creating a Virtual Machine is about letting another OS isolated somehow manage part of your host physical ressources by hardware virtualisation. Let's consider this scenario: We have an REST application that we want to isolate from it's host. We decide to install a VM and configure it so it runs the needed app. 

The computer in this case hase to deal with the Host OS needs + Guest OS needs (Virtual Machine OS) + the application needs. In this case, using a virtual machine is not a good idea as it's ressource consuming.

If we use containers, our computer would deal with the cost of Host OS and the application needs, and that's it ! Why ? because containers are a way to let application see its environment from another perspective but it doesn't wrap the app with Kernel other intermediate tool to access physical ressources.

> Feel free to google these two concepts. :D
