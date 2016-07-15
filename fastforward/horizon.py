import sys
from fastforward.cliutil import priority
from playback.api import Horizon

def install(args):
    try:
        target = Horizon(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)

    target.install(
            args.openstack_host, 
            args.memcached_servers, 
            args.time_zone)

@priority(20)
def make(parser):
    """provison Horizon with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def install_f(args):
        install(args)
    install_parser = s.add_parser('install', help='install horizon')
    install_parser.add_argument('--openstack-host',
                                help='configure the dashboard to use OpenStack services on the controller node e.g. CONTROLLER_VIP',
                                action='store',
                                default=None,
                                dest='openstack_host')
    install_parser.add_argument('--memcached-servers',
                                help='django memcache e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                                action='store',
                                default=None,
                                dest='memcached_servers')
    install_parser.add_argument('--time-zone',
                                help='the timezone of the server. This should correspond with the timezone of your entire OpenStack installation e.g. Asia/Shanghai',
                                action='store',
                                default=None,
                                dest='time_zone')
    install_parser.set_defaults(func=install_f)
