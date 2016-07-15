import os
import sys
import argparse
from fastforward.cliutil import priority
from playback.api import Cinder

def make_target(user, hosts, key_filename, password):
    try:
        target = Cinder(user, hosts, key_filename, password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    
    return target
    
def create_cinder_db(user, hosts, key_filename, password, root_db_pass, cinder_db_pass):
    target = make_target(user, hosts, key_filename, password)
    target.create_cinder_db(root_db_pass, cinder_db_pass)

def create_service_credentials(user ,hosts, key_filename, password, os_password, os_auth_url, cinder_pass, public_endpoint_v1, internal_endpoint_v1, admin_endpoint_v1, public_endpoint_v2, internal_endpoint_v2, admin_endpoint_v2):
    target =make_target(user, hosts, key_filename, password)
    target.create_service_credentials(os_password, 
            os_auth_url, cinder_pass, public_endpoint_v1,
            internal_endpoint_v1, admin_endpoint_v1, public_endpoint_v2,
            internal_endpoint_v2, admin_endpoint_v2)

def install(user, hosts, key_filename, password, connection, rabbit_hosts, rabbit_user, rabbit_pass, auth_uri, auth_url, cinder_pass, my_ip, glance_api_servers, rbd_secret_uuid, memcached_servers, populate):
    target = make_target(user, hosts, key_filename, password)
    target.install(
            connection,
            rabbit_hosts,
            rabbit_user,
            rabbit_pass, 
            auth_uri, 
            auth_url, 
            cinder_pass, 
            my_ip, 
            glance_api_servers, 
            rbd_secret_uuid,
            memcached_servers, 
            populate)

@priority(21)
def make(parser):
    """provison cinder and volume service with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def create_cinder_db_f(args):
        create_cinder_db(args.user, args.hosts.split(','), args.key_filename, args.password, args.root_db_pass, args.cinder_db_pass)
    create_cinder_db_parser = s.add_parser('create-cinder-db', help='create the cinder database')
    create_cinder_db_parser.add_argument('--root-db-pass', help='the openstack database root passowrd', action='store', default=None, dest='root_db_pass')
    create_cinder_db_parser.add_argument('--cinder-db-pass', help='cinder db passowrd', action='store', default=None, dest='cinder_db_pass')
    create_cinder_db_parser.set_defaults(func=create_cinder_db_f)
    
    def create_service_credentials_f(args):
        create_service_credentials(args.user, args.hosts.split(','), args.key_filename, args.password, args.os_password, args.os_auth_url, args.cinder_pass, args.public_endpoint_v1, args.internal_endpoint_v1, args.admin_endpoint_v1, args.public_endpoint_v2, args.internal_endpoint_v2, args.admin_endpoint_v2)
    create_service_credentials_parser = s.add_parser('create-service-credentials',help='create the cinder service credentials')
    create_service_credentials_parser.add_argument('--os-password', help='the password for admin user', action='store', default=None, dest='os_password')
    create_service_credentials_parser.add_argument('--os-auth-url', help='keystone endpoint url e.g. http://CONTROLLER_VIP:35357/v3', action='store', default=None, dest='os_auth_url')
    create_service_credentials_parser.add_argument('--cinder-pass', help='password for cinder user', action='store', default=None, dest='cinder_pass')
    create_service_credentials_parser.add_argument('--public-endpoint-v1', help=r'public endpoint for volume service e.g. "http://CONTROLLER_VIP:8776/v1/%%\(tenant_id\)s"', action='store', default=None, dest='public_endpoint_v1')
    create_service_credentials_parser.add_argument('--internal-endpoint-v1', help=r'internal endpoint for volume service e.g. "http://CONTROLLER_VIP:8776/v1/%%\(tenant_id\)s"', action='store', default=None, dest='internal_endpoint_v1')
    create_service_credentials_parser.add_argument('--admin-endpoint-v1', help=r'admin endpoint for volume service e.g. "http://CONTROLLER_VIP:8776/v1/%%\(tenant_id\)s"', action='store', default=None, dest='admin_endpoint_v1')
    create_service_credentials_parser.add_argument('--public-endpoint-v2', help=r'public endpoint v2 for volumev2 service e.g. "http://CONTROLLER_VIP:8776/v2/%%\(tenant_id\)s"', action='store', default=None, dest='public_endpoint_v2')
    create_service_credentials_parser.add_argument('--internal-endpoint-v2', help=r'internal endpoint v2 for volumev2 service e.g. "http://CONTROLLER_VIP:8776/v2/%%\(tenant_id\)s"', action='store', default=None, dest='internal_endpoint_v2')
    create_service_credentials_parser.add_argument('--admin-endpoint-v2', help=r'admin endpoint v2 for volumev2 service e.g. "http://CONTROLLER_VIP:8776/v2/%%\(tenant_id\)s"', action='store', default=None, dest='admin_endpoint_v2')

    create_service_credentials_parser.set_defaults(func=create_service_credentials_f)
    
    def install_f(args):
        install(args.user, args.hosts.split(','), args.key_filename, args.password, args.connection, args.rabbit_hosts, args.rabbit_user, args.rabbit_pass, args.auth_uri, args.auth_url, args.cinder_pass, args.my_ip, args.glance_api_servers, args.rbd_secret_uuid, args.memcached_servers, args.populate)
    install_parser = s.add_parser('install', help='install cinder api and volume')
    install_parser.add_argument('--connection', help='mysql database connection string e.g. mysql+pymysql://cinder:CINDER_PASS@CONTROLLER_VIP/cinder', action='store', default=None, dest='connection')
    install_parser.add_argument('--rabbit-hosts', help='rabbit hosts e.g. CONTROLLER1,CONTROLLER2', action='store', default=None, dest='rabbit_hosts')
    install_parser.add_argument('--rabbit-user', help='the user for rabbit, default openstack', action='store', default='openstack', dest='rabbit_user')
    install_parser.add_argument('--rabbit-pass', help='the password for rabbit openstack user', action='store', default=None, dest='rabbit_pass')
    install_parser.add_argument('--auth-uri', help='keystone internal endpoint e.g. http://CONTROLLER_VIP:5000', action='store', default=None, dest='auth_uri')
    install_parser.add_argument('--auth-url', help='keystone admin endpoint e.g. http://CONTROLLER_VIP:35357', action='store', default=None, dest='auth_url')
    install_parser.add_argument('--cinder-pass', help='password for cinder user', action='store', default=None, dest='cinder_pass')
    install_parser.add_argument('--my-ip', help='the host management ip', action='store', default=None, dest='my_ip')
    install_parser.add_argument('--glance-api-servers', help='glance host e.g. http://CONTROLLER_VIP:9292', action='store', default=None, dest='glance_api_servers')
    install_parser.add_argument('--rbd-secret-uuid', help='ceph rbd secret uuid', action='store', default=None, dest='rbd_secret_uuid')
    install_parser.add_argument('--memcached-servers', help='memcached servers e.g. CONTROLLER1:11211,CONTROLLER2:11211', action='store', default=None, dest='memcached_servers')
    install_parser.add_argument('--populate', help='Populate the cinder database', action='store_true', default=False, dest='populate')
    install_parser.set_defaults(func=install_f)
