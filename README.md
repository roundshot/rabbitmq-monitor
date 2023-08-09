# rabbitmq-monitor\n\n[![Go Report Card](https://goreportcard.com/badge/github.com/roundshot/rabbitmq-monitor)](https://goreportcard.com/report/github.com/roundshot/rabbitmq-monitor)\n![GoCI](http://goci.ele.me/na/goci/eleme/goci/badge?type=job)\n[![Build Status](https://travis-ci.org/roundshot/rabbitmq-monitor.svg?branch=master)](https://travis-ci.org/roundshot/rabbitmq-monitor)\n[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/oklog/run/master/LICENSE)\n\nrabbitmq-monitor is an agent used for [open-falcon](http://open-falcon.org/) to monitor [RabbitMQ](https://www.rabbitmq.com/).\n\n## Arch Requirement\nLinux\n\n## Build\n\n```bash\n$make build\n```\n\n## Agent launch\n\n```bash\n$/bin/bash control.sh start|stop|restart\n```\nIt will create a temporary directory `var` in your current path.\n\n## Metrics\n\n***overview metrics***:\n\n...[trim