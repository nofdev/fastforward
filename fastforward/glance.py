from fastforward.cliutil import priority
from playback.api import Glance

def make_target(user, hosts, key_filename, password):
    try:
        target = Glance(user, hosts, key_filename, password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)

    return target

def create_glance_db(args):
    target = make_target(args.user, args.hosts.split(','), args.key_filename, args.password)
    target.create_glance_db(
            args.root_db_pass, 
            args.glance_db_pass)

def create_service_credentials(args):
    target = make_target(args.user, args.hosts.split(','), args.key_filename, args.password)
    target.create_service_credentials(
            args.os_password, 
            args.os_auth_url, 
            args.glance_pass, 
            args.public_endpoint,
            args.internal_endpoint,
            args.admin_endpoint)

def install(args):
    target = make_target(args.user, args.hosts.split(','), args.key_filename, args.password)
    target.install_glance(
            args.connection, 
            args.auth_uri, 
            args.auth_url, 
            args.glance_pass,
            args.swift_store_auth_address,
            args.memcached_servers,
            args.populate)

@priority(15)
def make(parser):
    """provison Glance with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def create_glance_db_f(args):
        create_glance_db(args)
    create_glance_db_parser = s.add_parser('create-glance-db', help='create the glance database')
    create_glance_db_parser.add_argument('--root-db-pass', 
                                          help='the openstack database root passowrd',
                                          action='store', 
                                          default=None, 
                                          dest='root_db_pass')
    create_glance_db_parser.add_argument('--glance-db-pass', 
                                          help='glance db passowrd',
                                          action='store', 
                                          default=None, 
                                          dest='glance_db_pass')
    create_glance_db_parser.set_defaults(func=create_glance_db)

    def create_service_credentials_f(args):
        create_service_credentials(args)
    create_service_credentials_parser = s.add_parser('create-service-credentials',
                                                   help='create the glance service credentials')
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
    create_service_credentials_parser.add_argument('--glance-pass',
                                                    help='passowrd for glance user',
                                                    action='store',
                                                    default=None,
                                                    dest='glance_pass')
    create_service_credentials_parser.add_argument('--public-endpoint',
                                                    help='public endpoint for glance service e.g. http://CONTROLLER_VIP:9292',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint')
    create_service_credentials_parser.add_argument('--internal-endpoint',
                                                    help='internal endpoint for glance service e.g. http://CONTROLLER_VIP:9292',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint')
    create_service_credentials_parser.add_argument('--admin-endpoint',
                                                    help='admin endpoint for glance service e.g. http://CONTROLLER_VIP:9292',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint')
    create_service_credentials_parser.set_defaults(func=create_service_credentials)

    def install_f(args):
        install(args)
    install_parser = s.add_parser('install', help='install glance(default store: ceph)')
    install_parser.add_argument('--connection',
                        help='mysql database connection string e.g. mysql+pymysql://glance:GLANCE_PASS@CONTROLLER_VIP/glance',
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
    install_parser.add_argument('--glance-pass',
                        help='passowrd for glance user',
                        action='store',
                        default=None,
                        dest='glance_pass')
    install_parser.add_argument('--swift-store-auth-address',
                        help='DEPRECATED the address where the Swift authentication service is listening e.g. http://CONTROLLER_VIP:5000/v3/',
                        action='store',
                        default='DEPRECATED_BY_PLAYBACK',
                        dest='swift_store_auth_address')
    install_parser.add_argument('--memcached-servers',
                        help='memcached servers e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                        action='store',
                        default=None,
                        dest='memcached_servers')                        
    install_parser.add_argument('--populate',
                        help='populate the glance database',
                        action='store_true',
                        default=False,
                        dest='populate')
    install_parser.set_defaults(func=install_f)
