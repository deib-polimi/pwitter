import ConfigParser as cp
import os

CONFIG_FILE = 'pwitter.conf'
_CONF = cp.ConfigParser()
_CONF.readfp(open(CONFIG_FILE))

def get(opt, section='DEFAULT'):
    return _CONF.get(section, opt, vars=os.environ)
