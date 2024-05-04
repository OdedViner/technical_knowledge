
**Ansible modules** 
Ansible modules are standalone scripts that Ansible uses to perform specific tasks on managed nodes.  
These modules are the building blocks of Ansible automation, allowing users to perform a wide range of actions   
such as executing commands, managing files, installing software, configuring services, and more on remote systems.

**Plugins** 
plugins are pieces of code that extend or enhance the functionality of Ansible itself.  
They allow users to customize and extend Ansible's core capabilities to suit their specific needs.  
Plugins can be used to add new features, modify existing behaviors, or integrate Ansible with external systems or services.
Exapmple:  
Inventory Plugins: These plugins are responsible for generating the inventory of hosts that Ansible will manage.    
They can pull inventory information from various sources such as static files, dynamic cloud platforms, databases, or external inventory systems.

**Handlers**  
handlers are tasks that are triggered by other tasks and are only executed if the notified task changes something on the managed nodes.   
They are typically used to perform actions like restarting a service, reloading a configuration file.

Example:  
```
---
- name: Example Playbook
  hosts: servers
  tasks:
    - name: Ensure the Nginx package is installed
      yum:
        name: nginx
        state: present

    - name: Ensure the Nginx service is started
      service:
        name: nginx
        state: started
      notify: restart nginx  # Notify the 'restart nginx' handler if the service is started
                    ^
  handlers:         |
    - name: restart nginx
      service:
        name: nginx
        state: restarted
```

**Process:**  
1.Create a Role: First, let's create a role named file_copy using ansible-galaxy:

```$ ansible-galaxy init file_copy```

2.Edit the Role Tasks: Navigate into the file_copy role directory and edit the tasks/main.yml file:
```
---
- name: Copy file from /home/oviner/test.txt to /home/oviner/ansible/
  copy:
    src: /home/oviner/test.txt
    dest: /home/oviner/ansible/test.txt
``` 
**This task uses the copy module to copy the test.txt file from /home/oviner/ to /home/oviner/ansible/.

3.Create a Playbook named copy_file.yml:
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

`$ ansible-playbook -i inventory.ini copy_file.yml`