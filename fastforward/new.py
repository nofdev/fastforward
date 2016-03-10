def main(parser):
    parser.add_argument(
        'openstack',
        action='store_true',
        default=False,
        help='provision openstack')
    parser.set_defaults(func=new)
    args.func(args)

def new(args):
    pass