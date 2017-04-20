type hostPathProvisioner struct {
  // The directory to create PV-backing directories in
  pvDir string

  // Identity of this hostPathProvisioner, set to node's name. Used to identify
  // "this" provisioner's PVs.
  identity string
}

func NewHostPathProvisioner() controller.Provisioner {
  nodeName := os.Getenv("NODE_NAME")
  if nodeName == "" {
    glog.Fatal("env variable NODE_NAME must be set so that this provisioner can identify itself")
  }
  return &hostPathProvisioner{
    pvDir:    "/tmp/hostpath-provisioner",
    identity: nodeName,
  }
}