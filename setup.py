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
    license='GPLv2',
    install_requires=['playback == 0.3.4'],
    packages=find_packages(),
    entry_points={ 
       'console_scripts': [
           'env-deploy = playback.env:main',
           'mysql-deploy = playback.mysql:main',
           'haproxy-deploy = playback.haproxy:main',
           'rabbitmq-deploy = playback.rabbitmq:main',
           'keystone-deploy = playback.keystone:main',
           'glance-deploy = playback.glance:main',
           'nova-deploy = playback.nova:main',
           'nova-compute-deploy = playback.nova_compute:main',
           'neutron-deploy = playback.neutron:main',
           'neutron-agent-deploy = playback.neutron_agent:main',
           'horizon-deploy = playback.horizon:main',
           'cinder-deploy = playback.cinder:main',
           'swift-deploy = playback.swift:main',
           'swift-storage-deploy = playback.swift_storage:main',
           'manila-deploy = playback.manila:main',
           'manila-share-deploy = playback.manila_share:main'
           ],
       },
    )
