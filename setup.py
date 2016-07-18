import sys, os

try:
    from setuptools import setup, find_packages
except ImportError:
    print("fastforward now needs setuptools in order to build. Install it using"
          " your package manager (usually python-setuptools) or via pip (pip"
          " install setuptools).")
    sys.exit(1)

from fastforward import __version__, __author__

def read(fname):
    path = os.path.join(os.path.dirname(__file__), fname)
    try:
        f = open(path)
    except IOError:
        return None
    return f.read()

setup(name='fastforward',
    version=__version__,
    description='FastForward is a DevOps automate platform',
    long_description=read('README.md'),
    author=__author__,
    author_email='jiasir@icloud.com',
    url='https://github.com/nofdev/fastforward',
    license='MIT',
    install_requires=['playback == 0.3.7'],
    packages=find_packages(),
    entry_points={ 
       'console_scripts': [
           'ff = fastforward.cli:main',
           ],

        'openstack': [
            'environment = fastforward.environment:make',
            'mysql = fastforward.mysql:make',
            'haproxy = fastforward.haproxy:make',
            'rabbitmq = fastforward.rabbitmq:make',
            'keystone = fastforward.keystone:make',
            'glance = fastforward.glance:make',
            'nova = fastforward.nova:make',
            'nova-compute = fastforward.nova_compute:make',
            'neutron = fastforward.neutron:make',
            'neutron-agent = fastforward.neutron_agent:make',
            'horizon = fastforward.horizon:make',
            'cinder = fastforward.cinder:make',
            'swift = fastforward.swift:make',
            'swift-storage = fastforward.swift_storage:make',
            'manila = fastforward.manila:make',
            'manila-share = fastforward.manila_share:make'
        ],
       },
    )
