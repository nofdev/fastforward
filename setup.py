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
    install_requires=['playback == 0.3.4'],
    packages=find_packages(),
    entry_points={ 
       'console_scripts': [
           'ff = fastforward.cli:main',
           ],

        'openstack': [
            'environment = fastforward.environment:make',
            #'mysql-deploy = fastforward.mysql:make',
            #'haproxy-deploy = fastforward.haproxy:make',
            #'rabbitmq-deploy = fastforward.rabbitmq:make',
            #'keystone-deploy = fastforward.keystone:make',
            #'glance-deploy = fastforward.glance:make',
            #'nova-deploy = fastforward.nova:make',
            #'nova-compute-deploy = fastforward.nova_compute:make',
            #'neutron-deploy = fastforward.neutron:make',
            #'neutron-agent-deploy = fastforward.neutron_agent:make',
            #'horizon-deploy = fastforward.horizon:make',
            'cinder = fastforward.cinder:make',
            #'swift-deploy = fastforward.swift:make',
            #'swift-storage-deploy = fastforward.swift_storage:make',
            #'manila-deploy = fastforward.manila:make',
            #'manila-share-deploy = fastforward.manila_share:make'
        ],
       },
    )
