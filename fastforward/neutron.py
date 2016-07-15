import sys
from fastforward.cliutil import priority
from playback.api import Neutron

def create_neutron_db_subparser(s):
    create_neutron_db_parser = s.add_parser('create-neutron-db',
                                            help='create the neutron database')
    create_neutron_db_parser.add_argument('--root-db-pass', 
                                            help='the openstack database root passowrd',
                                            action='store', 
                                            default=None, 
                                            dest='root_db_pass')
    create_neutron_db_parser.add_argument('--neutron-db-pass', 
                                            help='neutron db passowrd',
                                            action='store', 
                                            default=None, 
                                            dest='neutron_db_pass')
    return create_neutron_db_parser
    
def create_service_credentials_subparser(s):
    create_service_credentials_parser = s.add_parser('create-service-credentials',
                                                            help='create the neutron service credentials')
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
    create_service_credentials_parser.add_argument('--neutron-pass',
                                                    help='the password for neutron user',
                                                    action='store',
                                                    default=None,
                                                    dest='neutron_pass')
    create_service_credentials_parser.add_argument('--public-endpoint',
                                                    help='public endpoint for neutron service e.g. http://CONTROLLER_VIP:9696',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint')
    create_service_credentials_parser.add_argument('--internal-endpoint',
                                                    help='internal endpoint for neutron service e.g. http://CONTROLLER_VIP:9696',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint')
    create_service_credentials_parser.add_argument('--admin-endpoint',
                                                    help='admin endpoint for neutron service e.g. http://CONTROLLER_VIP:9696',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint')
    return create_service_credentials_parser

def install_subparser(s):
    install_parser = s.add_parser('install', help='install neutron for self-service')
    install_parser.add_argument('--connection',
                        help='mysql database connection string e.g. mysql+pymysql://neutron:NEUTRON_PASS@CONTROLLER_VIP/neutron',
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
    install_parser.add_argument('--rabbit-hosts',
                        help='rabbit hosts e.g. controller1,controller2',
                        action='store',
                        default=None,
                        dest='rabbit_hosts')
    install_parser.add_argument('--rabbit-user',
                        help='the user for rabbit, default openstack',
                        action='store',
                        default='openstack',
                        dest='rabbit_user')
    install_parser.add_argument('--rabbit-pass',
                        help='the password for rabbit openstack user',
                        action='store',
                        default=None,
                        dest='rabbit_pass')
    install_parser.add_argument('--neutron-pass',
                        help='the password for neutron user',
                        action='store',
                        default=None,
                        dest='neutron_pass')
    install_parser.add_argument('--nova-url',
                        help='URL for connection to nova (Only supports one nova region currently) e.g. http://CONTROLLER_VIP:8774/v2.1',
                        action='store',
                        default=None,
                        dest='nova_url')
    install_parser.add_argument('--nova-pass',
                        help='passowrd for nova user',
                        action='store',
                        default=None,
                        dest='nova_pass')
    install_parser.add_argument('--public-interface',
                        help='public interface e.g. eth1',
                        action='store',
                        default=None,
                        dest='public_interface')
    install_parser.add_argument('--local-ip',
                        help=' underlying physical network interface that handles overlay networks(uses the management interface IP)',
                        action='store',
                        default=None,
                        dest='local_ip')
    install_parser.add_argument('--nova-metadata-ip',
                        help='IP address used by Nova metadata server e.g. CONTROLLER_VIP',
                        action='store',
                        default=None,
                        dest='nova_metadata_ip')
    install_parser.add_argument('--metadata-proxy-shared-secret',
                        help='metadata proxy shared secret',
                        action='store',
                        default=None,
                        dest='metadata_proxy_shared_secret')
    install_parser.add_argument('--memcached-servers',
                        help='memcached servers e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                        action='store',
                        default=None,
                        dest='memcached_servers')
    install_parser.add_argument('--populate',
                        help='Populate the neutron database',
                        action='store_true',
                        default=False,
                        dest='populate')
    return install_parser

def make_target(args):
    try:
        target = Neutron(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except Exception:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target
    
def create_neutron_db(args):
    target = make_target(args)
    target.create_neutron_db(args.root_db_pass, args.neutron_db_pass)

def create_service_credentials(args):
    target = make_target(args)
    target.create_service_credentials(
            args.os_password,
            args.os_auth_url,
            args.neutron_pass,
            args.public_endpoint,
            args.internal_endpoint,
            args.admin_endpoint)
            
def install(args):
    target = make_target(args)
    target.install_self_service(
            args.connection, 
            args.rabbit_hosts, 
            args.rabbit_user, 
            args.rabbit_pass, 
            args.auth_uri, 
            args.auth_url, 
            args.neutron_pass, 
            args.nova_url, 
            args.nova_pass, 
            args.public_interface, 
            args.local_ip, 
            args.nova_metadata_ip, 
            args.metadata_proxy_shared_secret, 
            args.memcached_servers,
            args.populate)

@priority(18)
def make(parser):
    """provison Neutron with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def create_neutron_db_f(args):
        create_neutron_db(args)
    create_neutron_db_parser = create_neutron_db_subparser(s)
    create_neutron_db_parser.set_defaults(func=create_neutron_db_f)
    
    def create_service_credentials_f(args):
        create_service_credentials(args)
    create_service_credentials_parser = create_service_credentials_subparser(s)
    create_service_credentials_parser.set_defaults(func=create_service_credentials_f)

    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)
