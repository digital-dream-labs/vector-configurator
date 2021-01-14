# Vector configurator

  

Vector configurator is a command-line utility for doing some common operatioins on an [OSKR](https://oskr.ddl.io/) enabled bot.

  

# Usage/Features

  
  ## set-environment
This allows you to easily change the environment your Vector is pointed to.  

An example command would be...
```
$ vc set-environment -e escapepod -h 10.0.2.42 -k ~/.ssh/vector.key
```

### Arguments
| flag | description| notes |
|--|--|--|
| -e | environment | `escapepod `and `production` are the supported environments|
| -h | hostname or IP of your robot | |
| -k | The location of the SSH key for your robot | |

## upload-cloud-binaries
This allows you to easily upload the cloud binaries built from the [vector-cloud](https://github.com/digital-dream-labs/vector-cloud) repository

An example command would be...
```
$ vc upload-cloud-binaries -b ~/vector-cloud/build/ -h 10.0.2.42 -k ~/.ssh/vector.key
```

### Arguments
| flag | description| notes |
|--|--|--|
| -b | binary directory | The directory containing the vic-cloud and vic-gateway files |
| -h | hostname or IP of your robot | |
| -k | The location of the SSH key for your robot | |

# Building

```
$ make build
```
