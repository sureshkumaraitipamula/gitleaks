package config

// DefaultConfig is the default gitleaks configuration. If --config={path-to-config} is set than the config located
// at {path-to-config} will be used. Alternatively, if --repo-config is set then gitleaks will attempt to
// use the config set in a gitleaks.toml or .gitleaks.toml file in the repo that is run with --repo-config set.
const DefaultConfig = `
title = "gitleaks config"

[[rules]]
    description = "AWS Access Key"
    regex = '''(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'''
    tags = ["key", "AWS"]


[allowlist]
    description = "Allowlisted files"
    files = ['''^\.?gitleaks.toml$''',
    '''(.*?)(png|jpg|gif|doc|docx|pdf|bin|xls|pyc|zip)$''',
    '''(go.mod|go.sum)$''']
`
