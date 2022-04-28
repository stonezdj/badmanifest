# Badmanifest

Sometimes Harbor might contains bad manifest files in repository.

This tool is used to check and fix bad manifest files.

```
./badmanifest -location=/data/registry
```

Once you find the bad manifest file, you can fix with two approaches:

1. Delete the manifest file and push the artifact again.
2. Find the missed blob that is referenced by the manifest file and copy the data file from another repository.