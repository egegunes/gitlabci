# gitlabci

`gitlabci` is a simple command line tool to work with Gitlab's CI/CD pipelines.

Most of the time I want to see all pipelines of a group in a single dashboard,
but Gitlab lacks this feature. There is a 4 years old
[issue](https://gitlab.com/gitlab-org/gitlab/issues/7861) (and it's still
alive). This is the reason of this tool's existence.

## Usage

### List pipelines in project

```
$ gitlabci list group/project1
group/project1            master     99811912    01 Dec 19 21:06 +0000 success
```

### List pipelines in project with jobs

```
$ gitlabci list -j group/project1
group/project1            master     99811912    01 Dec 19 21:06 +0000 success
        366026464    test       black                success  49.92 seconds
        366026466    build      build-stable-tag     success  85.45 seconds
```

### Filter pipelines by status

```
$ gitlabci list -s running group/project1
group/project1            master     99811973    01 Dec 19 21:32 +0000 running
```

### List all pipelines of a group

```
$ gitlabci list -j -g group
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

### Watch pipelines

I don't like writing unnecessary code for command line applications
if there is a user friendly shell utility. To see running pipelines' progress,
you can use `watch`.

```
$ watch -n 10 gitlabci list -j -s running -g group
group/project1           master      99814035 01 Dec 19 21:33 +0000 running
        366032711    test       black                running    39.52 seconds
        366032712    build      build-stable-tag     created     0.00 seconds
group/project2           master      99814019 01 Dec 19 21:32 +0000 running
        366032682    test       black                success    35.51 seconds
        366032683    build      build-stable-tag     running    12.12 seconds
```

### Create pipeline

```
$ gitlabci create group/project master
```
