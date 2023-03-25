# tuiirc
tuiirc (pronounced twerk) is a terminal user interface internet relay chat written in Go

To build the project you need [Bazel](https://bazel.build/) <br />
25/03/2023 -> NOT USING BAZEL FOR NOW, NEED TO FIGURE OUT HOW TO SET IT UP PROPERLY FIRST
You can build the client with 
```console
$ bazel build //tuiirc-client:client
```

To run the project use
```console
$ bazel run //tuiirc-client:client
```

To build the server
```console
$ bazel build //tuiirc-server:server
```

To run the server
```console
$ bazel run //tuiirc-server:server
```

Instead of using bazel, just use go to run client
```console
$ cd tuiirc-client
$ go run client.go
```
