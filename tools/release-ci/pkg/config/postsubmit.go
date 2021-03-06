package config

import "fmt"

// Postsubmit contains configuration for a post-submit job
type Postsubmit struct {
	Job
	Repo
}

// GcsPath is the base path for job data. Postsubmit jobs
// store their information only in the job log directory
func (p *Postsubmit) GcsPath() string {
	return fmt.Sprintf("%s%s/%s/%d/", JobLogPrefix, repositoryPath(p.RepoOwner, p.RepoName), p.JobName, p.BuildNumber)
}

// Aliases are the paths to symlink to the GcsPath.
// Postsubmit jobs have no aliases.
func (p *Postsubmit) Aliases() []string {
	return []string{}
}

// Type exposes the type of this configuration
func (p *Postsubmit) Type() Type {
	return PostsubmitType
}
