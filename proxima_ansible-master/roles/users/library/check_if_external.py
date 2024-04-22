#!/usr/bin/python

from ansible.module_utils.basic import AnsibleModule

from pwd import getpwnam
from grp import getgrnam

ANSIBLE_METADATA = {
    'metadata_version': '1.1',
    'status': ['preview'],
    'supported_by': 'community'
}

DOCUMENTATION = '''
---
module: check_if_external

short_description: Check if a user or a group is defined in a nss service other then local files.

version_added: "2.3"

description:
    - "Check if a user or a group is defined in a nss service other then local files"
    - "To determine if a local user or group should be created or not"

options:
    user:
        description:
            - User name to check
        required: false
    group:
        description:
            - Group name to check
        required: false
'''

EXAMPLES = '''
'''

RETURN = '''
user:
    description: User is defined not in local files
    type: bool
group:
    description: Group is defined not in local files
    type: bool
'''


def run_module():
    module_args = dict(
        user=dict(type='str', required=False, default=None),
        group=dict(type='str', required=False, default=None),
    )

    check_module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    user = check_module.params['user']
    group = check_module.params['group']

    result = dict(
        changed=False,
        user_name=user,
        user_external=False,
        group_name=group,
        group_external=False,
    )

    if not user and not group:
        check_module.fail_json(msg='You should provide a user, group or both!', **result)

    if user:
        if check_nss_user(user) and not check_files_user(user):
            result['user_external'] = True

    if group:
        if check_nss_group(group) and not check_files_group(group):
            result['group_external'] = True

    check_module.exit_json(**result)


def check_nss_user(name):
    """
    Args:
        name (str): Check if NSS query can find this user 
    Returns:
        bool
    """
    try:
        getpwnam(name)
        return True
    except KeyError:
        return False


def check_nss_group(name):
    """
    Args:
        name (str): Check if NSS query can find this group 
    Returns:
        bool
    """
    try:
        getgrnam(name)
        return True
    except KeyError:
        return False


def check_files_user(name):
    """
    Args:
        name (str): Check if this user is defined in /etc/passwd 
    Returns:
        bool
    """
    with open('/etc/passwd', 'r') as passwd:
        for line in passwd:
            if line.startswith(name + ':'):
                return True
    return False


def check_files_group(name):
    """
    Args:
        name (str): Check if this group is defined in /etc/group 
    Returns:
        bool
    """
    with open('/etc/group', 'r') as group:
        for line in group:
            if line.startswith(name + ':'):
                return True
    return False


def main():
    run_module()


if __name__ == '__main__':
    main()
