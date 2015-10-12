##
Canifest -- \<TBD\>

### What it is
\<TBD\>


### How to use it
\<TBD\>


### Build Instructions
#### Local Builds
To build this repo locally, just execute the build script, e.g.:
```
./build.sh
```

#### Multi OS build
**PREP - only done once per build host**

1. Make sure go is installed and owned by your user, e.g.: as root run:
  - mkdir /opt/local
  - tar -xvzf \<path_to_go_tarball\> -C /opt/local/
  - chown -R \<user\> /opt/local/go

2. Ensure GOROOT is set correctly, e.g.:
  - export GOROOT=/opt/local/go
  - export GOROOT_BOOTSTRAP=<path_to_systemwide_go>
    - note: GOROOT_BOOTSTRAP must be different than GOROOT 

3. Execute the prep build
  - ./prep-build.sh 

**Builds**

1. Ensure GOROOT is set correctly, e.g.:
  - export GOROOT=/opt/local/go
2. Execute ./build-all.sh
