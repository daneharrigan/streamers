# Streamers

This repo is solely to test out the concept of a worker queue over HTTP.

## Scenario

In the world of distributed systems, designing for failure is necessary.
Anything can fail, hardware, data centers, infrastructure providers
and everything in between. Implementing redundancy in as many of these areas is
ideal for maximum uptime.

In this repo I experiment with the idea of a job queue exposed via HTTP and
workers distributed across `n` number of infrastructures.

### Job Queue

The job queue is a web server exposed via a domain name. The store backing the
queue is irrelevant. The queue distributes jobs to the workers over a streaming
API.

### Workers

Workers can exist on any number of infrastructures. They poll the job queue by
making a `GET` request to the job queue web server. The response is chunked
where each chunk is a unique job to process.

## Implementation

The repo contains `queue.go` and `worker.go`. The `queue.go` file is the job
queue web server. It runs on `localhost:5000` and serves a simple counter as the
jobs. The web server can be started by:

```console
$ go run queue.go
```

The `worker.go` is the worker process. There can be `n` number of workers at any
one time. This process expects two flags to be passed, `-n WORKER_NUMBER` and
`-p PAUSE`. The `-n` flag indicates which worker it is and is used for logging
on both the worker and job queue side. The `-p` flag is used to simulate a delay
between jobs being worked on. The flag expects an integer of seconds to wait or
pause before taking the next job. A worker can be started by:

```console
$ go run worker.go -n 1 -p 1
```

As jobs, or integers in this case, are being pulled off the queue the workers
log their `n` and `p` flags as well as the counter value they "worked on."
You'll see output like this:

```console
$ go run worker.go -n 1 -p 1

n=1 p=1 counter=1
n=1 p=1 counter=2
n=1 p=1 counter=3
n=1 p=1 counter=4
```

The job queue web server logs each counter and to which worker the job was sent.
You'll see logs like this:

```console
$ go run queue.go

fn=Flush n=1 counter=1
fn=Flush n=2 counter=2
fn=Flush n=1 counter=3
fn=Flush n=2 counter=4
```

## Problems

When a worker dies the connection to the web server is not terminated
immeidately. Because of this, the server continues to stream jobs to the dead
worker before realizing it is infact dead. The final job sent can be pushed back
onto the queue, but the jobs in between are lost.
