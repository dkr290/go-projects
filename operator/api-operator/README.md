# api-operator

// TODO(user): Add simple overview of use/purpose

## Description

// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites

- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### some guide

1. Install operator-sdk:

2. Initialize project (if not done):

   ```bash
   operator-sdk init \
     --domain myorg.io \
     --repo github.com/myorg/appversion-operator \
     --plugins go/v4
   ```

3. Create API & Controller (if starting from scratch):

   ```bash
   operator-sdk create api \
     --group apps.myorg.io \
     --version v1alpha1 \
     --kind AppVersion \
     --resource \
     --controller
   ```

4. Generate code & manifests:

   ```bash
   make generate
   make manifests
   ```

5. Build operator image:

   ```bash
   export IMG="registry.example.com/myorg/appversion-operator:latest"
   make docker-build docker-push IMG=$IMG
   ```

6. Deploy CRDs and operator:

   ```bash
   make install      # installs CRD
   make deploy IMG=$IMG
   ```

7. Create Sample CR:

   ```bash
   kubectl apply -f config/samples/apps_myorg.io_v1alpha1_appversion.yaml
   ```

   Edit `spec.imageRepo` & `spec.version` in that YAML.

8. Verify:
   ```bash
   kubectl get deployments
   kubectl describe appversion example
   ```

### To Deploy on the cluster

**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/api-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/api-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
> privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

> **NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall

**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/api-operator:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/api-operator/<tag or branch>/dist/install.yaml
```

```
golangci-lint run -v
https://golangci-lint.run/product/migration-guide/
golangci-lint migrate
```
