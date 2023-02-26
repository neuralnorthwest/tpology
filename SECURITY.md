# tpology Security Policy

We want tpology to be safe and secure for everyone. We do this by following
these practices:

- **Dependencies**: Part of tpology's release process is to update all of its
dependencies to the latest versions, including transitive dependencies.
Additionally, we apply Dependabot to the repository, and if a security
vulnerability is found in a dependency, we will update the dependency as soon
as possible and release a new version of tpology.
- **CodeQL**: We use CodeQL to scan the codebase for security vulnerabilities.
Vulnerabilities of level HIGH or above are blockers to a release, and we will
address these before releasing any new version of tpology.

## Supported Versions

The tpology team strives to maintain backward compatibility within major
versions. This means that within a given major version, we do not backport bug
fixes or security fixes. However, we will backport security fixes to the
previous major version if possible. Versions older than the previous major
version are not supported.

## Reporting a Vulnerability

Please report vulnerabilities to
[neuralnorthwestllc+vuln@gmail.com](mailto:neuralnorthwestllc+vuln@gmail.com).
