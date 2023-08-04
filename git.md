git init: Initializes a new Git repository in the current directory.

git clone <repository_url>: Copies an existing Git repository from a remote server 

git status: Shows the current status of your working directory.

git add <file>: Adds changes from your working directory to the staging area.

******************************************************************************************
```
git commit -m "commit message" -s 
```
Commits the changes in the staging area to the repository.
The -m flag is used to specify the commit message
The -s flag adds a "Signed-off-by" line to the commit message, indicating that the author is acknowledging 
certain legal and licensing requirements. The "Signed-off-by" line typically includes the author's name and email address.

******************************************************************************************
git pull: Fetches changes from the remote repository and merges them into your current branch. 

******************************************************************************************
git push: Sends your committed changes to the remote repository. It updates the remote branch with your local changes.
If you're on the "feature-branch" and want to push it to the remote repository:
git push origin feature-branch
******************************************************************************************

git fetch --all, Git fetches all branches from all remotes, updating your local repository with the 
latest changes from the remote repositories

git branch: Lists all the branches in your repository. 

git checkout <branch>: Switches to the specified branch. 
If the branch doesn't exist, you'll get an error. You can create a new branch using git checkout -b <new_branch>.

******************************************************************************************
git rebase: 
The git rebase command is used to apply a sequence of commits from one branch on top of another branch.

Switch to your feature branch that you want to rebase:
git checkout your-feature-branch

Rebase your feature branch onto the updated main branch:
git rebase main
******************************************************************************************

******************************************************************************************
git merge <branch>: Merges the specified branch into the current branch. 
It combines the changes from the specified branch into your current branch.

**Example**:
```
Switch to the target branch: 
git checkout master
# Merge the changes: Once you are on the target branch (master in this case):
git merge feature-branch
```
******************************************************************************************
git log: Shows the commit history, displaying the commit hashes, authors, dates, and commit messages.

git remote: Lists the remote repositories linked to your local repository. 
By default, "origin" is the name given to the remote repository you cloned from.


******************************************************************************************
******************************************************************************************
Git Conflict steps:
```
1.Conflict message on git server.

2.git fetch --all (from pgsql_debug branch)

3.git rebase upstream/master 
output:CONFLICT (content): Merge conflict in setup.py
4.Fix issues in setup.py 
(====> head
<====)

5.git add <file>

6.git rebase --continue

7.git status (without changes)

8.git push --force

```
******************************************************************************************
******************************************************************************************
Cherry-Pick:
[
     Git commands demonstrates how to cherry-pick a specific commit from the upstream/release-4.8 
     branch into a new branch named release-4.8-cherry-pick-pr-XYZ and then push that new branch to the 
     remote repository named origin. Here's a step-by-step breakdown of the commands:
]
```
Create and switch to a new branch named release-4.8-cherry-pick-pr-XYZ, which will be based on the upstream/release-4.8 branch:
git checkout -b release-4.8-cherry-pick-pr-XYZ upstream/release-4.8

Cherry-pick the desired commit (COMMIT_ID) from the upstream/release-4.8 branch onto the release-4.8-cherry-pick-pr-XYZ branch:
git cherry-pick COMMIT_ID

Push the newly created branch release-4.8-cherry-pick-pr-XYZ to the remote repository named origin:
git push origin release-4.8-cherry-pick-pr-XYZ
```