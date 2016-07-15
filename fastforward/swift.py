import sys
from fastforward.cliutil import priority
from playback.api import Swift

def create_service_credentials_subparser(s):
    create_service_credentials_parser = s.add_parser('create-service-credentials', help='create the swift service credentials')
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
    create_service_credentials_parser.add_argument('--swift-pass',
                                                    help='password for swift user',
                                                    action='store',
                                                    default=None,
                                                    dest='swift_pass')
    create_service_credentials_parser.add_argument('--public-endpoint',
                                                    help=r'public endpoint for swift service e.g. "http://CONTROLLER_VIP:8080/v1/AUTH_%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint')
    create_service_credentials_parser.add_argument('--internal-endpoint',
                                                    help=r'internal endpoint for swift service e.g. "http://CONTROLLER_VIP:8080/v1/AUTH_%%\(tenant_id\)s"',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint')
    create_service_credentials_parser.add_argument('--admin-endpoint',
                                                    help='admin endpoint for swift service e.g. http://CONTROLLER_VIP:8080/v1',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint')
    return create_service_credentials_parser

def install_subparser(s):
    install_parser = s.add_parser('install',help='install swift proxy')
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
    install_parser.add_argument('--swift-pass',
                                help='password for swift user',
                                action='store',
                                default=None,
                                dest='swift_pass')
    install_parser.add_argument('--memcached-servers',
                                help='memcache servers e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                                action='store',
                                default=None,
                                dest='memcached_servers')
    install_parser.add_argument('--with-memcached',
                                help='install memcached on remote server, if you have other memcached on the controller node, you can use --memcached-sersers',
                                action='store_true',
                                default=False,
                                dest='with_memcached')
    return install_parser

def finalize_install_subparser(s):
    finalize_install_parser = s.add_parser('finalize-install', help='finalize swift installation')
    finalize_install_parser.add_argument('--swift-hash-path-suffix',
                                        help='swift_hash_path_suffix and swift_hash_path_prefix are used as part of the hashing algorithm when determining data placement in the cluster. These values should remain secret and MUST NOT change once a cluster has been deployed',
                                        action='store',
                                        default=None,
                                        dest='swift_hash_path_suffix')
    finalize_install_parser.add_argument('--swift-hash-path-prefix',
                                        help='swift_hash_path_suffix and swift_hash_path_prefix are used as part of the hashing algorithm when determining data placement in the cluster. These values should remain secret and MUST NOT change once a cluster has been deployed',
                                        action='store',
                                        default=None,
                                        dest='swift_hash_path_prefix')
    return finalize_install_parser

def make_target(args):
    try:
        target = Swift(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target

def create_service_credentials(args):
    target = make_target(args)
    target.create_service_credentials(
            args.os_password, 
            args.os_auth_url, 
            args.swift_pass, 
            args.public_endpoint,
            args.internal_endpoint,
            args.admin_endpoint)

def install(args):
    target = make_target(args)
    target.install(
            args.auth_uri, 
            args.auth_url, 
            args.swift_pass, 
            args.memcached_servers,
            args.with_memcached)

def finalize_install(args):
    target = make_target(args)
    target.finalize_install(
            args.swift_hash_path_suffix, 
            args.swift_hash_path_prefix)

@priority(22)
def make(parser):
    """provison Swift Proxy service with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def create_service_credentials_f(args):
        create_service_credentials(args)
    create_service_credentials_parser = create_service_credentials_subparser(s)
    create_service_credentials_parser.set_defaults(func=create_service_credentials_f)

    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)

    def finalize_install_f(args):
        finalize_install(args)
    finalize_install_parser = finalize_install_subparser(s)
    finalize_install_parser.set_defaults(func=finalize_install_f)

