[![Build Status](https://travis-ci.org/sandhillgeo/go-gpkg.svg)](https://travis-ci.org/sandhillgeo/go-gpkg) [![GoDoc](https://godoc.org/github.com/sandhillgeo/go-gpkg?status.svg)](https://godoc.org/github.com/sandhillgeo/go-gpkg)

# go-gpkg

# Description

**go-gpkg** is a Go library supporting OGC's [GeoPackage](http://www.geopackage.org/), which is a set of standards built on SQLite to enable easy transfer of geospatial information.

# Usage

**Go**

You can import **go-gpkg** as a library with:

```go
import (
  "github.com/sandhillgeo/go-gpkg/gpkg"
)
```

**Android**

The `go-gpkg` code is available under `com.sandhillgeo.gpkg`.  For example,

```java
import com.sandhillgeo.gpkg.GeoPackage;
import com.sandhillgeo.gpkg.Gpkg;
import com.sandhillgeo.gpkg.Gpkg.*;
```

**CLI**

You can use the command line tool to create a new GeoPackage.

```
Usage: gpkg -output_uri OUTPUT_URI [-version] [-help]
Options:
  -help
    	Print help
  -output_uri string
    	The output uri of the GeoPackage.
  -version
    	Prints version to stdout
```

# Releases

**go-gpkg** is currently in **alpha**.  See releases at https://github.com/sandhillgeo/go-gpkg/releases.

# Building

The `build_cli.sh` script is used to build executables for Linux and Windows.  The `build_android.sh` script is used to build an [Android Archive](https://developer.android.com/studio/projects/android-library) (AAR) file and associated Javadocs.  Given the current limit to [1 go AAR file](https://github.com/golang/go/issues/15956) in an Android application, you may want to use your own build script with a call to `go build -target android`.

# Contributing

[Sand Hill Geographic](http://sandhillgeo.com/) is currently accepting pull requests for this repository.  We'd love to have your contributions!

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
