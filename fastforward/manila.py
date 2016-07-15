import sys
from fastforward.cliutil import priority
from playback.manila import Manila

def create_manila_db_subparser(s):
    create_manila_db_parser = s.add_parser('create-manila-db', help='create manila database')
    create_manila_db_parser.add_argument('--root-db-pass',
                                            help='the MySQL database root passowrd',
                                            action='store',
                                            default=None,
                                            dest='root_db_pass')
    create_manila_db_parser.add_argument('--manila-db-pass',
                                            help='manila database password',
                                            action='store',
                                            default=None,
                                            dest='manila_db_pass')
    return create_manila_db_parser

def create_service_credentials_subparser(s):
    create_service_credentials_parser = s.add_parser('create-service-credentials', help='create the service credentials')
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
    create_service_credentials_parser.add_argument('--manila-pass',
                                                    help='passowrd for manila user',
                                                    action='store',
                                                    default=None,
                                                    dest='manila_pass')
    create_service_credentials_parser.add_argument('--public-endpoint-v1',
                                                    help=r'public endpoint for manila service e.g. "http://CONTROLLER_VIP:8786/v1/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint_v1')
    create_service_credentials_parser.add_argument('--internal-endpoint-v1',
                                                    help=r'internal endpoint for manila service e.g. "http://CONTROLLER_VIP:8786/v1/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint_v1')
    create_service_credentials_parser.add_argument('--admin-endpoint-v1',
                                                    help=r'admin endpoint for manila service e.g. "http://CONTROLLER_VIP:8786/v1/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint_v1')
    create_service_credentials_parser.add_argument('--public-endpoint-v2',
                                                    help=r'public endpoint for manila service e.g. "http://CONTROLLER_VIP:8786/v2/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint_v2')
    create_service_credentials_parser.add_argument('--internal-endpoint-v2',
                                                    help=r'internal endpoint for manila service e.g. "http://CONTROLLER_VIP:8786/v2/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint_v2')
    create_service_credentials_parser.add_argument('--admin-endpoint-v2',
                                                    help=r'admin endpoint for manila service e.g. "http://CONTROLLER_VIP:8786/v2/%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint_v2')
    return create_service_credentials_parser

def install_subparser(s):
    install_parser = s.add_parser('install',help='install manila')
    install_parser.add_argument('--connection',
                                help='mysql manila database connection string e.g. mysql+pymysql://manila:MANILA_PASS@CONTROLLER_VIP/manila',
                                action='store',
                                default=None,
                                dest='connection')
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
    install_parser.add_argument('--manila-pass',
                                help='passowrd for manila user',
                                action='store',
                                default=None,
                                dest='manila_pass')
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
    install_parser.add_argument('--populate',
                                help='Populate the manila database',
                                action='store_true',
                                default=False,
                                dest='populate')
    return install_parser

def make_target(args):
    try:
        target = Manila(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target

def create_manila_db(args):
    target = make_target(args)
    target.create_manila_db(args.root_db_pass, args.manila_db_pass)

def create_service_credentials(args):
    target = make_target(args)
    target.create_service_credentials(
            args.os_password,
            args.os_auth_url,
            args.manila_pass,
            args.public_endpoint_v1,
            args.internal_endpoint_v1,
            args.admin_endpoint_v1,
            args.public_endpoint_v2,
            args.internal_endpoint_v2,
            args.admin_endpoint_v2)

def install(args):
    target = make_target(args)
    target.install_manila(
            args.connection,
            args.auth_uri,
            args.auth_url,
            args.manila_pass,
            args.my_ip,
            args.memcached_servers,
            args.rabbit_hosts,
            args.rabbit_user,
            args.rabbit_pass,
            args.populate)

@priority(24)
def make(parser):
    """provison Manila with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def create_manila_db_f(args):
        create_manila_db(args)
    create_manila_db_parser = create_manila_db_subparser(s)
    create_manila_db_parser.set_defaults(func=create_manila_db_f)

    def create_service_credentials_f(args):
        create_service_credentials(args)
    create_service_credentials_parser = create_service_credentials_subparser(s)
    create_service_credentials_parser.set_defaults(func=create_service_credentials_f)

    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)

