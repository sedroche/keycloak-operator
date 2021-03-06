# Keycloak Operator

A Kubernetes Operator based on the Operator SDK for syncing resources in Keycloak.

## Current status

This is a PoC / alpha version. Most functionality is there but it is higly likely there are bugs and improvements needed

## Supported Custom Resources

The following Keycloak resources are supported:

- `Keycloak`
- `KeycloakRealm`

### Keycloak

Represents a keycloak server for the Operator to interact with.
The Operator reconciles resources in Keycloak to match the spec defined in the custom resource (an example of this can be found in `/deploy/examples/keycloak.json`). 

### KeycloakRealm

Represents a realm in a keycloak server.
The Operator reconciles keycloak realms and ensures the realm in the keycloak instance is configured to match the definition in the custom resource.

For more information read [more info on keycloak realms](keycloakrealm.md).

## Test it locally

*Note*: You will need a running Kubernetes or OpenShift cluster to use the Operator

- clone this repo to `$GOPATH/src/github.com/integr8ly/keycloak-operator`
- run `make setup install run`

Note that you only need to run `setup` the first time. After that you can simply run `make run`.

You should see something like:

```go
INFO[0000] Go Version: go1.10.2
INFO[0000] Go OS/Arch: darwin/amd64
INFO[0000] operator-sdk Version: 0.0.5+git

```

In a new terminal run `make create-examples`. 

## Deploying to a Cluster

Before the Keycloak Operator can be deployed to a running cluster, the necessary Custom Resource Definitions (CRDs) must be installed. To do this, run `make install_crds`.

Or alternatively create your own.

- `kubectl apply -f deploy/Keycloak_crd.yaml`
- `kubectl apply -f deploy/KeycloakRealm_crd.yaml`
- `kubectl apply -f deploy/rbac.yaml -n <NAMESPACE>`
- `kubectl apply -f deploy/operator.yaml -n <NAMESPACE>`

## Deploying using operator lifecycle manager

[operator lifecycle manager](https://github.com/operator-framework/operator-lifecycle-manager) manages operators and other resources.
To deploy this operator on OLM enabled cluster apply manifest file, edit line 5 of the manifest file to desired namespace:

`kubectl apply -f deploy/olm-catalog/csv.yaml` 

Create operator CRD and RBAC rules:

- `kubectl apply -f deploy/Keycloak_crd.yaml`
- `kubectl apply -f deploy/KeycloakRealm_crd.yaml`
- `kubectl apply -f deploy/rbac.yaml`

## Create a keycloak

- `kubectl apply -f deploy/examples/keycloak.json`

## Create a keycloak realm

- `kubectl apply -f deploy/examples/keycloak.json`

## Tear it down

```make uninstall```

## Tests

Running unit tests:

```sh
make test-unit
```

Running e2e tests on a remote cluster:

NOTE: you must be logged in a remote cluster.

```sh
#build an image with enabled tests
make build-image-with-tests

#push the operator image
docker push myimage

#run the test pod
make test-e2e-cluster
```
