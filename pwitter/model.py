from schematics.models import Model
from schematics.types import StringType, FloatType
import redis, uuid, json

from pwitter import config
from textblob import TextBlob

_r = redis.StrictRedis(
    host=config.get('DB_ADDRESS'),
    port=int(config.get('DB_PORT')),
    db=0
)

class Entity(Model):
    KEY_PREFIX = 'entity'

    def __init__(self, *args, **kwargs):
        Model.__init__(self, *args, **kwargs)
        self.ID = str(uuid.uuid1())

    @property
    def key(self):
        return ':'.join([self.KEY_PREFIX, str( self.ID )])

    def save(self):
        # using hashes
        # http://redis.io/commands#hash
        _r.hmset(self.key, self.to_primitive())
        self.on_save()

    def on_save(self):
        pass

    @classmethod
    def one(cls, key):
        raw = _r.hgetall(key)
        return cls(raw)

    @classmethod
    def _all(cls):
        # read at http://redis.io/commands/KEYS for
        # problems with KEYS command
        keys = _r.keys(cls.KEY_PREFIX + '*')
        return [cls.one(k) for k in keys]


class Pweet(Entity):
    '''
    We are using a redis Ordered Set to index keys by score.

    See http://redis.io/topics/data-types
    '''
    KEY_PREFIX = 'pweet'
    _SET_KEY = 'pweet-set'
    _LIMIT = 20

    user = StringType(default='anonymous')
    body = StringType(default='')
    polarity = FloatType(default=0.0)
    subjectivity = FloatType(default=0.0)

    def __init__(self, *args, **kwargs):
        Entity.__init__(self, *args, **kwargs)
        sentiment = TextBlob(self.body)
        self.polarity = sentiment.polarity
        self.subjectivity = sentiment.subjectivity

    def on_save(self):
        _r.zadd(
            self._SET_KEY,
            self.polarity, # the score
            self.key
        )

    @classmethod
    def count(cls):
        return _r.zcard(cls._SET_KEY)

    @classmethod
    def by_polarity(cls, low=-1.0, high=1.0):
        keys = _r.zrangebyscore(
            cls._SET_KEY, low, high,
            start=0, num=cls._LIMIT
        )
        return [cls.one(k) for k in keys]
