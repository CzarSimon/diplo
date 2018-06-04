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
