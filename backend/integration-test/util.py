import json

import requests
import psycopg2

from config import BASE_HEADERS


def make_headers(token=None):
    """Makes request headers.

    :param token: Authorization token.
    :return: Request headers as dict
    """
    if token is None:
        return  BASE_HEADERS
    headers = BASE_HEADERS
    headers['Authorization'] = f'Bearer {token}'
    return headers


def get_request(URL, headers):
    """Makes a get request, checks and returns the response.

    :param URL: url to make the request to.
    :param headers: Header dict to include in the request.
    :return: Deserialized JSON body of the response.
    """
    resp = requests.get(URL, headers=headers)
    check_response(resp)
    return resp.json()


def post_request(URL, headers, body=None):
    """Makes a post request, checks and returns the response.

    :param URL: url to make the request to.
    :param headers: Header dict to include in the request.
    :param body: Body to include in the request.
    :return: Deserialized JSON body of the response.
    """
    resp = requests.post(URL, headers=headers, data=json.dumps(body))
    check_response(resp)
    return resp.json()


def put_request(URL, headers, body=None):
    """Makes a put request, checks and returns the response.

    :param URL: url to make the request to.
    :param headers: Header dict to include in the request.
    :param body: Body to include in the request.
    :return: Deserialized JSON body of the response.
    """
    resp = requests.put(URL, headers=headers, data=json.dumps(body))
    check_response(resp)
    return resp.json()


def delete_request(URL, headers):
    """Makes a delete request, checks and returns the response.

    :param URL: url to make the request to.
    :param headers: Header dict to include in the request.
    :return: Deserialized JSON body of the response.
    """
    resp = requests.delete(URL, headers=headers)
    check_response(resp)
    return resp.json()


def check_response(response):
    """Checks that a response was successfull, raises an exceptions if not.

    :param response: Response from a request.
    """
    status = response.status_code
    UNAUTHORIZED_CODES = [401, 403]
    if status in UNAUTHORIZED_CODES:
        print('Unauthorized to make request. Status: {}'.format(status))
    response.raise_for_status()


def clear_table(conn, table_name):
    """Clears the contents of a table.

    :param conn: Database connection to use.
    :param table_name: Name of the table to clear.
    """
    SQL = f'DELETE FROM {table_name}'
    try:
        cur = conn.cursor()
        cur.execute(SQL)
        conn.commit()
        cur.close()
    except Exception as e:
        print(str(e))
        conn.rollback()


def create_query(key, values):
    return '&'.join([f'{key}={value}' for value in values])
