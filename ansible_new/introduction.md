1.Create a Role: First, let's create a role named file_copy using ansible-galaxy:
ansible-galaxy init file_copy

2.Edit the Role Tasks: Navigate into the file_copy role directory and edit the tasks/main.yml file:
```
---
- name: Copy file from /home/oviner/test.txt to /home/oviner/ansible/
  copy:
    src: /home/oviner/test.txt
    dest: /home/oviner/ansible/test.txt
```

This task uses the copy module to copy the test.txt file from /home/oviner/ to /home/oviner/ansible/.

3.Create a Playbook: Now, let's create a playbook named copy_file.yml:
```
---
- name: Copy File Playbook
  hosts: all
  become: true
  roles:
    - file_copy
```

name: Name of the playbook.
hosts: Targets for this playbook (you can specify specific hosts or groups here).
become: Whether to execute tasks with elevated privileges (like sudo).
roles: The role(s) to include in this playbook. In this case, it includes the file_copy role.

4.Inventory Setup: Create an inventory file named inventory.ini:

```
# inventory
[servers]
server1 ansible_host=127.0.0.1
```

5.Run the Playbook: Run the playbook using the ansible-playbook command:

ansible-playbook -i inventory.ini copy_file.yml