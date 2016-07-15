import sys
import argparse
import pkg_resources

from fastforward import __version__ 

def get_parser():
    parser = argparse.ArgumentParser(
        formatter_class=argparse.RawDescriptionHelpFormatter,
        description='FastForward is a DevOps automate platform'
        )
    parser.add_argument(
        '-v', '--version',
        action='version', version=__version__ ,
    )
    parser.add_argument(
        '--user',
        help='the username to connect to the remote host', action='store', default='ubuntu', dest='user'
    )
    parser.add_argument(
        '--hosts',
        help='the remote host to connect to ', action='store', default=None, dest='hosts'
    )
    parser.add_argument(
        '-i', '--key-filename',
        help='referencing file paths to SSH key files to try when connecting', action='store', dest='key_filename', default=None
    )
    parser.add_argument(
        '--password',
        help='the password used by the SSH layer when connecting to remote hosts', action='store', dest='password', default=None
    )
    subparser = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    openstack_parser = subparser.add_parser(
       'openstack',
       help='provision and manage OpenStack'
        )
    openstack_subparser = openstack_parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
    )
    entry_points = [
        (ep.name, ep.load())
        for ep in pkg_resources.iter_entry_points('openstack')
        ]
    entry_points.sort(
        key=lambda (name, fn): getattr(fn, 'priority', 100),
        )
    for (name, fn) in entry_points:
        p = openstack_subparser.add_parser(
            name,
            description=fn.__doc__,
            help=fn.__doc__,
            )
        fn(p)
    return parser

def _main():
    parser = get_parser()
    if len (sys.argv) < 2 :
        parser.print_help()
        sys.exit()
    else :
        args = parser.parse_args()
    return args.func(args)

def main():
    try:
        _main()
    except:
        pass

if __name__ == '__main__':
    main()