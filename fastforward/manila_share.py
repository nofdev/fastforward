import sys
from fastforward.cliutil import priority
from playback.manila_share import ManilaShare

def install_subparser(s):
    install_parser = s.add_parser('install',help='install manila share node')
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
    install_parser.add_argument('--nova-pass',
                                help='passowrd for nova user',
                                action='store',
                                default=None,
                                dest='nova_pass')
    install_parser.add_argument('--cinder-pass',
                                help='passowrd for cinder user',
                                action='store',
                                default=None,
                                dest='cinder_pass')
                            
    return install_parser

def make_target(args):
    try:
        target = ManilaShare(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target

def install(args):
    target = make_target(args)
    target.install_manila_share(
            args.connection,
            args.auth_uri,
            args.auth_url,
            args.manila_pass,
            args.my_ip,
            args.memcached_servers,
            args.rabbit_hosts,
            args.rabbit_user,
            args.rabbit_pass,
            args.neutron_endpoint,
            args.neutron_pass,
            args.nova_pass,
            args.cinder_pass)

@priority(25)
def make(parser):
    """provison Manila with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)
