Create a compressed tarball (.tar.gz) :

```
tar -czvf dir.tar.gz dir

-c option tells tar to create a new archive
-z specifies that you want to compress the archive using gzip 
-v enables verbose mode to see the progress
-f indicates the name of the archive file to be created.
```

Extract the contents of a tarball (.tar.gz or .tar): 
```
tar -xvf file.tar.gz
-x: Extract the contents of the archive.
-v: Enable verbose mode to see the list of files being extracted.
-f: Specifies the archive file to be extracted, in this cas
```


Example 'oc' client installation:
```
wget https://openshift-release-artifacts.apps.ci.l2s4.p1.openshiftapps.com/
tar -xvf  openshift-client-linux.tar.gz
sudo cp oc /usr/local/bin/
```