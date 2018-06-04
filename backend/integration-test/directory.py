import logging

import psycopg2

import config
import util


log = logging.getLogger(__name__)


def directory_tests(env):
    _setup_directory_env(env)
    _clear_db()
    env['directory']['users']['user-1'] = create_user(config.USER_1)
    env['directory']['users']['user-2'] = create_user(config.USER_2)
    user_1_token = login_user(config.USER_1)
    user_1_token = renew_token(user_1_token)
    env['directory']['users']['user-1']['token'] = user_1_token
    test_get_me(env)
    test_get_users(env)
    log.info('Directory: OK')


def create_user(user):
    resp = util.post_request(
        config.CREATE_USER_ROUTE, util.make_headers(), user)
    validate_token(resp['token'])
    return resp


def login_user(user):
    resp = util.post_request(
        config.LOGIN_USER_ROUTE, util.make_headers(), user)
    validate_token(resp['token'])
    return resp['token']


def renew_token(token):
    resp = util.get_request(config.RENEW_TOKEN_ROUTE, util.make_headers(token))
    validate_token(resp['token'])
    return resp['token']


def test_get_me(env):
    user_1 = env['directory']['users']['user-1']['user']
    token = env['directory']['users']['user-1']['token']
    resp = util.get_request(config.GET_USER_ROUTE, util.make_headers(token))
    if resp['id'] != user_1['id']:
        raise ValueError('Wrong user id returned from: {}'.format(
            config.GET_USER_ROUTE))


def test_get_users(env):
    ids = [user['user']['id'] for user in env['directory']['users'].values()]
    get_users_route = '{}?{}'.format(
        config.GET_USERS_ROUTE, '&'.join([f'userId={id}' for id in ids]))
    token = env['directory']['users']['user-1']['token']
    users = util.get_request(get_users_route, util.make_headers(token))
    found_ids = set([user['id'] for user in users])
    for id in ids:
        if not id in found_ids:
            raise ValueError(f'Id: {id} not found')



def validate_token(token):
    token_parts = token.split('.')
    if len(token_parts) != 3:
        raise TypeError(f'Invalid token: {token}')


def _setup_directory_env(env):
    env['directory'] = {
        'users': {
            'user-1': None,
            'user-2': None
        }
    }

def _clear_db():
    try:
        conn = psycopg2.connect(config.DIRECTORY_CONN_STR)
        for table in config.DIRECTORY_TABLES:
            util.clear_table(conn, table)
    except Exception as e:
        log.error(str(e))
        raise e
