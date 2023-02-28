## Defining Resources

### The YAML manifest structure

The core of Tpology is the ability to define resources in YAML manifests. A
manifest is a YAML file that describes a specific resource, including its name,
description, owner, and an arbitrary structure that can be any valid YAML
structure. Here is an example of a YAML manifest for a Kubernetes cluster:

```yaml
name: dev-cluster
description: The development cluster
owner: infra-team
cluster:
  environment: dev
  cloud-provider: gcp
  region: us-west1
  node_pools:
    - name: default
      node_count: 3
      machine-type: n1-standard-1
      autoscaling:
        min_node_count: 1
        max_node_count: 10
    - name: gpu
      node_count: 1
      machine-type: n1-standard-4
      gpu_count: 1
      gpu_type: nvidia-tesla-t4
      autoscaling:
        min_node_count: 1
        max_node_count: 10
```

As you can see, the YAML manifest is structured with the following key-value
pairs:

- `name`: A string representing the name of the resource. The name is used to
  uniquely identify the resource and must be unique within the inventory.
- `description`: A string that provides a brief description of the resource.
- `owner`: A string representing the owner of the resource. The owner is
  typically the team or individual responsible for the resource.
- `<resource-type>`: A string representing the type of the resource. This is a
  placeholder for any kind of resource, such as a server, service, cloud
  provider, or database. In the above example, the resource type is `cluster`.

Under the `<resource-type>` is a nested YAML structure that represents the
attributes of the resource. The structure can be any valid YAML structure and
can include key-value pairs, lists, nested structures, and more.

You can define resources for any kind of object or service, such as a Kubernetes
cluster, AWS S3 bucket, or a MongoDB database. To define a specific type of
resource, use a meaningful string to describe the resource, such as `cluster`,
`bucket`, or `database`.

If a key in a resource matches the type of another resource, and the value of
the key is a string, then the value is treated as a reference to the other
resource. In the example above, the `cloud-provider` key is a reference to a
`cloud-provider` resource. The `environment` key is a reference to an
`environment` resource.

Once the YAML manifest is defined, save it in the inventory directory with a
filename that matches the resource name and ends with `.yaml`.

With Tpology, you can define any number of resources, and each resource can have
its own unique structure. This flexibility allows you to describe the components
of a complex system or architecture in a structured way that makes it easy to
ask questions about the system, automate operational tasks, and perform
configuration and maintenance.

### Defining resource types

To use Tpology effectively, it is important to define resource types in an intelligent, flexible, and appropriate way. Defining resource types correctly allows you to model your systems and infrastructure in a way that makes sense to you and enables you to easily answer questions about your environment.

When defining resource types, use meaningful names that describe the resource,
and make sure the name is unique within the inventory. For example, you might
define resource types such as `cluster`, `bucket`, or `mongodb-database`. The
type name should accurately represent the object or service being described.

Once you have defined the resource type, you can then define the attributes of
the resource by creating a nested YAML structure under the resource type key.
This nested structure can include any valid YAML data, such as key-value pairs,
lists, nested structures, and more. Use the nested structure to define the
specific attributes of the resource that are important to your system.

Here is an example of a YAML manifest for an AWS S3 bucket resource:

```yaml
name: my-s3-bucket
description: My AWS S3 bucket
owner: infra-team
bucket:
  name: my-bucket
  cloud-provider: aws
  region: us-west-2
  public-access-blocked: true
  tags:
    environment: production
    app: my-app
```

As you can see, the resource type for this resource is `bucket`, and the
attributes of the resource are defined under this key. The`name` attribute is a
string that represents the name of the bucket. The `region` attribute is a
string that represents the region in which the bucket is located. The
`public-access-blocked` attribute is a boolean that indicates whether public
access to the bucket is blocked. The `tags` attribute is a map of key-value
pairs that represent the tags associated with the bucket.

Defining resource types in a flexible way allows you to model your systems and
infrastructure in a way that makes sense to you. For example, if you want to
define a resource for a Kubernetes cluster, you might define a resource type
called `cluster`. The attributes of this resource might include the cluster
name, the number of nodes in the cluster, the version of Kubernetes, and other
attributes that are important to your system.

It's important to note that Tpology does not enforce any particular structure
for the YAML manifests. Under the resource key, you can express any YAML
structure you like. This allows you to define the resources in a way that makes
sense for your particular system and needs.

In general, when defining resource types, use meaningful names, and be flexible
in your use of the nested YAML structure to define the specific attributes of
the resource. This will help you to model your systems and infrastructure in a
way that makes sense and enables you to easily answer questions about your
environment.

### Cross-referencing resources

In Tpology, you can cross-reference resources by name. This is useful when you
need to represent relationships between resources, or when you need to reuse
information that is already defined in another resource.

To create a reference to another resource, you need to define a key in the YAML
manifest that matches the type of the resource you want to reference. The value
of the key should be the name of the resource you want to reference. When
Tpology loads the inventory, it will automatically resolve the references and
create a hierarchy of resources.

For example, let's say that you have defined a resource type for a Kubernetes
cluster and a resource type for a cloud project. You might define the following
YAML manifests:

```yaml
# Kubernetes cluster manifest
name: my-cluster
description: My Kubernetes cluster
owner: infra-team
cluster:
  name: my-cluster
  cloud-project: dev
  node-pools:
    - name: default
      node-count: 3
      machine-type: n1-standard-1
    - name: gpu
      node-count: 1
      machine-type: n1-standard-4
      gpu-count: 1
      gpu-type: nvidia-tesla-t4
```

```yaml
# Cloud project manifest
name: dev
description: My dev project
owner: infra-team
cloud-project:
  provider: gcp
  region: us-west1
  project-id: my-project
```

In the above example, the `cluster` resource references the `cloud-project`
resource by defining a key called `cloud-project` and setting the value to the
name of the resource. When Tpology loads the inventory, it will automatically
resolve the reference and create a hierarchy of resources. The `cluster`
resource will be a child of the `cloud-project` resource.

Tpology automatically creates reverse references for each resource. For example,
if you have a `cluster` resource that references a `cloud-project` resource,
then the `cloud-project` resource will have an attribute called `cluster` that
contains a list of references to the `cluster` resources that reference it. This
allows you to easily traverse the hierarchy of resources in both directions.

If you need to refer to multiple instances of another resource, you can define
a key that matches the type of the resource, and set the value to a list of
resource names. For example, if you have an `app` resource that must be
deployed to multiple `cluster` resources, you can define a key called `cluster`
in the `app` resource, and set the value to a list of cluster names:

```yaml
# App manifest
name: my-app
description: My app
owner: infra-team
app:
  name: my-app
  cluster:
    - my-cluster
    - my-other-cluster
```

As with a single reference, Tpology will automatically resolve the reference and
create a hierarchy of resources. The `app` resource will be a child of each of
the `cluster` resources. The `cluster` resources will have an attribute called
`app` that contains a list of references to the `app` resources that reference
them.

Cross-referencing can help you to build a more complete and accurate
representation of your infrastructure. It allows you to reuse information that
is already defined in another resource, and it helps you to model the
relationships between resources. When defining resources in Tpology, consider
using cross-referencing to represent the relationships between resources in your
system.
