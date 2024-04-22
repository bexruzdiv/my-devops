from itertools import chain


class FilterModule(object):
    @staticmethod
    def _to_list(arg=None):
        if arg is None:
            arg = []
        if not isinstance(arg, list):
            arg = [arg]
        return list(chain(arg))

    @staticmethod
    def _list_format(elements, format_string, *args):
        if elements is None:
            return []
        if not isinstance(elements, list):
            elements = [elements]
        return [
            format_string.format(element, *args) for element in list(chain(elements))
        ]

    def filters(self):
        return {
            'to_list': self._to_list,
            'list_format': self._list_format,
        }
