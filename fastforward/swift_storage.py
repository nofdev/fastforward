import sys
from fastforward.cliutil import priority
from playback.api import SwiftStorage

def prepare_disks_subparser(s):
    prepare_disks_parser = s.add_parser('prepare-disks',
                                        help='prepare the disks for storage')
    prepare_disks_parser.add_argument('--name',
                                        help='the device name, e.g. sdb,sdc',
                                        action='store',
                                        default=None,
                                        dest='name')
    return prepare_disks_parser

def install_subparser(s):
    install_parser = s.add_parser('install', help='install swift storage')
    install_parser.add_argument('--address', 
                                help='the management interface ip for rsync', 
                                action='store', 
                                dest='address')
    install_parser.add_argument('--bind-ip', 
                                help='the management interface ip for swift storage binding', 
                                action='store', 
                                dest='bind_ip')
    return install_parser

def create_account_builder_file_subparser(s):
    create_account_builder_file_parser = s.add_parser('create-account-builder-file', help='create account ring')
    create_account_builder_file_parser.add_argument('--partitions', 
                                                    help='2^10 (1024) maximum partitions e.g. 10', 
                                                    action='store', 
                                                    default=None,
                                                    dest='partitions')
    create_account_builder_file_parser.add_argument('--replicas', 
                                                    help='3 replicas of each object e.g. 3', 
                                                    action='store', 
                                                    default=None,
                                                    dest='replicas')
    create_account_builder_file_parser.add_argument('--moving', 
                                                    help='1 hour minimum time between moving a partition more than once e.g. 1', 
                                                    action='store',
                                                    default=None, 
                                                    dest='moving')
    return create_account_builder_file_parser

def account_builder_add_subparser(s):
    account_builder_add_parser= s.add_parser('account-builder-add', help='Add each storage node to the account ring')
    account_builder_add_parser.add_argument('--region', 
                                            help='swift storage region e.g. 1', 
                                            action='store', 
                                            default=None,
                                            dest='region')
    account_builder_add_parser.add_argument('--zone', 
                                            help='swift storage zone e.g. 1', 
                                            action='store', 
                                            default=None,
                                            dest='zone')
    account_builder_add_parser.add_argument('--ip', 
                                            help='the IP address of the management network on the storage node e.g. STORAGE_NODE_IP', 
                                            action='store', 
                                            default=None,
                                            dest='ip')
    account_builder_add_parser.add_argument('--device', 
                                            help='a storage device name on the same storage node e.g. sdb', 
                                            action='store', 
                                            default=None,
                                            dest='device')
    account_builder_add_parser.add_argument('--weight',
                                            help='the storage device weight e.g. 100',
                                            action='store',
                                            default=None,
                                            dest='weight')
    return account_builder_add_parser

def create_container_builder_file_subparser(s):
    create_container_builder_file_parser = s.add_parser('create-container-builder-file', help='create container ring')
    create_container_builder_file_parser.add_argument('--partitions', 
                                                        help='2^10 (1024) maximum partitions e.g. 10', 
                                                        action='store', 
                                                        default=None,
                                                        dest='partitions')
    create_container_builder_file_parser.add_argument('--replicas', 
                                                        help='3 replicas of each object e.g. 3', 
                                                        action='store', 
                                                        default=None,
                                                        dest='replicas')
    create_container_builder_file_parser.add_argument('--moving', 
                                                        help='1 hour minimum time between moving a partition more than once e.g. 1', 
                                                        action='store',
                                                        default=None, 
                                                        dest='moving')
    return create_container_builder_file_parser

