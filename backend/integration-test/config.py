import logging
from logging.config import dictConfig
import os


DIRECTORY_CONN_STR = "dbname='{}' user='{}' host='{}' password='{}'".format(
    'directory',
    os.getenv('DIPLO_DIRECTORY_DB_USER'),
    os.getenv('DIPLO_DIRECTORY_DB_HOST'),
    os.getenv('DIPLO_DIRECTORY_DB_PASSWORD'))


CHAT_CONN_STR = "dbname='{}' user='{}' host='{}' password='{}'".format(
    'chat',
    os.getenv('DIPLO_CHAT_DB_USER'),
    os.getenv('DIPLO_CHAT_DB_HOST'),
    os.getenv('DIPLO_CHAT_DB_PASSWORD'))


DIRECTORY_TABLES = [
    'account'
]


BASE_HEADERS = {
    'Content-Type': 'application/json'
}


USER_1 = {
    "email": "metternic@gmail.com",
    "username": "mr.austria",
    "password": "password",
    "givenName": "Clemens",
    "surname": "Von Metternic"
}


USER_2 = {
    "email": "castelraegh@gmail.com",
    "username": "mr.england",
    "password": "password",
    "givenName": "Robert Stewart",
    "surname": "Castelraegh"
}


BASE_URL = 'http://localhost:8080/api'


CREATE_USER_ROUTE = f'{BASE_URL}/directory/user'
LOGIN_USER_ROUTE = f'{BASE_URL}/directory/user/login'
RENEW_TOKEN_ROUTE = f'{BASE_URL}/directory/user/token/renew'

GET_USER_ROUTE = f'{BASE_URL}/directory/me'
GET_USERS_ROUTE = f'{BASE_URL}/directory/users'


LOGGING_CONIFG = {
    'version': 1,
    'formatters': {
        'default': {
            'format': '%(asctime)s - %(name)s - %(levelname)s - %(message)s'
        }
    },
    'handlers': {
        'console': {
            'class': 'logging.StreamHandler',
            'formatter': 'default',
            'level': logging.INFO
        }
    },
    'root': {
        'handlers': ['console'],
        'level': logging.DEBUG
    }
}


dictConfig(LOGGING_CONIFG)
