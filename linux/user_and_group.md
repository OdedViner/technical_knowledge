## Users:  

**User Account:**   
A user account represents an individual or a system entity with a unique username and user ID (UID). 
User accounts are used to log in to the system and perform various tasks.

**Username:**   
A username is a human-readable identifier associated with a user account. It is used for authentication during the login process.

**User ID (UID):**   
Each user account is associated with a numeric user ID (UID), which is a unique identifier for the user. 
The UID is used internally by the system for user management.

**Home Directory:**   
Each user typically has a home directory where they store their files and configuration settings. 
The home directory is usually located in /home/username.

**Shell:** 
A shell is a command-line interface that allows users to interact with the system. 
Users are assigned a default shell when their accounts are created.

**User Groups:**  
Users can belong to one or more user groups, which determine their access permissions to files, directories, and resources. 
By default, users are often members of a primary group with the same name as their username.

**Password:**   
Users have passwords associated with their accounts to authenticate themselves during login. 
Passwords are stored securely as hashes in the /etc/shadow file.

## Groups:  
**Group Account:**  
A group account is a collection of user accounts that share common permissions and access rights to files and resources. 
Group accounts are identified by a unique group name and group ID (GID).

**Group Name:**   
A group name is a human-readable identifier for a group account. 
It is used to associate multiple users with the same group.

**Group ID (GID):**  
Each group account is associated with a numeric group ID (GID), which is a unique identifier for the group. 
The GID is used internally by the system for group management.

**Primary Group:**   
Each user account is typically associated with a primary group, which shares the same name as the user's username. 
This primary group is set in the user's /etc/passwd entry.

**Secondary Groups:**   
Users can belong to multiple secondary groups in addition to their primary group. 
Secondary groups are defined in the /etc/group file.
