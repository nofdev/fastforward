import os
import sys
from playback.api import Cmd
from playback.api import PrepareHost
from fastforward.cliutil import priority

def prepare_host(user, hosts, key_filename, password, public_interface): 
    try:
        remote = PrepareHost(user, hosts, key_filename, password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)

    # host networking
    remote.setup_external_interface(public_interface)

    # ntp
    remote.setup_ntp()

    # openstack packages
    remote.set_openstack_repository()

def gen_pass():
    os.system('openssl rand -hex 10')
    
def cmd(user, hosts, key_filename, password, run):
    try:
        remote = Cmd(user, hosts, key_filename, password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    remote.cmd(run)

@priority(10)
def make(parser):
    """prepare OpenStack basic environment"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def prepare_host_f(args):
        prepare_host(args.user, args.hosts.split(','), args.key_filename, args.password, args.public_interface)
    prepare_host_parser = s.add_parser('prepare-host', help='prepare the OpenStack environment')
    prepare_host_parser.add_argument('--public-interface', help='public interface e.g. eth1', action='store', default='eth1', dest='public_interface')
    prepare_host_parser.set_defaults(func=prepare_host_f)
    
    def gen_pass_f(args):
        gen_pass()
    gen_pass_parser = s.add_parser('gen-pass', help='generate the password')
    gen_pass_parser.set_defaults(func=gen_pass_f)
    
    def cmd_f(args):
        cmd(args.user, args.hosts.split(','), args.key_filename, args.password, args.run)
    cmd_parser = s.add_parser('cmd', help='run command line on the target host')
    cmd_parser.add_argument('--run', help='the command running on the remote node', action='store', default=None, dest='run')
    cmd_parser.set_defaults(func=cmd_f)
