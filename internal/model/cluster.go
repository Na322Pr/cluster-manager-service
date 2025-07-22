package model

import (
	"github.com/hashicorp/nomad/api"
	"time"
)

type Cluster struct {
	job     NomadJob
	version int
}

type NomadJob struct {
	ID          string
	Name        string
	Type        string
	Datacenters []string
	TaskGroups  []TaskGroup
}

type TaskGroup struct {
	Name    string
	Count   int
	Network Network
	Tasks   []Task
}

type Network struct {
	Ports []Port
}

type Port struct {
	Label string
	To    int
}

type Task struct {
	Name      string
	Driver    string
	Config    map[string]interface{}
	Env       map[string]string
	Resources Resources
	Services  []Service
}

type Resources struct {
	CPU    int
	Memory int
}

type Service struct {
	Name      string
	PortLabel string
	Checks    []ServiceCheck
}

type ServiceCheck struct {
	Name        string
	Type        string
	Path        string
	Interval    time.Duration
	Timeout     time.Duration
	GrpcService string
	GrpcUseTLS  bool
	PortLabel   string
}

func (j *NomadJob) ConvertToNomadJob() *api.Job {
	nomadJob := &api.Job{
		ID:          stringToPtr(j.ID),
		Name:        stringToPtr(j.Name),
		Type:        stringToPtr(j.Type),
		Datacenters: j.Datacenters,
	}

	// Convert TaskGroups
	var nomadTaskGroups []*api.TaskGroup
	for _, tg := range j.TaskGroups {
		nomadTG := &api.TaskGroup{
			Name:  stringToPtr(tg.Name),
			Count: intToPtr(tg.Count),
		}

		// Convert Network if it exists
		nomadNetwork := &api.NetworkResource{
			DynamicPorts: []api.Port{},
		}

		for _, port := range tg.Network.Ports {
			nomadPort := api.Port{
				Label: port.Label,
				To:    port.To,
			}
			nomadNetwork.DynamicPorts = append(nomadNetwork.DynamicPorts, nomadPort)
		}
		nomadTG.Networks = []*api.NetworkResource{nomadNetwork}

		// Convert Tasks
		var nomadTasks []*api.Task
		for _, t := range tg.Tasks {
			nomadTask := &api.Task{
				Name:   t.Name,
				Driver: t.Driver,
				Config: t.Config,
				Env:    t.Env,
				Resources: &api.Resources{
					CPU:      intToPtr(t.Resources.CPU),
					MemoryMB: intToPtr(t.Resources.Memory),
				},
			}

			// Convert Services
			var nomadServices []*api.Service
			for _, s := range t.Services {
				nomadService := &api.Service{
					Name:      s.Name,
					PortLabel: s.PortLabel,
				}

				// Convert ServiceChecks with gRPC support
				var nomadChecks []api.ServiceCheck
				for _, sc := range s.Checks {
					nomadCheck := api.ServiceCheck{
						Name:        sc.Name,
						Type:        sc.Type,
						Path:        sc.Path,
						Interval:    sc.Interval,
						Timeout:     sc.Timeout,
						PortLabel:   sc.PortLabel,
						GRPCService: sc.GrpcService,
						GRPCUseTLS:  sc.GrpcUseTLS,
					}
					nomadChecks = append(nomadChecks, nomadCheck)
				}
				nomadService.Checks = nomadChecks
				nomadServices = append(nomadServices, nomadService)
			}
			nomadTask.Services = nomadServices
			nomadTasks = append(nomadTasks, nomadTask)
		}

		nomadTG.Tasks = nomadTasks
		nomadTaskGroups = append(nomadTaskGroups, nomadTG)
	}

	nomadJob.TaskGroups = nomadTaskGroups
	return nomadJob
}

func stringToPtr(s string) *string { return &s }
func intToPtr(i int) *int          { return &i }

func (c *Cluster) SetClusterSize(size int) {
	c.job.TaskGroups[0].Count = size
	c.version += 1
}

func (c *Cluster) GetNomadJob() *api.Job {
	return c.job.ConvertToNomadJob()
}

func (c *Cluster) GetJobID() string {
	return c.job.ID
}

func (c *Cluster) GetTaskGroupName() string {
	return c.job.TaskGroups[0].Name
}

func (c *Cluster) GetClusterSize() *int {
	return intToPtr(c.job.TaskGroups[0].Count)
}

func NewSampleCluster() *Cluster {
	return &Cluster{
		job: NomadJob{
			ID:          "go-service",
			Name:        "go-service",
			Type:        "service",
			Datacenters: []string{"dc1"},
			TaskGroups: []TaskGroup{
				{
					Name:  "app",
					Count: 3,
					Network: Network{
						Ports: []Port{
							{
								Label: "http",
								To:    7000,
							},
							{
								Label: "grpc",
								To:    7001,
							},
						},
					},
					Tasks: []Task{
						{
							Name:   "server",
							Driver: "docker",
							Config: map[string]interface{}{
								"image": "na322pr/kv-storage-service:latest",
								"ports": []string{"grpc"},
							},
							Env: map[string]string{
								"NODE_ID":    "${NOMAD_ALLOC_INDEX}",
								"GRPC_PORT":  "${NOMAD_PORT_grpc}",
								"HTTP_PORT":  "${NOMAD_PORT_http}",
								"SEED_NODES": "go-service-0.service.consul:7001,go-service-1.service.consul:7001",
							},
							Resources: Resources{
								CPU:    5,
								Memory: 64,
							},
							Services: []Service{
								{
									Name:      "go-service",
									PortLabel: "grpc",
									Checks: []ServiceCheck{
										{
											Name:        "grpc-health-check",
											Type:        "grpc",
											PortLabel:   "grpc",
											Interval:    15 * time.Second,
											Timeout:     5 * time.Second,
											GrpcService: "kv_storage_service.KeyValueStorage",
											GrpcUseTLS:  false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		version: 1,
	}
}