def container_builder_add_subparser(s):
    container_builder_add_parser = s.add_parser('container-builder-add', help='Add each storage node to the container ring')
    container_builder_add_parser.add_argument('--region', 
                                                help='swift storage region e.g. 1', 
                                                action='store', 
                                                default=None,
                                                dest='region')
    container_builder_add_parser.add_argument('--zone', 
                                                help='swift storage zone e.g. 1', 
                                                action='store', 
                                                default=None,
                                                dest='zone')
    container_builder_add_parser.add_argument('--ip', 
                                                help='the IP address of the management network on the storage node e.g. STORAGE_NODE_IP', 
                                                action='store', 
                                                default=None,
                                                dest='ip')
    container_builder_add_parser.add_argument('--device', 
                                                help='a storage device name on the same storage node e.g. sdb', 
                                                action='store', 
                                                default=None,
                                                dest='device')
    container_builder_add_parser.add_argument('--weight',
                                                help='the storage device weight e.g. 100',
                                                action='store',
                                                default=None,
                                                dest='weight')
    return container_builder_add_parser

def create_object_builder_file_subparser(s):
    create_object_builder_file_parser = s.add_parser('create-object-builder-file', help='create object ring')
    create_object_builder_file_parser.add_argument('--partitions', 
                                                    help='2^10 (1024) maximum partitions e.g. 10', 
                                                    action='store', 
                                                    default=None,
                                                    dest='partitions')
    create_object_builder_file_parser.add_argument('--replicas', 
                                                    help='3 replicas of each object e.g. 3', 
                                                    action='store', 
                                                    default=None,
                                                    dest='replicas')
    create_object_builder_file_parser.add_argument('--moving', 
                                                    help='1 hour minimum time between moving a partition more than once e.g. 1', 
                                                    action='store',
                                                    default=None, 
                                                    dest='moving')
    return create_object_builder_file_parser

def object_builder_add_subparser(s):
    object_builder_add_parser = s.add_parser('object-builder-add', help='Add each storage node to the object ring')
    object_builder_add_parser.add_argument('--region', 
                                            help='swift storage region e.g. 1', 
                                            action='store', 
                                            default=None,
                                            dest='region')
    object_builder_add_parser.add_argument('--zone', 
                                            help='swift storage zone e.g. 1', 
                                            action='store', 
                                            default=None,
                                            dest='zone')
    object_builder_add_parser.add_argument('--ip', 
                                            help='the IP address of the management network on the storage node e.g. STORAGE_NODE_IP', 
                                            action='store', 
                                            default=None,
                                            dest='ip')
    object_builder_add_parser.add_argument('--device', 
                                            help='a storage device name on the same storage node e.g. sdb', 
                                            action='store', 
                                            default=None,
                                            dest='device')
    object_builder_add_parser.add_argument('--weight',
                                            help='the storage device weight e.g. 100',
                                            action='store',
                                            default=None,
                                            dest='weight')
    return object_builder_add_parser

def sync_builder_file_subparser(s):
    sync_builder_file_parser = s.add_parser('sync-builder-file',
                                            help='copy the account.ring.gz, container.ring.gz, and object.ring.gz files to the /etc/swift directory on each storage node and any additional nodes running the proxy service')
    sync_builder_file_parser.add_argument('--to', 
                                            help='the target hosts where the *.ring.gz file to be added', 
                                            action='store', 
                                            default=None,
                                            dest='to')
    return sync_builder_file_parser

def account_builder_rebalance_subparser(s):
    account_builder_rebalance_parser = s.add_parser('account-builder-rebalance', help='Rebalance the account ring')
    return account_builder_rebalance_parser

def container_builder_rebalance_subparser(s):
    container_builder_rebalance_parser = s.add_parser('container-builder-rebalance', help='Rebalance the container ring')
    return container_builder_rebalance_parser

def object_builder_rebalance_subparser(s):
    object_builder_rebalance_parser = s.add_parser('object-builder-rebalance', help='Rebalance the object ring')
    return object_builder_rebalance_parser

