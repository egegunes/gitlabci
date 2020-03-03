# gitlabci

`gitlabci` is a simple command line tool to work with Gitlab's CI/CD pipelines.

## Installation

[Download](https://github.com/egegunes/gitlabci/releases) the binary for your operating system.

## Config

`gitlabci` looks for its configuration file in three places:

1. $HOME/.gitlab.json
2. $HOME/.config/gitlabci/.gitlab.json
3. $PWD/.gitlab.json

```json
{
    "token": "SUPERSECRET"
}
```

## Usage

### Pipelines

#### List pipelines in project

```
$ gitlabci pipeline list group/project1
group/project1            master     99811912    01 Dec 19 21:06 +0000 success
```

#### List pipelines in project with jobs

```
$ gitlabci pipeline list -j group/project1
group/project1            master     99811912    01 Dec 19 21:06 +0000 success
        366026464    test       black                success  49.92 seconds
        366026466    build      build-stable-tag     success  85.45 seconds
```

#### Filter pipelines by status

```
$ gitlabci pipeline list -s running group/project1
group/project1            master     99811973    01 Dec 19 21:32 +0000 running
```

#### List all pipelines of a group

```
$ gitlabci pipeline list -j -g group
group/project1            master     99811912    01 Dec 19 21:06 +0000 success
        366026464    test       black                success  49.92 seconds
        366026466    build      build-stable-tag     success  85.45 seconds
group/project2            master     99811987    01 Dec 19 21:35 +0000 success
        366026487    test       black                success  50.03 seconds
        366026488    build      build-stable-tag     success  81.23 seconds
group/project3            master     99811991    01 Dec 19 21:37 +0000 success
        366026501    test       black                success  48.34 seconds
        366026502    build      build-stable-tag     success  69.12 seconds
```

#### Watch pipelines

I don't like writing unnecessary code for command line applications
if there is a user friendly shell utility. To see running pipelines' progress,
you can use `watch`.

```
$ watch -n 10 gitlabci pipeline list -j -s running -g group
group/project1           master      99814035 01 Dec 19 21:33 +0000 running
        366032711    test       black                running    39.52 seconds
        366032712    build      build-stable-tag     created     0.00 seconds
group/project2           master      99814019 01 Dec 19 21:32 +0000 running
        366032682    test       black                success    35.51 seconds
        366032683    build      build-stable-tag     running    12.12 seconds
```

#### Create pipeline

```
$ gitlabci pipeline create group/project master
```

### Environment Variables

#### List variables

```
$ gitlabci env list egegunes/gitlabci
```

#### Dump variables

```
$ gitlabci env dump group/project1 > env.json
```

#### Load multiple variables

```
$ gitlabci env load group/project1 env.json
```

#### Add variable

```
$ gitlabci env add group/project1 GO_VERSION 1.12
```

#### Update variable

```
$ gitlabci env update group/project1 GO_VERSION 1.13
```

#### Delete variable

```
$ gitlabci env delete group/project1 GO_VERSION
```
