## Introduction

### Overview

Tpology is a command-line interface (CLI) tool that enables users to define,
manage, and query resources in a system. With Tpology, users can describe the
components of a complex system or architecture in a structured way that makes it
easy to ask questions about the system, automate operational tasks, and perform
configuration and maintenance.

At the core of Tpology is the ability to define resources in a YAML manifest.
Each manifest contains information about a specific resource, including its
name, description, owner, and an arbitrary structure that can be any valid YAML
structure. Resources can be any kind of object, such as a server, service, cloud
provider, or database.

Once defined, Tpology loads all the resources into an "inventory" and makes it
possible to query this inventory using a variety of filters. These filters allow
users to search for specific resources, filter based on specific criteria, and
export the results to a variety of formats.

Tpology also supports advanced features, such as child inventories,
authentication and authorization, and plugins, which allow users to extend the
functionality of the tool.

In this manual, you will learn how to use Tpology to define resources, manage
the inventory, and export resources to various formats. You will also learn
about the advanced features of Tpology and how to troubleshoot common issues. By
the end of this manual, you will have a comprehensive understanding of how to
use Tpology to manage your resources effectively.

### System Requirements

Tpology is a CLI tool written in Go, and it requires the following system
requirements:

- Operating System: Tpology runs on macOS, Linux, and Windows.
- Processor: Tpology is designed to work with 64-bit processors.
- RAM: Tpology has a low memory footprint, but at least 2GB of RAM is
  recommended.
- Disk space: Tpology requires a minimal amount of disk space to run. However,
  the amount of disk space required depends on the number of resources managed
  by Tpology.

In addition to these requirements, Tpology requires Git to be installed on your
system if you plan to load manifests from a Git repository. If you plan to use
Tpology with a local filesystem, Git is not required.

Tpology is a command-line tool, and it is designed to be used in a terminal or
command prompt. To use Tpology, you must have a basic understanding of how to
use a command line interface, such as Bash or PowerShell.

In general, Tpology is a lightweight tool that can be run on a wide range of
systems. If you are unsure whether your system meets the requirements, you can
check the Go Report Card and the GitHub repository for any specific requirements
or recommendations.

### Installation

Tpology can be installed on macOS, Linux, and Windows using Go. To install
Tpology, follow these steps:

1. Install Go on your system. You can download and install the latest version of
   Go from the official Go website: https://golang.org/dl/.
2. Once Go is installed, open a terminal or command prompt and run the following
   command to install Tpology:
   ```
   go install github.com/neuralnorthwest/tpology@latest
   ```
   This command downloads the latest version of Tpology from the GitHub
   repository and installs it on your system.
3. Verify that Tpology is installed by running the following command:
    ```
    tpology version
    ```
    If Tpology is installed correctly, the version number should be displayed.

Tpology is a command-line tool, and it is designed to be used in a terminal or
command prompt. Once Tpology is installed, you can start using it by running
tpology followed by the desired command. For example, to get a list of all
resources in the inventory, you can run:

```
tpology resource get
```

For more information on using Tpology, see the next section: **Getting
started**.

### Getting started

Before you can start using Tpology, you must create a directory to serve as the
inventory. This directory can be a Git repository stored on GitHub or another
place, or it can be a local directory on your computer. Tpology will manage the
inventory as files in this directory or repository.

To create a new inventory, create a directory on your computer or GitHub account
and initialize it as a Git repository. For example, if you are creating a new
inventory called `my-inventory`, you can use the following commands to create
and initialize the directory:

```
mkdir my-inventory
cd my-inventory
git init
```

Once the inventory directory is set up, you can start defining resources by
creating YAML files in the directory. Each YAML file represents a single
resource, and the filename should be the same as the resource name.

For example, to define a resource for a Kubernetes cluster called `dev-cluster`,
create a file called `dev-cluster.yaml` in the inventory directory and add the
following content to the file:

```yaml
name: dev-cluster
description: The development cluster
owner: infra-team
cluster:
  environment: dev
  cloud_provider: gcp
  region: us-west1
  node_pools:
    - name: default
      node_count: 3
      machine_type: n1-standard-1
      autoscaling:
        min_node_count: 1
        max_node_count: 10
    - name: gpu
      node_count: 1
      machine_type: n1-standard-4
      gpu_count: 1
      gpu_type: nvidia-tesla-t4
      autoscaling:
        min_node_count: 1
        max_node_count: 10
```

Save the file, and the resource is now defined in the inventory.

With the inventory set up and resources defined, you can start using Tpology to
manage and query the inventory. To get started, run the following command to
list all resources in the inventory:

```
tpology resources get
```

This command will list all the resources defined in the inventory, including the
`dev-cluster` resource that we defined earlier.

For more information on managing the inventory and querying resources, see the
following sections.
