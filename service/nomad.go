package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/nomad/api"
)

func createNomadJob(name string, script bool, content []byte) (*api.Job, error) {
	taskName := "nginx"

	tmpDirPath := filepath.Join("/tmp", name)
	os.RemoveAll(tmpDirPath)
	err := os.Mkdir(tmpDirPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error creating temp dir: %v", err)
	}

	if err := os.Chmod(tmpDirPath, 0755); err != nil {
		return nil, fmt.Errorf("error chmod temp dir: %v", err)
	}

	contentPath := filepath.Join(tmpDirPath, "index.html")
	volume := tmpDirPath + ":/usr/share/nginx/html"
	perm := os.FileMode(0644)
	if script {
		contentPath = filepath.Join(tmpDirPath, "index.sh")
		perm = os.FileMode(0755)
		volume = fmt.Sprintf("%s:/tmp/index.sh", contentPath)
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
				Services: []*api.Service{
					{
						Name:      name,
						PortLabel: "http",
						Provider:  "nomad",
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
