# APTLIFY [![Build Status](https://travis-ci.org/queeno/aptlify.svg?branch=master)](https://travis-ci.org/queeno/aptlify)

Aptlify is a tool to drive aptly through a JSON configuration file.

## Commands

Plan - Create an change plan
Apply - Apply the staged changes
Mirror - Update the mirrors
Dump - Show the current aptly state

## How to use

Create a configuration file in `$HOME/.aptlify.conf`

The file should include 4 sections: `gpg_keys`, `mirrors`, `repos`, `snapshots`

  - `gpg_keys` should contain gpg public key fingerprints within an array of `fingerprint`.
    The keys will be downloaded from `keys.gnupg.net` and added to `trustedkeys.gpg`
  - `mirrors` should contain an array of aptly mirror definitions, including the following
    mandatory fields:
    - `name`: The name of the mirrors
    - `url`: The mirror upstream URL
    - `dist`: The upstream mirror distribution
    `mirrors` can also contain the following optional fields:
    - `component`: The upstream mirror component
    - `filter`: An array of filters which describe the package names and versions that will be filtered. The array must include:
      - `name`: The name of the package to fetch from the mirror.
        `version`: The version of the package to fetch from the mirror.
    - `filter-with-deps`: Whether to pass the `filter-with-deps` flag to the command to fetch the filtered package dependencies.
  - `repos` should contain an array of aptly repo definitions:
      - `name`: The name of the repo to create.
  - `snapshots`: should contain a list snapshot to take from repos and mirrors.
    The resulting snapshot will contain all the packages within the declared repos and mirrors.
    - `name`: The name of the combined snapshot to create.
      `resources`: An array of mirrors or repos to include in the snapshot.
      - `name`: The name of the resource to include in the snapshot.
        `type`: The type of the resource (either mirror or repo)
        `filter`: The filter to apply to the resource. If no filter is specified, the whole content of the resource is included in the snapshot.
        - `name`: The name of the package to filter.
          `version`: The version of the package to filter.

### Example `~/.aptlify.conf`

```
{
  "gpg_keys":
    {
      "fingerprint": [
        "9E3E53F19C7DE460",
        "353525F9",
        "505D97A41C61B9CD",
        "1C61B9CD"
      ]
    },
  "mirrors": [
    {
      "name": "hmrc-aptly",
      "url": "http://repo.aptly.info",
      "dist": "squeeze",
      "component": "main",
      "filter": [
        {
          "name": "aptly",
          "version": "0.9.7"
        }
      ]
    },
    {
      "name": "haproxy",
      "url": "http://ppa.launchpad.net/vbernat/haproxy-1.5/ubuntu",
      "dist": "trusty"
    }
  ],
  "repos": [
    {
      "name": "hmrc-internal"
    },
    {
      "name": "test"
    }
  ],
  "snapshots": [
    {
      "name": "hmrc",
      "resources": [
        {
          "type": "mirror",
          "name": "haproxy",
          "filter": [
            { "name": "haproxy", "version": "1.5" }
          ]
        }
      ]
    }
  ]
}
```
