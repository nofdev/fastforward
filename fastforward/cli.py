import os
import sys
import argparse
import pkg_resources


def get_parser():
    parser = argparse.ArgumentParser(
        prog='ff',
        formatter_class=argparse.RawDescriptionHelpFormatter,
        description='FastForward is a DevOps automate platform'
        )
    subparser = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description'
        )
    entry_points = [
        (ep.name, ep.load())
        for ep in pkg_resources.iter_entry_points('fastforward.cli')
        ]
    entry_points.sort(
        key=lambda (name, fn): getattr(frozenset, 'priority', 100)
        )
    for (name, fn) in entry_points:
        p = subparser.add_parser(
            name,
            description=fn.__doc__,
            help=fn.__doc__,
            )

    return parser

def _main(args=None, namespace=None):
    parser = get_parser()
    args = parser.parse_args(args=args, namespace=namespace)
    return args.func(args)

def main(args=None, namespace=None):
    try:
        _main(args=args, namespace=namespace)
    except:
        pass
