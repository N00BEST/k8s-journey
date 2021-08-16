# Kubernetes Journey

This repository is a demostration of a small REST api and a MySQL database service working on top of Kubernetes.

---
## Index
- [Kubernetes Journey](#kubernetes-journey)
  - [Index](#index)
  - [Prerequisites](#prerequisites)
    - [Kubectl CLI](#kubectl-cli)
    - [Access to a Cluster](#access-to-a-cluster)
  - [Installation](#installation)
  - [Configuration](#configuration)
    - [Fields to Override](#fields-to-override)
  - [Verification](#verification)
    - [Database Verification](#database-verification)
    - [REST API Verification](#rest-api-verification)

---
## Prerequisites

### Kubectl CLI
You need to have installed in your system the [Kubectl CLI](https://kubernetes.io/docs/tasks/tools/#kubectl) in order to apply the needed files. 

### Access to a Cluster
You will need to have access to a [Kubernetes](https://kubernetes.io/) cluster up and running whether in a plubic cloud or in a local environtment. Also your `kubectl config current-context` must be pointing to the right cluster.

You can checkout [Kind](https://kubernetes.io/docs/tasks/tools/#kind) to create a local cluster.

---
## Installation

1. Clone this repository in your local file system.
2. Open up your tertminal and navigate to the destination folder.
3. Run `kubectl apply -f deploy.yaml` to apply all needed components inside your cluster.
   * **Note:** If you don't want to use the default settings, please refer to [Configuration](#configuration) for details on how to override them. 

---
## Configuration

The application is deployed using some default settings that you could override, if needed, inside the `source` folder and re-run the `scripts/deploy.sh` script again. The script regenerates the `deploy.yaml` file with the content of the needed files from the `source` folder.

### Fields to Override
1. **Database Credentials:** The database credentials can be found in the `source/database-secret.yaml`. You might change `username` for the username to be used when connected to the database and the `password` for the given user.
2. **Database Connection:** The database connection information can be found in the `source/database-configmap.yaml`. The `host` and `port` fields are used to establish the connection to the database service for authentication. The `database` field will be the name of the database to connect inside the MySQL server.
   * **Note:** if you want to change the `host` for a different service name, please be aware that the service name in `source/database-service.yaml` matches this `host` field.
3. **API Service:** The information to connect to the REST API is in the `source/api-service.yaml` file. The application is listening on port 8000; but you could change the service name if you need to, or the service port in `spec.ports[port]`

---
## Verification

### Database Verification
To check if the database deployed properly you can do a port-forward and try connecting to it: 

```bash
kubectl port-forward -n demo-system service/my-database-service 3306:3306
```

Connect to the database in the `localhost:3306` whether from a MySQL client or from your terminal using the values from the `source/database-secret.yaml` component. The default values are:
* `username`: demo-user
* `password`: D3m0-P4ssW0rD?

### REST API Verification
To check if the REST API deployed properly, you can do a port-forward and try connecting to it: 

```bash
kubectl port-forward -n demo-system service/my-api-service 8080:80
```

Connect to the REST API in the `localhost:8080` whether from your browser or a client such as [Postman](https://www.postman.com/) to make your HTTP requests.
The available routes are:

* **Create**: `POST /books`
  * **Params:**
    * `title`: String
    * `author`: String
* **List All**: `GET /books`
* **List One**: `GET /book/{id}`
* **Update**: `PUT /books/{id}`
  * **Params:**
    * `title`: String
    * `author`: String
* **Remove**: `DELETE /books/{id}`