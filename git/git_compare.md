
**Compare code btween areas:**
```
woking dir -> staging area -> local repository -> remote
```

**Compare working dir and staging area:**  
`git diff`

**Compare working dir and local repository:**  
`git diff HEAD`

**Compare staging area and local repository:**  
`git diff --staged HEAD`

**Limiting Comparisons to one File (or Path)**  
`git diff -- README.md`

**Comparing Between Commits:**
```
$ git log --oneline 
866adce (HEAD -> main, origin/main) ansible new
32903d7 ansible new
554fa95 add terrraform
3bb7599 add info on fio
c2426c5 add storage_performance_testing.md
```
```
$ git diff c2426c5 HEAD
```
```
$ git diff c2426c5 3bb7599
```

**Comparing Between Local and Remote Master Branches:**  
`git diff master origin/master`

