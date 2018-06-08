import logging

import psycopg2

import config
import util


log = logging.getLogger(__name__)


def chat_tests(env):
    _setup_chat_env(env)
    _clear_db()
    users = [user for user in env['directory']['users'].values()]
    delete_users(users)
    create_users(users)
    channel = create_channel(env['chat']['game'], users[:-1])
    env['chat']['channels'][channel['id']] = channel
    subscribe_user_to_channel(channel, users[-1])
    send_messages(channel, users[0], users[-1])
    test_get_channels(channel, env['chat']['game'], users[0]['token'])
    test_get_messages(channel, users[0]['token'])
    log.info('Chat: OK')


def create_users(users):
    """Creates directory users in chat db

    :param users: List of users to create.
    """
    tokens = [user['token'] for user in users]
    for token in tokens:
        util.post_request(config.CHAT_USER_ROUTE, util.make_headers(token))


def delete_users(users):
    """Deletes directory users by first creating them.

    :param users: List of users to create.
    """
    create_users(users)
    tokens = [user['token'] for user in users]
    for token in tokens:
        util.delete_request(config.CHAT_USER_ROUTE, util.make_headers(token))


def create_channel(game_id, recievers):
    token = recievers[0]['token']
    channel_route = config.CREATE_CHANNEL_ROUTE.format(game_id)
    query = util.create_query('recieverId', [r['user']['id'] for r in recievers])
    channel = util.post_request(f'{channel_route}?{query}', util.make_headers(token))
    return channel


def subscribe_user_to_channel(channel, user):
    URL = config.CHANNEL_ROUTE.format(channel['id'])
    util.put_request(URL, util.make_headers(user['token']))


def send_messages(channel, user_1, user_2):
    use_first_user = True
    URL = config.MESSAGE_ROUTE.format(channel['id'])
    for msg in config.MELIAN_DIALOGE:
        token = user_1['token'] if use_first_user else user_2['token']
        util.post_request(URL, util.make_headers(token), _make_chat_msg(msg))
        use_first_user = not use_first_user


def test_get_channels(channel, game_id, token):
    URL = config.GET_CHANNEL_ROUTE.format(game_id)
    recieved_channels = util.get_request(URL, util.make_headers(token))
    if len(recieved_channels) != 1:
        raise ValueError('Wrong number of channels: {}'.format(len(recieved_channels)))
    if recieved_channels[0]['id'] != channel['id']:
        raise ValueError('Wrong channel returned: {}'.format(channel['id']))


def test_get_messages(channel, token):
    URL = config.MESSAGE_ROUTE.format(channel['id'])
    messages = util.get_request(URL, util.make_headers(token))
    if len(messages) != len(config.MELIAN_DIALOGE):
        raise ValueError('Wrong number of messages recieved: {}'.format(len(messages)))
    for i in range(len(messages)):
        expected = config.MELIAN_DIALOGE[i]
        actual = messages[i]
        log.debug('Message Id: {}'.format(actual['id']))
        log.debug(actual['text'])
        if actual['text'] != expected:
            raise ValueError('Wrong message text for message: {}'.format(actual['id']))


def _make_chat_msg(text):
    return {'text': text}


def _setup_chat_env(env):
    env['chat'] = {
        'game': 'game-1',
        'channels': {}
    }


def _clear_db():
    """Clears the data in chat tables"""
    try:
        conn = psycopg2.connect(config.CHAT_CONN_STR)
        for table in config.CHAT_TABLES:
            util.clear_table(conn, table)
    except Exception as e:
        log.error(str(e))
        raise e
