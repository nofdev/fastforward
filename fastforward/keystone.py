from fastforward.cliutil import priority
from playback.api import Keystone

def make_target(args):
    try:
        target = Keystone(user=args.user, hosts=args.hosts.split(','), key_filename=args.key_filename, password=args.password)
    except AttributeError:
        err_hosts = 'No hosts found. Please using --hosts param.'
        sys.stderr.write(err_hosts)
        sys.exit(1)
    
    return target

def create_keystone_db(args):
    target = make_target(args)
    target.create_keystone_db(
            args.root_db_pass, 
            args.keystone_db_pass)

def install(args):
    target = make_target(args)
    target.install_keystone(
            args.admin_token, 
            args.connection, 
            args.memcached_servers,
            args.populate)

def create_entity_and_endpoint(args):
    target = make_target(args)
    target.create_entity_and_endpoint(
        args.os_token,
        args.os_url,
        args.public_endpoint,
        args.internal_endpoint,
        args.admin_endpoint)

def create_projects_users_roles(args):
    target = make_target(args)
    target.create_projects_users_roles(
        args.os_token,
        args.os_url,
        args.admin_pass,
        args.demo_pass)
    target.update_keystone_paste_ini()

def create_keystone_db_subparser(s):
    create_keystone_db_parser = s.add_parser('create-keystone-db', help='create the keystone database', )
    create_keystone_db_parser.add_argument('--root-db-pass', 
                                           help='the openstack database root passowrd',
                                           action='store', 
                                           default=None, 
                                           dest='root_db_pass')
    create_keystone_db_parser.add_argument('--keystone-db-pass', 
                                           help='keystone db passowrd',
                                           action='store', 
                                           default=None, 
                                           dest='keystone_db_pass')
    return create_keystone_db_parser

def install_subparser(s):
    install_parser = s.add_parser('install', help='install keystone')
    install_parser.add_argument('--admin-token',
                                help='define the value of the initial administration token',
                                action='store',
                                default=None,
                                dest='admin_token')
    install_parser.add_argument('--connection',
                                help='database connection string e.g. mysql+pymysql://keystone:PASS@CONTROLLER_VIP/keystone',
                                action='store',
                                default=None,
                                dest='connection')
    install_parser.add_argument('--memcached-servers',
                                help='memcached servers. e.g. CONTROLLER1:11211,CONTROLLER2:11211',
                                action='store',
                                default=None,
                                dest='memcached_servers')
    install_parser.add_argument('--populate',
                                help='populate the keystone database',
                                action='store_true',
                                default=False,
                                dest='populate')
    return install_parser

def create_entity_and_endpoint_subparser(s):
    create_entity_and_endpoint_parser = s.add_parser('create-entity-and-endpoint',
                                                       help='create the service entity and API endpoints',)
    create_entity_and_endpoint_parser.add_argument('--os-token',
                                                    help='the admin token',
                                                    action='store',
                                                    default=None,
                                                    dest='os_token')
    create_entity_and_endpoint_parser.add_argument('--os-url',
                                                    help='keystone endpoint url e.g. http://CONTROLLER_VIP:35357/v3',
                                                    action='store',
                                                    default=None,
                                                    dest='os_url')
    create_entity_and_endpoint_parser.add_argument('--public-endpoint',
                                                    help='the public endpoint e.g. http://CONTROLLER_VIP:5000/v3',
                                                    action='store',
                                                    default=None,
                                                    dest='public_endpoint')
    create_entity_and_endpoint_parser.add_argument('--internal-endpoint',
                                                    help='the internal endpoint e.g. http://CONTROLLER_VIP:5000/v3',
                                                    action='store',
                                                    default=None,
                                                    dest='internal_endpoint')
    create_entity_and_endpoint_parser.add_argument('--admin-endpoint',
                                                    help='the admin endpoint e.g. http://CONTROLLER_VIP:35357/v3',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_endpoint')
    return create_entity_and_endpoint_parser

def create_projects_users_roles_subparser(s):
    create_projects_users_roles_parser = s.add_parser('create-projects-users-roles',
                                                        help='create an administrative and demo project, user, and role for administrative and testing operations in your environment')
    create_projects_users_roles_parser.add_argument('--os-token',
                                                    help='the admin token',
                                                    action='store',
                                                    default=None,
                                                    dest='os_token')
    create_projects_users_roles_parser.add_argument('--os-url',
                                                    help='keystone endpoint url e.g. http://CONTROLLER_VIP:35357/v3',
                                                    action='store',
                                                    default=None,
                                                    dest='os_url')
    create_projects_users_roles_parser.add_argument('--admin-pass',
                                                    help='passowrd for admin user',
                                                    action='store',
                                                    default=None,
                                                    dest='admin_pass')
    create_projects_users_roles_parser.add_argument('--demo-pass',
                                                    help='passowrd for demo user',
                                                    action='store',
                                                    default=None,
                                                    dest='demo_pass')
    return create_projects_users_roles_parser

@priority(14)
def make(parser):
    """provison Keystone with HA"""
    s = parser.add_subparsers(
        title='commands',
        metavar='COMMAND',
        help='description',
        )

    def create_keystone_db_f(args):
        create_keystone_db(args)
    create_keystone_db_parser = create_keystone_db_subparser(s)
    create_keystone_db_parser.set_defaults(func=create_keystone_db_f)

    def install_f(args):
        install(args)
    install_parser = install_subparser(s)
    install_parser.set_defaults(func=install_f)

    def create_entity_and_endpoint_f(args):
        create_entity_and_endpoint(args)
    create_entity_and_endpoint_parser = create_entity_and_endpoint_subparser(s)
    create_entity_and_endpoint_parser.set_defaults(func=create_entity_and_endpoint_f)

    def create_projects_users_roles_f(args):
        create_projects_users_roles(args)
    create_projects_users_roles_parser = create_projects_users_roles_subparser(s)
    create_projects_users_roles_parser.set_defaults(func=create_projects_users_roles_f)
