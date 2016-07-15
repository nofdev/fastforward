import sys
from playback.api import RabbitMq
from fastforward.cliutil import priority

def install_subparser(s):
    install_parser = s.add_parser('install', help='install RabbitMQ HA')
    install_parser.add_argument('--erlang-cookie', help='setup elang cookie',
                                action='store', default=None, dest='erlang_cookie')
    install_parser.add_argument('--rabbit-user', help='set rabbit user name',
                                action='store', default=None, dest='rabbit_user')
    install_parser.add_argument('--rabbit-pass', help='set rabbit password',
                                action='store', default=None, dest='rabbit_pass')
    return install_parser

def join_cluster_subparser(s):
    join_cluster_parser = s.add_parser('join-cluster', help='join the rabbit cluster')
    join_cluster_parser.add_argument('--name', help='the joined name, e.g. rabbit@CONTROLLER1',
                                    action='store', default=None, dest='name')
    return join_cluster_parser

def make_target(args):
    try:
        target = RabbitMq(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target

def install(args):
    target = make_target(args)
    target.install(args.erlang_cookie, args.rabbit_user, args.rabbit_pass)

def join_cluster(args):
    target = make_target(args)
    target.join_cluster(args.name)

@priority(13)
def make(parser):
    """provision RabbitMQ with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)

    def join_cluster_f(args):
        join_cluster(args)
    join_cluster_parser = join_cluster_subparser(s)
    join_cluster_parser.set_defaults(func=join_cluster_f)