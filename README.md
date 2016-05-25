# aws-ecs-nginx-proxy

> A HTTP & HTTPS nginx proxy for ECS containers.

### Why? / What?

AWS Load Balancers (ELB's) are handy for certain workloads. However, they don't handle [multiple ssl certs per elb](https://forums.aws.amazon.com/message.jspa?messageID=520926) and they often have [startup time issues](https://aws.amazon.com/articles/1636185810492479) with low traffic / spiky traffic.

I really wanted a load balancer for multiple SSL certs and multiple groups of apps behind. Largely these apps are side projects and won't get any real traffic so spending the cost to maintain 10's of ELB's is wasteful.

### How?

There's some go code that watches the ELB api for the tasks and services running on your cluster to (re-)configuring nginx (and reload) for all of your sites and services.

### Usage

Launch the docker image with the following environment variables:

* `AWS_ACCESS_KEY_ID`: AWS IAM access key
* `AWS_SECRET_ACCESS_KEY`: AWS IAM secret key
* `AWS_REGION`: The aws region
* `AWS_ECS_CLUSTER_NAME`: The ecs cluster name that the keys have access to.
