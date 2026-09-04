# Repo Sync

Repo-Sync is divided into a pusher and puller design. 

The puller pulls Git repositories and bundles them into a specific structure on the filesystem, optionally compressing the output.

The pusher reads a filesystem for bundles and pushes them into Git repositories, optionally decompressing the input.

## Structure of pulls

The output of a pull-run is as follows:

```txt
repos.json
bundles/
  repo-name-1/
    0.bundle
    1.bundle
    ...
    100.bundle
  repo-name-2/
    0.bundle
    1.bundle
    ...
    77.bundle
  ...
  repo-name-100/
    0.bundle
```

Bundles are organized per repository, and are numbered in creation order for easier maintenance.

This allows all outputs to be combined into a single folder, and also allows the other end (push-run) to trivially validate bundle order.

The structure of repos.json is as follows:

```json
{
  "repositories": [
    {
      "name": "repo-name-1", // Every repository must have a unique name
      "srcUrl": "https://github.com/...", // URL of the repository to pull from
      "dstUrl": "https://github.com/...", // URL of the repository to push to
      "bundles": [
        {
          "order": 0, // Every bundle must have a unique order number per repository
          "fizesize": 123456, // Filesize of the bundle in bytes
          "sha256": "abcdef", // SHA256 sum of the bundle
          "branches": [ // Every single branch is saved
            {
              "name": "branch_name_1",
              "lastRev": "abcdef", // Hash of the last commit
              "curRev": "fedcba" // Hash of the current commit
            }
          ]
        }
      ]
    }
  ]
}
```