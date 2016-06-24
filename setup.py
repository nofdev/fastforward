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
    author_email='taio@outlook.com',
    url='https://github.com/nofdev/fastforward',
    license='GPLv2',
    install_requires=['playback == 0.2.3'],
    packages=find_packages(),
    entry_points={ 
       'console_scripts': [
           'ff-openstack-env = playback.env:main',
           'ff-openstack-mysql = playback.mysql:main',
           'ff-openstack-haproxy = playback.haproxy:main',
           'ff-openstack-rabbitmq = playback.rabbitmq:main',
           'ff-openstack-keystone = playback.keystone:main',
           'ff-openstack-glance = playback.glance:main',
           'ff-openstack-nova = playback.nova:main',
           'ff-openstack-nova-compute = playback.nova_compute:main',
           'ff-openstack-neutron = playback.neutron:main',
           'ff-openstack-neutron-agent = playback.neutron_agent:main',
           'ff-openstack-horizon = playback.horizon:main',
           'ff-openstack-cinder = playback.cinder:main',
           'ff-openstack-swift = playback.swift:main',
           'ff-openstack-swift-storage = playback.swift_storage:main',
           'ff-openstack-ceph-deploy = ceph_deploy.cli:main'
           ],
       },
    )
