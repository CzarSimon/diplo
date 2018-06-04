import json

import config
import directory


def main():
    env = {}
    directory.directory_tests(env)
    print(json.dumps(env, sort_keys=True, indent=4))


if __name__ == '__main__':
    main()
