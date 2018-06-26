# Contributions

This project uses automated semantic versioning.  The technical aspects of that being that you are expected to submit pull requests for 
individual work items from a branch that is for that work item.  Find an issue; Create a branch "issue-#"; work and submit a 
pull request.

Version numbers are calculated automatically using symantec versioning:

* https://juhani.gitlab.io/go-semrel-gitlab/#introduction

Pull request messages are to be formatted using the Angular commit message format:

* https://juhani.gitlab.io/go-semrel-gitlab/#getting-started
* https://semantic-release.gitbooks.io/semantic-release/content/#how-does-it-work

| Type     | Description                                                                                            |
| -------- | ------------------------------------------------------------------------------------------------------ |
| feat     | A new feature                                                                                          |
| fix      | A bug fix                                                                                              |
| docs     | Documentation only changes                                                                             |
| style    | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc) |
| refactor | A code change that neither fixes a bug nor adds a feature                                              |
| perf     | A code change that improves performance                                                                |
| test     | Adding missing or correcting existing tests                                                            |
| chore    | Changes to the build process or auxiliary tools and libraries such as documentation generation         |

Examples:
```
fix(my service): did all the appropriate things.
```

```
feat(teh ui): now in color!  
```

```
chore(documentation): moved all documentation to new location.

all the things in detail down here!
```

```
perf(service): refactored obsolete functionality

BREAKING CHANGE: my service will no more do all of the appropriate things.

```