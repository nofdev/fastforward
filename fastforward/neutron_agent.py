import sys
from fastforward.cliutil import priority
from playback.api import NeutronAgent

def install_subparser(s):
    install_parser = s.add_parser('install', help='install neutron agent')
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
    install_parser.add_argument('--neutron-pass',
                                help='the password for neutron user',
                                action='store',
                                default=None,
                                dest='neutron_pass')
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
    install_parser.add_argument('--memcached-servers',
                                help='memcached servers e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                                action='store',
                                default=None,
                                dest='memcached_servers')
    return install_parser

def make_target(args):
    try:
        target = NeutronAgent(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target
    
def install(args):
    target = make_target(args)
    target.install(
            args.rabbit_hosts,
            args.rabbit_user,  
            args.rabbit_pass, 
            args.auth_uri, 
            args.auth_url, 
            args.neutron_pass, 
            args.public_interface, 
            args.local_ip, 
            args.memcached_servers)

@priority(19)
def make(parser):
    """provison Neutron agent nodes"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    
    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)
