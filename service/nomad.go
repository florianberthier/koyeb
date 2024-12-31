package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/nomad/api"
)

func (s *Service) createNomadJob(name string, script bool, content []byte) (*api.Job, error) {
	taskName := "nginx"

	tmpDir, err := os.MkdirTemp("", "nginx-content")
	if err != nil {
		return nil, fmt.Errorf("error creating temp dir: %v", err)
	}

	contentPath := ""
	volume := fmt.Sprintf("%s:/usr/share/nginx/html", tmpDir)
	perm := os.FileMode(0644)
	if script {
		contentPath = filepath.Join(tmpDir, "index.sh")
		perm = os.FileMode(0755)
		volume = fmt.Sprintf("%s:/tmp/index.sh", contentPath)
	} else {
		contentPath = filepath.Join(tmpDir, "index.html")
	}

	err = os.WriteFile(contentPath, content, perm)
	if err != nil {
		return nil, fmt.Errorf("error writing content to file: %v", err)
	}

	cpu := 500
	memory := 256

	job := &api.Job{
		ID:   &name,
		Name: &name,
		TaskGroups: []*api.TaskGroup{
			{
				Name: &taskName,
				Networks: []*api.NetworkResource{
					{
						ReservedPorts: []api.Port{
							{
								Label: "http",
								Value: 20676,
								To:    80,
							},
						},
					},
				},
				Tasks: []*api.Task{
					{
						Name:   taskName,
						Driver: "docker",
						Config: map[string]interface{}{
							"image": "nginx:latest",
							"ports": []string{"http"},
							"volumes": []string{
								volume,
							},
						},

						Resources: &api.Resources{
							CPU:      &cpu,
							MemoryMB: &memory,
						},
					},
				},
			},
		},
	}

	if script {
		job.TaskGroups[0].Tasks[0].Config["command"] = "sh"
		job.TaskGroups[0].Tasks[0].Config["args"] = []string{
			"-c",
			"chmod +x /tmp/index.sh && /bin/sh /tmp/index.sh > /usr/share/nginx/html/index.html && chmod 644 /usr/share/nginx/html/index.html && nginx -g 'daemon off;'",
		}
	}
	return job, nil
}
