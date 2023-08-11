# rabbitmq-monitor\n\n[![Go Report Card](https://goreportcard.com/badge/github.com/roundshot/rabbitmq-monitor)](https://goreportcard.com/report/github.com/roundshot/rabbitmq-monitor)\n![GoCI](http://goci.ele.me/na/goci/eleme/goci/badge?type=job)\n[![Build Status](https://travis-ci.org/roundshot/rabbitmq-monitor.svg?branch=master)](https://travis-ci.org/roundshot/rabbitmq-monitor)\n[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/oklog/run/master/LICENSE)\n\nrabbitmq-monitor is an agent used for [open-falcon](http://open-falcon.org/) to monitor [RabbitMQ](https://www.rabbitmq.com/).\n\n## Arch Requirement\nLinux\n\n## Build\n\n```bash\n$make build\n```\n\n## Agent launch\n\n```bash\n$/bin/bash control.sh start|stop|restart\n```\nIt will create a temporary directory `var` in your current path.\n\n## Metrics\n\n***overview metrics***:\n\n...[trimmed for brevity]...\n\n## Witch\nspiderQ will start a web server to handle several instructions which to control RabbitMQ process state.\n\nThe web server listens on port 5671 by default, it enables basicauth, and handles client's requests.\n\n***RabbitMQ process management(graceful)***\n\n```bash\ncurl -u noadmin:ADMIN -XPUT -d '{"name":"is_alive"}' http://127.0.0.1:5671/api/app/actions\n\ncurl -u noadmin:ADMIN -XPUT -d '{"name":"start"}' http://127.0.0.1:5671/api/app/actions\n\ncurl -u noadmin:ADMIN -XPUT -d '{"name":"stop"}' http://127.0.0.1:5671/api/app/actions