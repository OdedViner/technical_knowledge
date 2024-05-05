**Resolve Conflict:**
1.Conflict message on git server (github)

2.`git fetch --all`

3.`git rebase upstream/master`  
output:CONFLICT (content): Merge conflict in setup.py  
4.Fix issues in setup.py    
(====> head  
<====)  

5.`git add <file>`

6.`git rebase --continue`

7.`git status` (without changes)

8.`git push --force`