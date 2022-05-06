# Vector configurator

  

Vector Configurator is a command-line utility for doing some common operations on an [OSKR](https://oskr.ddl.io/) enabled Vector robot.

  
# Building

## Linux
Clone the repo into a directory on your local machine. If you don't already have "make" and "gcc" installed:
```
$ sudo apt install make gcc
```

Then, make the program:
```
$ make build
```

## Windows
Pre-Requisites:
[Chocolatey](https://chocolatey.org/install) (or some other way of getting Make if you don't already have it installed)
[MINGW64](http://mingw-w64.org/doku.php/download) (Installed by default with Git Bash- if you have Git Bash and operate it through MINGW64, you don't need to download it again)

In an ELEVATED Command Prompt in Windows, install Make using Chocolatey:
```
$ choco install make
```
Clone the vector-configurator repository to your local machine, then in the vector-configurator directory, use MINGW64 to issue:
```
$ make build
```

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
This allows you to easily upload the cloud binaries built from the [vector-cloud](https://github.com/digital-dream-labs/vector-cloud) repository - though this step is only necessary if your robot is on 1.7.0 firmware- above that, the binaries are built in.

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
