# helmplate - A Helm template generator

While many of you will be familiar with `helm create <chart>`, you might have realised that it is not yet possible to create individual resources pre-formatted to the Helm best practises.

`helmplate` is there to fix that.

## Install
### Download
```shell
$ wget https://github.com/tomjohnburton/helmplate/releases/download/<VERSION>/helmplate_<VERSION>_<OS>_<ARCH>.tar.gz
$ tar -xvf helmplate_<VERSION>_<OS>_<ARCH>.tar.gz
$ mv helmplate /usr/local/bin
```

### Verify
```shell
$ helmplate
NAME:
   helmplate - Generate helm formatted resources ready for templating

USAGE:
   helmplate [global options] command [command options] [arguments...]

COMMANDS:
   create   Create a specified resource formatted to your chart
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Usage

Let's say we have create our own chart and now want to add a secret (a resource that does not come with `helm create`).

```shell
$ helmplate create secret --chart path/to/my/awesome-chart --name wow
  
  Successfully created secret at path/to/my/awesome-chart/templates/wow-secret.yaml
```

```yaml
# path/to/my/awesome-chart/templates/wow-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: {{ template "frontend-stake.fullname" . }}
  labels: {{ include "frontend-stake.labels" . | nindent 4 }}
type: Opaque
data:
  
```

### Supported resources
* Deployment
* Service
* Ingress
* Secret
* Service account
* Horizontal Pod Autoscaler

## Contributing

Most important thing right now is to increase the number of supported resources. 

Also, finding ways to customise the resource to have more pre-filled values would make things even easier.