def make_target(args):
    try:
        target = SwiftStorage(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        sys.stderr.write('No hosts found. Please using --hosts param.')
        sys.exit(1)
    return target

def prepare_disks(args):
    target = make_target(args)
    target.prepare_disks(args.name)

def install(args):
    target = make_target(args)
    target.install(args.address, args.bind_ip)

def create_account_builder_file(args):
    target = make_target(args)
    target.create_account_builder_file(args.partitions, args.replicas, args.moving)

def account_builder_add(args):
    target = make_target(args)
    target.account_builder_add(args.region, args.zone, 
            args.ip, args.device, args.weight)

def create_container_builder_file(args):
    target = make_target(args)
    target.create_container_builder_file(args.partitions,
            args.replicas,
            args.moving)

def container_builder_add(args):
    target = make_target(args)
    target.container_builder_add(args.region, 
            args.zone, args.ip, 
            args.device, args.weight)

def create_object_builder_file(args):
    target = make_target(args)
    target.create_object_builder_file(args.partitions,
            args.replicas, args.moving)

def object_builder_add(args):
    target = make_target(args)
    target.object_builder_add(args.region, args.zone, 
            args.ip, args.device, args.weight)

def sync_builder_file(args):
    target = make_target(args)
    target.get_builder_file()
    target.sync_builder_file(hosts=args.to.split(','))
    os.remove('account.ring.gz')
    os.remove('container.ring.gz')
    os.remove('object.ring.gz')

def account_builder_rebalance(args):
    target = make_target(args)
    target.account_builder_rebalance()

def container_builder_rebalance(args):
    target = make_target(args)
    target.container_builder_rebalance()

def object_builder_rebalance(args):
    target = make_target(args)
    target.object_builder_rebalance()

@priority(23)
def make(parser):
    """provison Swift Storage service"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def prepare_disks_f(args):
        prepare_disks(args)
    prepare_disks_parser = prepare_disks_subparser(s)
    prepare_disks_parser.set_defaults(func=prepare_disks_f)

    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)

    def create_account_builder_file_f(args):
        create_account_builder_file(args)
    create_account_builder_file_parser = create_account_builder_file_subparser(s)
    create_account_builder_file_parser.set_defaults(func=create_account_builder_file_f)

    def account_builder_add_f(args):
        account_builder_add(args)
    account_builder_add_parser = account_builder_add_subparser(s)
    account_builder_add_parser.set_defaults(func=account_builder_add_f)

    def create_container_builder_file_f(args):
        create_container_builder_file(args)
    create_container_builder_file_parser = create_container_builder_file_subparser(s)
    create_container_builder_file_parser.set_defaults(func=create_container_builder_file_f)

    def container_builder_add_f(args):
        container_builder_add(args)
    container_builder_add_parser = container_builder_add_subparser(s)
    container_builder_add_parser.set_defaults(func=container_builder_add_f)

    def create_object_builder_file_f(args):
        create_object_builder_file(args)
    create_object_builder_file_parser = create_object_builder_file_subparser(s)
    create_object_builder_file_parser.set_defaults(func=create_object_builder_file_f)

    def object_builder_add_f(args):
        object_builder_add(args)
    object_builder_add_parser = object_builder_add_subparser(s)
    object_builder_add_parser.set_defaults(func=object_builder_add_f)

    def sync_builder_file_f(args):
        sync_builder_file(args)
    sync_builder_file_parser = sync_builder_file_subparser(s)
    sync_builder_file_parser.set_defaults(func=sync_builder_file_f)

    def account_builder_rebalance_f(args):
        account_builder_rebalance(args)
    account_builder_rebalance_parser = account_builder_rebalance_subparser(s)
    account_builder_rebalance_parser.set_defaults(func=account_builder_rebalance_f)

    def container_builder_rebalance_f(args):
        container_builder_rebalance(args)
    container_builder_rebalance_parser = container_builder_rebalance_subparser(s)
    container_builder_rebalance_parser.set_defaults(func=container_builder_rebalance_f)

    def object_builder_rebalance_f(args):
        object_builder_rebalance(args)
    object_builder_rebalance_parser = object_builder_rebalance_subparser(s)
    object_builder_rebalance_parser.set_defaults(func=object_builder_rebalance_f)
