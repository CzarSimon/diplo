import re

import requests


_DIALOGE_URL = 'https://www.mtholyoke.edu/acad/intrel/melian.htm'


def _get_html():
    resp = requests.get(_DIALOGE_URL)
    return resp.text


def _get_messages(text):
    pattern = re.compile('</strong>(.*?)<br>', re.DOTALL |  re.IGNORECASE)
    messages = [group[2:] for group in pattern.findall(text)]
    return _format_messages(messages)


def _format_messages(messages):
    no_newlines = [msg.replace('\r\n', ' ') for msg in messages]
    return [msg.replace('  ', ' ') for msg in no_newlines]


def main():
    text = _get_html()
    messages = _get_messages(text)
    print(messages)


if __name__ == '__main__':
    main()
