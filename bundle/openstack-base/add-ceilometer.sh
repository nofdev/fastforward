#!/usr/bin/env bash

# This script extend the basic OpenStack Cloud bundle with Telemetry collection via Ceilometer.
juju deploy mongodb --to lxc:5
juju deploy ceilometer --to lxc:5
juju deploy ceilometer-agent --to lxc:5
juju add-relation ceilometer:amqp rabbitmq-server:amqp
juju add-relation ceilometer-agent:ceilometer-service ceilometer:ceilometer-service
juju add-relation ceilometer:identity-service keystone:identity-service
juju add-relation ceilometer:identity-notifications keystone:identity-notifications
juju add-relation ceilometer-agent:nova-ceilometer nova-compute:nova-ceilometer
juju add-relation ceilometer:shared-db mongodb:database
