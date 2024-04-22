from crypt import crypt
import string
import random


class FilterModule(object):
    @staticmethod
    def _to_pwd(password, seed=None):
        if password is None or password == '' or password == '!' or password == '!!':
            return '!'
        if password.startswith('$6$'):
            return password
        letters = string.ascii_letters + string.digits + './'
        if seed is not None:
            random.seed(seed)
        salt = '$6$' + ''.join(random.choice(letters) for _char in range(16))
        return crypt(password, salt)

    def filters(self):
        return {
            'to_pwd': self._to_pwd
        }


filter = FilterModule()

print(filter._to_pwd("ieth7ach7mee~jaiFee9wahn", "324"))