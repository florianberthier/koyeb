package service

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNomadJobNoScript(t *testing.T) {
	name := "test-job"
	content := []byte("<html><body>Test</body></html>")

	job, err := createNomadJob(name, false, content)

	require.NoError(t, err)
	assert.NotNil(t, job)
	assert.Equal(t, "test-job", *job.Name)
	assert.Equal(t, "test-job", *job.ID)

	assert.Len(t, job.TaskGroups, 1)
	taskGroup := job.TaskGroups[0]
	assert.Len(t, taskGroup.Tasks, 1)
	task := taskGroup.Tasks[0]

	expectedVolume := filepath.Join("/tmp", name) + ":/usr/share/nginx/html"
	fmt.Println(task.Config["volumes"])
	assert.Contains(t, task.Config["volumes"], expectedVolume)
}

func TestCreateNomadJobWithScript(t *testing.T) {
	name := "test-job-script"
	content := []byte("#!/bin/sh\necho 'Hello World!'")

	job, err := createNomadJob(name, true, content)

	require.NoError(t, err)
	assert.NotNil(t, job)
	assert.Equal(t, "test-job-script", *job.Name)
	assert.Equal(t, "test-job-script", *job.ID)

	assert.Len(t, job.TaskGroups, 1)
	taskGroup := job.TaskGroups[0]
	assert.Len(t, taskGroup.Tasks, 1)
	task := taskGroup.Tasks[0]

	args := task.Config["args"].([]string)
	assert.Equal(t, []string{
		"-c",
		"chmod +x /tmp/index.sh && /bin/sh /tmp/index.sh > /usr/share/nginx/html/index.html && chmod 644 /usr/share/nginx/html/index.html && nginx -g 'daemon off;'",
	}, args)
}
