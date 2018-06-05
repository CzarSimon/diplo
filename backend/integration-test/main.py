import json

import config
import directory
import chat


def main(print_env):
    env = {}
    directory.directory_tests(env)
    chat.chat_tests(env)
    if print_env:
        print(json.dumps(env, sort_keys=True, indent=4))


if __name__ == '__main__':
    main(False)
