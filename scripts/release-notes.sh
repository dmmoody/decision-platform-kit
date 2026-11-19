#!/usr/bin/env bash
# Print release notes (markdown) for commits since the latest v* tag.
set -euo pipefail

latest="$(git tag -l 'v*' --sort=-v:refname | head -n1 || true)"
if [[ -z "${latest}" ]]; then
  git log --format='- %s (%h)'
  exit 0
fi

git log --format='- %s (%h)' "${latest}..HEAD"
