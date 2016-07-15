import sys
from playback.api import HaproxyInstall
from playback.api import HaproxyConfig
from playback.templates.haproxy_cfg import conf_haproxy_cfg
from fastforward.cliutil import priority

def install(args):
    try:
        target = HaproxyInstall(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError as e:
        sys.stderr.write(e.message)
        sys.exit(1)
    target.install()

def config(args):
    try:
        target = HaproxyConfig(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    if args.upload_conf:
        target.upload_conf(args.upload_conf)
    if args.configure_keepalived:
        target.configure_keepalived(args.router_id, args.priority, 
                args.state, args.interface, args.vip)

def gen_conf():
    with open('haproxy.cfg', 'w') as f:
        f.write(conf_haproxy_cfg)

@priority(12)
def make(parser):
    """provision HAProxy with Keepalived"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def install_f(args):
        install(args)
    install_parser = s.add_parser('install', help='install HAProxy')
    install_parser.set_defaults(func=install_f)

    def config_f(args):
        config(args)
    config_parser = s.add_parser('config', help='configure HAProxy')
    config_parser.add_argument('--upload-conf', help='upload configuration file to the target host', 
                               action='store', default=False, dest='upload_conf')
    config_parser.add_argument('--configure-keepalived', help='configure keepalived', 
                               action='store_true', default=False, dest='configure_keepalived')
    config_parser.add_argument('--router_id', help='Keepalived router id e.g. lb1', 
                               action='store', default=False, dest='router_id')
    config_parser.add_argument('--priority', help='Keepalived priority e.g. 150', 
                               action='store', default=False, dest='priority')
    config_parser.add_argument('--state', help='Keepalived state e.g. MASTER', 
                               action='store', default=False, dest='state')
    config_parser.add_argument('--interface', help='Keepalived binding interface e.g. eth0', 
                               action='store', default=False, dest='interface')
    config_parser.add_argument('--vip', help='Keepalived virtual ip e.g. CONTROLLER_VIP', 
                               action='store', default=False, dest='vip')
    config_parser.set_defaults(func=config_f)

    def gen_conf_f(args):
        gen_conf()
    gen_conf_parser = s.add_parser('gen-conf', help='generate the example configuration to the current location')
    gen_conf_parser.set_defaults(func=gen_conf_f)
