import sys
from fastforward.cliutil import priority
from playback.api import Nova

def create_nova_db_subparser(s):
    create_nova_db_parser = s.add_parser('create-nova-db',
                                        help='create the nova and nova_api database')
    create_nova_db_parser.add_argument('--root-db-pass', 
                                        help='the MySQL database root passowrd',
                                        action='store', 
                                        default=None, 
                                        dest='root_db_pass')
    create_nova_db_parser.add_argument('--nova-db-pass', 
                                        help='nova and nova_api database passowrd',
                                        action='store', 
                                        default=None, 
                                        dest='nova_db_pass')
    return create_nova_db_parser

def create_service_credentials_subparser(s):
    create_service_credentials_parser = s.add_parser('create-service-credentials', help='create the nova service credentials',)
    create_service_credentials_parser.add_argument('--os-password',
                                                    help='the password for admin user',
                                                    action='store',
                                                    default=None,
                                                    dest='os_password')
    create_service_credentials_parser.add_argument('--os-auth-url',
                                                    help='keystone endpoint url e.g. http://CONTROLLER_VIP:35357/v3',
                                                    action='store',
                                                    default=None,
                                                    dest='os_auth_url')
    create_service_credentials_parser.add_argument('--nova-pass',
                                                    help='passowrd for nova user',
                                                    action='store',
                                                    default=None,
                                                    dest='nova_pass')
    create_service_credentials_parser.add_argument('--public-endpoint',
                                                    help=r'public endpoint for nova service e.g. "http://CONTROLLER_VIP:8774/v2.1/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint')
    create_service_credentials_parser.add_argument('--internal-endpoint',
                                                    help=r'internal endpoint for nova service e.g. "http://CONTROLLER_VIP:8774/v2.1/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint')
    create_service_credentials_parser.add_argument('--admin-endpoint',
                                                    help=r'admin endpoint for nova service e.g. "http://CONTROLLER_VIP:8774/v2.1/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint')
    return create_service_credentials_parser
    
def install_subparser(s):
    install_parser = s.add_parser('install',help='install nova')
    install_parser.add_argument('--connection',
                                help='mysql nova database connection string e.g. mysql+pymysql://nova:NOVA_PASS@CONTROLLER_VIP/nova',
                                action='store',
                                default=None,
                                dest='connection')
    install_parser.add_argument('--api-connection',
                                help='mysql nova_api database connection string e.g. mysql+pymysql://nova:NOVA_PASS@CONTROLLER_VIP/nova_api',
                                action='store',
                                default=None,
                                dest='api_connection')
    install_parser.add_argument('--auth-uri',
                                help='keystone internal endpoint e.g. http://CONTROLLER_VIP:5000',
                                action='store',
                                default=None,
                                dest='auth_uri')
    install_parser.add_argument('--auth-url',
                                help='keystone admin endpoint e.g. http://CONTROLLER_VIP:35357',
                                action='store',
                                default=None,
                                dest='auth_url')
    install_parser.add_argument('--nova-pass',
                                help='passowrd for nova user',
                                action='store',
                                default=None,
                                dest='nova_pass')
    install_parser.add_argument('--my-ip',
                                help='the host management ip',
                                action='store',
                                default=None,
                                dest='my_ip')
    install_parser.add_argument('--memcached-servers',
                                help='memcached servers e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                                action='store',
                                default=None,
                                dest='memcached_servers')
    install_parser.add_argument('--rabbit-hosts',
                                help='rabbit hosts e.g. CONTROLLER1,CONTROLLER2',
                                action='store',
                                default=None,
                                dest='rabbit_hosts')
    install_parser.add_argument('--rabbit-user',
                                help='the user for rabbit openstack user, default openstack',
                                action='store',
                                default='openstack',
                                dest='rabbit_user')
    install_parser.add_argument('--rabbit-pass',
                                help='the password for rabbit openstack user',
                                action='store',
                                default=None,
                                dest='rabbit_pass')
    install_parser.add_argument('--glance-api-servers',
                                help='glance host e.g. http://CONTROLLER_VIP:9292',
                                action='store',
                                default=None,
                                dest='glance_api_servers')
    install_parser.add_argument('--neutron-endpoint',
                                help='neutron endpoint e.g. http://CONTROLLER_VIP:9696',
                                action='store',
                                default=None,
                                dest='neutron_endpoint')
    install_parser.add_argument('--neutron-pass',
                                help='the password for neutron user',
                                action='store',
                                default=None,
                                dest='neutron_pass')
    install_parser.add_argument('--metadata-proxy-shared-secret',
                                help='metadata proxy shared secret',
                                action='store',
                                default=None,
                                dest='metadata_proxy_shared_secret')
    install_parser.add_argument('--populate',
                                help='Populate the nova database',
                                action='store_true',
                                default=False,
                                dest='populate')
    return install_parser
    
def make_target(args):
    try:
        target = Nova(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target
    
def create_nova_db(args):
    target = make_target(args)
    target.create_nova_db(args.root_db_pass, args.nova_db_pass)

def create_service_credentials(args):
    target = make_target(args)
    target.create_service_credentials(args.os_password, 
            args.os_auth_url, args.nova_pass, 
            args.public_endpoint, args.internal_endpoint, 
            args.admin_endpoint)

def install(args):
    target = make_target(args)
    target.install_nova(args.connection, args.api_connection, args.auth_uri, args.auth_url,
            args.nova_pass, args.my_ip, args.memcached_servers, args.rabbit_hosts, args.rabbit_user, 
            args.rabbit_pass, args.glance_api_servers, args.neutron_endpoint, args.neutron_pass,
            args.metadata_proxy_shared_secret, args.populate)

@priority(16)
def make(parser):
    """provison Nova with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def create_nova_db_f(args):
        create_nova_db(args)
    create_nova_db_parser = create_nova_db_subparser(s)
    create_nova_db_parser.set_defaults(func=create_nova_db_f)
    
    def create_service_credentials_f(args):
        create_service_credentials(args)
    create_service_credentials_parser = create_service_credentials_subparser(s)
    create_service_credentials_parser.set_defaults(func=create_service_credentials_f)
    
    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)
