import sys
from playback.api import MysqlConfig
from playback.api import MysqlManage
from playback.api import MysqlInstallation
from fastforward.cliutil import priority

def install(args):
    try:
        target = MysqlInstallation(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        err_hosts = 'No hosts found. Please using --hosts param.'
        sys.stderr.write(err_hosts)
        sys.exit(1)
    target.enable_repo()
    target.install()

def config(args):
    try:
        target = MysqlConfig(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        err_hosts = 'No hosts found. Please using --hosts param.'
        sys.stderr.write(err_hosts)
        sys.exit(1)
    target.update_mysql_config(args.wsrep_cluster_address, 
            args.wsrep_node_name, args.wsrep_node_address)

def manage(args):
    try:
        target = MysqlManage(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        err_hosts = 'No hosts found. Please using --hosts param.'
        sys.stderr.write(err_hosts)
        sys.exit(1)
    if args.wsrep_new_cluster:
        target.start_wsrep_new_cluster()
    if args.start:
        target.start_mysql()
    if args.stop:
        target.stop_mysql()
    if args.change_root_password:
        target.change_root_password(args.change_root_password)
    if args.show_cluster_status:
        if args.root_db_pass == None:
            raise Exception('--root-db-pass is empty\n')
        target.show_cluster_status(args.root_db_pass)

def install_subparser(s):
    install_parser = s.add_parser('install', help='install Galera Cluster for MySQL')
    return install_parser

def config_subparser(s):
    config_parser = s.add_parser('config', help='setup Galera Cluster for MySQL')
    config_parser.add_argument('--wsrep-cluster-address', help='the IP addresses for each cluster node e.g. gcomm://CONTROLLER1_IP,CONTROLLER2_IP', 
                                action='store', dest='wsrep_cluster_address')
    config_parser.add_argument('--wsrep-node-name', help='the logical name of the cluster node e.g. galera1', 
                                action='store', dest='wsrep_node_name')
    config_parser.add_argument('--wsrep-node-address', help='the IP address of the cluster node e.g. CONTROLLER1_IP', 
                                action='store', dest='wsrep_node_address')
    return config_parser

def manage_subparser(s):
    manage_parser = s.add_parser('manage', help='manage Galera Cluster for MySQL')
    manage_parser.add_argument('--wsrep-new-cluster', help='initialize the Primary Component on one cluster node',
                                action='store_true', default=False, dest='wsrep_new_cluster')
    manage_parser.add_argument('--start', help='start the database server on all other cluster nodes',
                                action='store_true', default=False, dest='start')
    manage_parser.add_argument('--stop', help='stop the database server',
                                action='store_true', default=False, dest='stop')
    manage_parser.add_argument('--change-root-password', help='change the root password',
                                action='store', default=False, dest='change_root_password')
    manage_parser.add_argument('--show-cluster-status', help='show the cluster status',
                                action='store_true', default=False, dest='show_cluster_status')
    manage_parser.add_argument('--root-db-pass', help='the password of root user',
                                action='store', default=None, dest='root_db_pass')
    return manage_parser

@priority(11)
def make(parser):
    """provision MariaDB Galera Cluster"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )
    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)

    def config_f(args):
        config(args)
    config_parser = config_subparser(s)
    config_parser.set_defaults(func=config_f)

    def manage_f(args):
        manage(args)
    manage_parser = manage_subparser(s)
    manage_parser.set_defaults(func=manage_f)
