#!/usr/bin/env bash
# Print the next semver tag (vX.Y.Z) from commits since the latest v* tag.
# Bump rules use uppercase TYPE: commit prefixes (FEAT, FIX, …).
set -euo pipefail

latest="$(git tag -l 'v*' --sort=-v:refname | head -n1 || true)"
if [[ -z "${latest}" ]]; then
  latest="v0.0.0"
fi

ver="${latest#v}"
IFS=. read -r major minor patch <<<"${ver}"

if [[ "${latest}" == "v0.0.0" ]]; then
  log="$(git log --format='%s')"
else
  log="$(git log --format='%s' "${latest}..HEAD")"
fi

if [[ -z "${log}" ]]; then
  echo "no commits since ${latest}" >&2
  exit 2
fi

bump="patch"
while IFS= read -r subject; do
  [[ -z "${subject}" ]] && continue
  prefix="$(sed -n 's/^\([A-Z][A-Z0-9]*\):.*/\1/p' <<<"${subject}")"
  case "${prefix}" in
  BREAKING | MAJOR)
    bump="major"
    ;;
  FEAT | ARCH | REFACT)
    if [[ "${bump}" != "major" ]]; then
      bump="minor"
    fi
    ;;
  FIX | PERF | DOC | CHORE | TEST | STYLE | CI) ;;
  esac
done <<<"${log}"

# On v0, a "major" signal bumps minor instead of jumping to v1.
if [[ "${major}" == "0" && "${bump}" == "major" ]]; then
  bump="minor"
fi

case "${bump}" in
major)
  major=$((major + 1))
  minor=0
  patch=0
  ;;
minor)
  minor=$((minor + 1))
  patch=0
  ;;
patch)
  patch=$((patch + 1))
  ;;
esac

echo "v${major}.${minor}.${patch}"
