% crio.conf(5) Kubernetes Container Runtime Daemon for Open Container Initiative Containers
% Aleksa Sarai
% OCTOBER 2016

# NAME

crio.conf - configuration file of the CRI-O OCI Kubernetes Container Runtime daemon

# DESCRIPTION

The CRI-O configuration file specifies all of the available configuration options and command-line flags for the [crio(8) OCI Kubernetes Container Runtime daemon][crio], but in a TOML format that can be more easily modified and versioned.

CRI-O supports partial configuration reload during runtime, which can be done by sending SIGHUP to the running process. Currently supported options in `crio.conf` are explicitly marked with 'This option supports live configuration reload'.

The containers-registries.conf(5) file can be reloaded as well by sending SIGHUP to the `crio` process.

The default crio.conf is located at /etc/crio/crio.conf.

# FORMAT

The [TOML format][toml] is used as the encoding of the configuration file. Every option and subtable listed here is nested under a global "crio" table. No bare options are used. The format of TOML can be simplified to:

    [table]
    option = value

    [table.subtable1]
    option = value

    [table.subtable2]
    option = value

## CRIO TABLE

CRI-O reads its storage defaults from the containers-storage.conf(5) file located at /etc/containers/storage.conf. Modify this storage configuration if you want to change the system's defaults. If you want to modify storage just for CRI-O, you can change the storage configuration options here.

**root**="/var/lib/containers/storage"
Path to the "root directory". CRI-O stores all of its data, including containers images, in this directory.

**runroot**="/var/run/containers/storage"
Path to the "run directory". CRI-O stores all of its state in this directory.

**storage_driver**="overlay"
Storage driver used to manage the storage of images and containers. Please refer to containers-storage.conf(5) to see all available storage drivers.

**storage_option**=[]
List to pass options to the storage driver. Please refer to containers-storage.conf(5) to see all available storage options.

**log_dir**="/var/log/crio/pods"
The default log directory where all logs will go unless directly specified by the kubelet. The log directory specified must be an absolute directory.

**version_file**="/var/run/crio/version"
Location for CRI-O to lay down the temporary version file.
It is used to check if crio wipe should wipe containers, which should
always happen on a node reboot

**version_file_persist**=""
Location for CRI-O to lay down the persistent version file.
It is used to check if crio wipe should wipe images, which should
only happen when CRI-O has been upgraded

**imagestore**=""
Store newly pulled images in the specified path, rather than the path provided by --root.

**internal_wipe**=true
**This option is currently DEPRECATED, and will be removed in the future.**
Whether CRI-O should wipe containers after a reboot and images after an upgrade when the server starts.
If set to false, one must run `crio wipe` to wipe the containers and images in these situations.

**internal_repair**=true
InternalRepair is whether CRI-O should check if the container and image storage was corrupted after a sudden restart.
If it was, CRI-O also attempts to repair the storage.

**clean_shutdown_file**="/var/lib/crio/clean.shutdown"
Location for CRI-O to lay down the clean shutdown file.
It is used to check whether crio had time to sync before shutting down.
If not found, crio wipe will clear the storage directory.

## CRIO.API TABLE

The `crio.api` table contains settings for the kubelet/gRPC interface.

**listen**="/var/run/crio/crio.sock"
Path to AF_LOCAL socket on which CRI-O will listen.

**stream_address**="127.0.0.1"
IP address on which the stream server will listen.

**stream_port**="0"
The port on which the stream server will listen. If the port is set to "0", then CRI-O will allocate a random free port number.

**stream_enable_tls**=false
Enable encrypted TLS transport of the stream server.

**stream_idle_timeout**=""
Length of time until open streams terminate due to lack of activity.

**stream_tls_cert**=""
Path to the x509 certificate file used to serve the encrypted stream. This file can change and CRI-O will automatically pick up the changes within 5 minutes.

**stream_tls_key**=""
Path to the key file used to serve the encrypted stream. This file can change and CRI-O will automatically pick up the changes within 5 minutes.

**stream_tls_ca**=""
Path to the x509 CA(s) file used to verify and authenticate client communication with the encrypted stream. This file can change and CRI-O will automatically pick up the changes within 5 minutes.

**grpc_max_send_msg_size**=83886080
Maximum grpc send message size in bytes. If not set or <=0, then CRI-O will default to 80 _ 1024 _ 1024.

**grpc_max_recv_msg_size**=83886080
Maximum grpc receive message size. If not set or <= 0, then CRI-O will default to 80 _ 1024 _ 1024.

## CRIO.RUNTIME TABLE

The `crio.runtime` table contains settings pertaining to the OCI runtime used and options for how to set up and manage the OCI runtime.

**default_runtime**="crun"
The _name_ of the OCI runtime to be used as the default. This option supports live configuration reload.

**default_ulimits**=[]
A list of ulimits to be set in containers by default, specified as "<ulimit name>=<soft limit>:<hard limit>", for example:"nofile=1024:2048". If nothing is set here, settings will be inherited from the CRI-O daemon.

**no_pivot**=false
If true, the runtime will not use `pivot_root`, but instead use `MS_MOVE`.

**decryption_keys_path**="/etc/crio/keys/"
Path where the keys required for image decryption are located

**conmon**=""
Path to the conmon binary, used for monitoring the OCI runtime. Will be searched for using $PATH if empty.
This option is currently deprecated, and will be replaced with RuntimeHandler.MonitorPath.

**conmon_cgroup**=""
Cgroup setting for conmon
This option is currently deprecated, and will be replaced with RuntimeHandler.MonitorCgroup.

**conmon_env**=[]
Environment variable list for the conmon process, used for passing necessary environment variables to conmon or the runtime.
This option is currently deprecated, and will be replaced with RuntimeHandler.MonitorEnv.

**default_env**=[]
Additional environment variables to set for all the containers. These are overridden if set in the container image spec or in
the container runtime configuration.

**selinux**=false
If true, SELinux will be used for pod separation on the host.
This option is deprecated, and be interpreted from whether SELinux is enabled on the host in the future.

**seccomp_profile**=""
Path to the seccomp.json profile which is used as the default seccomp profile for the runtime. If not specified, then the internal default seccomp profile will be used.
This option is currently deprecated, and will be replaced by the SeccompDefault FeatureGate in Kubernetes.

**privileged_seccomp_profile**=""
Enable a seccomp profile for privileged containers from the local path.

**apparmor_profile**=""
Used to change the name of the default AppArmor profile of CRI-O. The default profile name is "crio-default".

**blockio_config_file**=""
Path to the blockio class configuration file for configuring the cgroup blockio controller.

**blockio_reload**=false
If true, the runtime reloads blockio_config_file and rescans block devices in the system before applying blockio parameters.

**cdi_spec_dirs**=[]
Directories to scan for Container Device Interface Specifications to enable CDI device injection. For more details about CDI and the syntax of CDI Spec files please refer to https://github.com/container-orchestrated-devices/container-device-interface.

Directories later in the list have precedence over earlier ones. The default directory list is:

```
  cdi_spec_dirs = [
	  "/etc/cdi",
	  "/var/run/cdi",
  ]
```

**irqbalance_config_file**="/etc/sysconfig/irqbalance"
Used to change irqbalance service config file which is used by CRI-O.
For CentOS/SUSE, this file is located at /etc/sysconfig/irqbalance. For Ubuntu, this file is located at /etc/default/irqbalance.

**irqbalance_config_restore_file**="/etc/sysconfig/orig_irq_banned_cpus"
Used to set the irqbalance banned cpu mask to restore at CRI-O startup. If set to 'disable', no restoration attempt will be done.

**rdt_config_file**=""
Path to the RDT configuration file for configuring the resctrl pseudo-filesystem.

**cgroup_manager**="systemd"
Cgroup management implementation used for the runtime.

**default_capabilities**=[]
List of default capabilities for containers. If it is empty or commented out, only the capabilities defined in the container json file by the user/kube will be added.

The default list is:

```
  default_capabilities = [
	  "CHOWN",
	  "DAC_OVERRIDE",
	  "FSETID",
	  "FOWNER",
	  "SETGID",
	  "SETUID",
	  "SETPCAP",
	  "NET_BIND_SERVICE",
	  "KILL",
  ]
```

**add_inheritable_capabilities**=false
Add capabilities to the inheritable set, as well as the default group of permitted, bounding and effective.
If capabilities are expected to work for non-root users, this option should be set.

**default_sysctls**=[]
List of default sysctls. If it is empty or commented out, only the sysctls defined in the container json file by the user/kube will be added.

One example would be allowing ping inside of containers. On systems that support `/proc/sys/net/ipv4/ping_group_range`, the default list could be:

```
  default_sysctls = [
       "net.ipv4.ping_group_range = 0   2147483647",
  ]
```

**allowed_devices**=[]
List of devices on the host that a user can specify with the "io.kubernetes.cri-o.Devices" allowed annotation.

**additional_devices**=[]
List of additional devices. Specified as "<device-on-host>:<device-on-container>:<permissions>", for example: "--additional-devices=/dev/sdc:/dev/xvdc:rwm". If it is empty or commented out, only the devices defined in the container json file by the user/kube will be added.

**hooks_dir**=["*path*", ...]
Each `*.json` file in the path configures a hook for CRI-O containers. For more details on the syntax of the JSON files and the semantics of hook injection, see `oci-hooks(5)`. CRI-O currently support both the 1.0.0 and 0.1.0 hook schemas, although the 0.1.0 schema is deprecated.

Paths listed later in the array have higher precedence (`oci-hooks(5)` discusses directory precedence).

For the annotation conditions, CRI-O uses the Kubernetes annotations, which are a subset of the annotations passed to the OCI runtime. For example, `io.kubernetes.cri-o.Volumes` is part of the OCI runtime configuration annotations, but it is not part of the Kubernetes annotations being matched for hooks.

For the bind-mount conditions, only mounts explicitly requested by Kubernetes configuration are considered. Bind mounts that CRI-O inserts by default (e.g. `/dev/shm`) are not considered.

**default_mounts**=[]
List of default mounts for each container. **Deprecated:** this option will be removed in future versions in favor of `default_mounts_file`.

**default_mounts_file**=""
Path to the file specifying the defaults mounts for each container. The format of the config is /SRC:/DST, one mount per line. Notice that CRI-O reads its default mounts from the following two files:

    1) `/etc/containers/mounts.conf` (i.e., default_mounts_file): This is the override file, where users can either add in their own default mounts, or override the default mounts shipped with the package.

    2) `/usr/share/containers/mounts.conf`: This is the default file read for mounts. If you want CRI-O to read from a different, specific mounts file, you can change the default_mounts_file. Note, if this is done, CRI-O will only add mounts it finds in this file.

**pids_limit**=-1
Maximum number of processes allowed in a container.
This option is deprecated. The Kubelet flag `--pod-pids-limit` should be used instead.

**log_filter**=""
Filter the log messages by the provided regular expression. This option supports live configuration reload. For example 'request:.\*' filters all gRPC requests.

**log_level**="info"
Changes the verbosity of the logs based on the level it is set to. Options are fatal, panic, error, warn, info, debug, and trace. This option supports live configuration reload.

**log_size_max**=-1
Maximum size allowed for the container log file. Negative numbers indicate that no size limit is imposed. If it is positive, it must be >= 8192 to match/exceed conmon's read buffer. The file is truncated and re-opened so the limit is never exceeded.
This option is deprecated. The Kubelet flag `--container-log-max-size` should be used instead.

**log_to_journald**=false
Whether container output should be logged to journald in addition to the kubernetes log file.

**container_exits_dir**="/var/run/crio/exits"
Path to directory in which container exit files are written to by conmon.

**container_attach_socket_dir**="/var/run/crio"
Path to directory for container attach sockets.

**bind_mount_prefix**=""
A prefix to use for the source of the bind mounts. This option would be useful when running CRI-O in a container and the / directory on the host is mounted as /host in the container. Then if CRI-O runs with the --bind-mount-prefix=/host option, CRI-O would add the /host directory to any bind mounts it hands over to CRI. If Kubernetes asked to have /var/lib/foobar bind mounted into the container, then CRI-O would bind mount /host/var/lib/foobar. Since CRI-O itself is running in a container with / or the host mounted on /host, the container would end up with /var/lib/foobar from the host mounted in the container rather than /var/lib/foobar from the CRI-O container.

**read_only**=false
If set to true, all containers will run in read-only mode.

**uid_mappings**=""
The UID mappings for the user namespace of each container. A range is specified in the form containerUID:HostUID:Size. Multiple ranges must be separated by comma.
This option is deprecated, and will be replaced with native Kubernetes user namespace support in the future.

**minimum_mappable_uid**=-1
The lowest host UID which can be specified in mappings supplied, either as part of a **uid_mappings** or as part of a request received over CRI, for a pod that will be run as a UID other than 0.
This option is deprecated, and will be replaced with Kubernetes user namespace support (KEP-127) in the future.

**gid_mappings**=""
The GID mappings for the user namespace of each container. A range is specified in the form containerGID:HostGID:Size. Multiple ranges must be separated by comma.
This option is deprecated, and will be replaced with Kubernetes user namespace support (KEP-127) in the future.

**minimum_mappable_gid**=-1
The lowest host GID which can be specified in mappings supplied, either as part of a **gid_mappings** or as part of a request received over CRI, for a pod that will be run as a UID other than 0.
This option is deprecated, and will be replaced with Kubernetes user namespace support (KEP-127) in the future.

**ctr_stop_timeout**=30
The minimal amount of time in seconds to wait before issuing a timeout regarding the proper termination of the container.

**drop_infra_ctr**=true
Determines whether we drop the infra container when a pod does not have a private PID namespace, and does not use a kernel separating runtime (like kata).
Requires **manage_ns_lifecycle** to be true.

**infra_ctr_cpuset**=""
Determines the CPU set to run infra containers. If not specified, the CRI-O will use all online CPUs to run infra containers.
You can specify CPUs in the Linux CPU list format.
To get better isolation for guaranteed pods, set this parameter to be equal to kubelet reserved-cpus.

**shared_cpuset**=""
Determines the CPU set which is allowed to be shared between guaranteed containers,
regardless of, and in addition to, the exclusiveness of their CPUs.
This field is optional and would not be used if not specified.
You can specify CPUs in the Linux CPU list format.

**namespaces_dir**="/var/run"
The directory where the state of the managed namespaces gets tracked. Only used when manage_ns_lifecycle is true

**pinns_path**=""
The path to find the pinns binary, which is needed to manage namespace lifecycle

**absent_mount_sources_to_reject**=[]
A list of paths that, when absent from the host, will cause a container creation to fail (as opposed to the current behavior of creating a directory).

**device_ownership_from_security_context**=false
Changes the default behavior of setting container devices uid/gid from CRI's SecurityContext (RunAsUser/RunAsGroup) instead of taking host's uid/gid.

**enable_criu_support**=true
Enable CRIU integration, requires that the criu binary is available in $PATH. (default: true)

**enable_pod_events**=false
Enable CRI-O to generate the container pod-level events in order to optimize the performance of the Pod Lifecycle Event Generator (PLEG) module in Kubelet.

**hostnetwork_disable_selinux**=true
Determines whether SELinux should be disabled within a pod when it is running in the host network namespace.

**disable_hostport_mapping**=false
Enable/Disable the container hostport mapping in CRI-O. Default value is set to 'false'.

**timezone**=""
To set the timezone for a container in CRI-O. If an empty string is provided, CRI-O retains its default behavior. Use 'Local' to match the timezone of the host machine.

### CRIO.RUNTIME.RUNTIMES TABLE

The "crio.runtime.runtimes" table defines a list of OCI compatible runtimes. The runtime to use is picked based on the runtime handler provided by the CRI. If no runtime handler is provided, the runtime will be picked based on the level of trust of the workload. This option supports live configuration reload. This option supports live configuration reload.

**runtime_path**=""
Path to the OCI compatible runtime used for this runtime handler.

**runtime_root**=""
Root directory used to store runtime data

**runtime_type**="oci"
Type of the runtime used for this runtime handler. "oci", "vm"

**inherit_default_runtime**=false
Override the runtime path, runtime config path, runtime root and runtime type from the default runtime on load.

**runtime_config_path**=""
Path to the runtime configuration file, should only be used with VM runtime types

**privileged_without_host_devices**=false
Whether this runtime handler prevents host devices from being passed to privileged containers.

**allowed_annotations**=[]
**This field is currently DEPRECATED. If you'd like to use allowed_annotations, please use a workload.**
A list of experimental annotations this runtime handler is allowed to process.
The currently recognized values are:
"io.kubernetes.cri-o.userns-mode" for configuring a user namespace for the pod.
"io.kubernetes.cri-o.Devices" for configuring devices for the pod.
"io.kubernetes.cri-o.ShmSize" for configuring the size of /dev/shm.
"io.kubernetes.cri-o.UnifiedCgroup.$CTR_NAME" for configuring the cgroup v2 unified block for a container.
"io.containers.trace-syscall" for tracing syscalls via the OCI seccomp BPF hook.
"seccomp-profile.kubernetes.cri-o.io" for setting the seccomp profile for: - a specific container by using: "seccomp-profile.kubernetes.cri-o.io/<CONTAINER_NAME>" - a whole pod by using: "seccomp-profile.kubernetes.cri-o.io/POD"
Note that the annotation works on containers as well as on images.
For images, the plain annotation `seccomp-profile.kubernetes.cri-o.io`
can be used without the required `/POD` suffix or a container name.

**container_min_memory**=""
The minimum memory that must be set for a container. This value can be used to override the currently set global value for a specific runtime. If not set, a global default value of "12 MiB" will be used.

**platform_runtime_paths**={}
A mapping of platforms to the corresponding runtime executable paths for the runtime handler.

**no_sync_log**=false
If set to true, the runtime will not sync the log file on rotate or container exit. This option is only valid for the 'oci'
runtime type. Setting this option to true can cause data loss, e.g. when a machine crash happens.

**default_annotations**={}
A mapping of keys to values of annotations set on containers run by this runtime handler, if not overridden by the pod spec.

**stream_websockets**=false
Enable the WebSocket protocol for container exec, attach and port forward.
conmon-rs (`runtime_type = "pod"`) supports this configuration for exec and attach. Forwarding ports will be supported in future releases.

### CRIO.RUNTIME.WORKLOADS TABLE

The "crio.runtime.workloads" table defines a list of workloads - a way to customize the behavior of a pod and container.
A workload is chosen for a pod based on whether the workload's **activation_annotation** is an annotation on the pod.

**activation_annotation**=""
activation_annotation is the pod annotation that activates these workload settings.

**annotation_prefix**=""
annotation_prefix is the way a pod can override a specific resource for a container.
The full annotation must be of the form `$annotation_prefix.$resource/$ctrname = $value`.

**allowed_annotations**=[]
allowed_annotations is a slice of experimental annotations that this workload is allowed to process.
The currently recognized values are:
"io.kubernetes.cri-o.userns-mode" for configuring a user namespace for the pod.
"io.kubernetes.cri-o.cgroup2-mount-hierarchy-rw" for mounting cgroups writably when set to "true".
"io.kubernetes.cri-o.Devices" for configuring devices for the pod.
"io.kubernetes.cri-o.ShmSize" for configuring the size of /dev/shm.
"io.kubernetes.cri-o.UnifiedCgroup.$CTR_NAME" for configuring the cgroup v2 unified block for a container.
"io.containers.trace-syscall" for tracing syscalls via the OCI seccomp BPF hook.
"io.kubernetes.cri-o.seccompNotifierAction" for enabling the seccomp notifier feature.
"io.kubernetes.cri-o.umask" for setting the umask for container init process.
"io.kubernetes.cri.rdt-class" for setting the RDT class of a container
"seccomp-profile.kubernetes.cri-o.io" for setting the seccomp profile for: - a specific container by using: "seccomp-profile.kubernetes.cri-o.io/<CONTAINER_NAME>" - a whole pod by using: "seccomp-profile.kubernetes.cri-o.io/POD"
Note that the annotation works on containers as well as on images.
"io.kubernetes.cri-o.DisableFIPS" for disabling FIPS mode for a pod within a FIPS-enabled Kubernetes cluster.

#### Using the seccomp notifier feature:

This feature can help you to debug seccomp related issues, for example if
blocked syscalls (permission denied errors) have negative impact on the
workload.

To be able to use this feature, configure a runtime which has the annotation
"io.kubernetes.cri-o.seccompNotifierAction" in the `allowed_annotations` array.

It also requires at least runc 1.1.0 or crun 0.19 which support the notifier
feature.

If everything is setup, CRI-O will modify chosen seccomp profiles for containers
if the annotation "io.kubernetes.cri-o.seccompNotifierAction" is set on the Pod
sandbox. CRI-O will then get notified if a container is using a blocked syscall
and then terminate the workload after a timeout of 5 seconds if the value of
"io.kubernetes.cri-o.seccompNotifierAction=stop".

This also means that multiple syscalls can be captured during that period, while
the timeout will get reset once a new syscall has been discovered.

This also means that the Pods "restartPolicy" has to be set to "Never",
otherwise the kubelet will restart the container immediately.

Please be aware that CRI-O is not able to get notified if a syscall gets blocked
based on the seccomp defaultAction, which is a general runtime limitation.

### CRIO.RUNTIME.WORKLOAD.RESOURCES TABLE

The resources table is a structure for overriding certain resources for pods using this workload.
This structure provides a default value, and can be overridden by using the AnnotationPrefix.

**cpushares**=""
Specifies the number of CPU shares this pod has access to.

**cpuset**=""
Specifies the cpuset this pod has access to.

## CRIO.IMAGE TABLE

The `crio.image` table contains settings pertaining to the management of OCI images.

CRI-O reads its configured registries defaults from the system wide containers-registries.conf(5) located in /etc/containers/registries.conf.

**default_transport**="docker://"
Default transport for pulling images from a remote container storage.

**global_auth_file**=""
The path to a file like /var/lib/kubelet/config.json holding credentials necessary for pulling images from secure registries.

**pause_image**="registry.k8s.io/pause:3.10.1"
The on-registry image used to instantiate infra containers.
The value should start with a registry host name.
This option supports live configuration reload.

**pause_image_auth_file**=""
The path to a file like /var/lib/kubelet/config.json holding credentials specific to pulling the pause_image from above. This option supports live configuration reload.

**pause_command**="/pause"
The command to run to have a container stay in the paused state. This option supports live configuration reload.

**pinned_images**=[]
A list of images to be excluded from the kubelet's garbage collection. It allows specifying image names using either exact, glob, or keyword patterns. Exact matches must match the entire name, glob matches can have a wildcard \* at the end, and keyword matches can have wildcards on both ends. By default, this list includes the `pause` image if configured by the user, which is used as a placeholder in Kubernetes pods.

**signature_policy**=""
Path to the file which decides what sort of policy we use when deciding whether or not to trust an image that we've pulled. It is not recommended that this option be used, as the default behavior of using the system-wide default policy (i.e., /etc/containers/policy.json) is most often preferred. Please refer to containers-policy.json(5) for more details.

**signature_policy_dir**="/etc/crio/policies"
Root path for pod namespace-separated signature policies. The final policy to be used on image pull will be <SIGNATURE_POLICY_DIR>/\<NAMESPACE\>.json. If no pod namespace is being provided on image pull (via the sandbox config), or the concatenated path is non existent, then the signature_policy or system wide policy will be used as fallback. Must be an absolute path.

**image_volumes**="mkdir"
Controls how image volumes are handled. The valid values are mkdir, bind and ignore; the latter will ignore volumes entirely.

**insecure_registries**=[]
List of registries to skip TLS verification for pulling images.
This option is deprecated. Use registries.conf instead.

**big_files_temporary_dir**=""
Path to the temporary directory to use for storing big files, used to store image blobs and data streams related to containers image management.

**separate_pull_cgroup**=""
[EXPERIMENTAL] If its value is set, then images are pulled into the specified cgroup. If its value is set to "pod", then the pod's cgroup is used. It is currently supported only with the systemd cgroup manager.

**auto_reload_registries**=false
If true, CRI-O will automatically reload the mirror registry when there is an update to the 'registries.conf.d' directory. Default value is set to 'false'.

**pull_progress_timeout**="0s"
The timeout for an image pull to make progress until the pull operation gets canceled. This value will be also used for calculating the pull progress interval to pull_progress_timeout / 10. Can be set to 0 to disable the timeout as well as the progress output.

**oci_artifact_mount_support**=true
This option is whether CRI-O enables OCI Artifact mount.
If true, CRI-O can mount OCI artifacts as volumes.

## CRIO.NETWORK TABLE

The `crio.network` table containers settings pertaining to the management of CNI plugins.

**cni_default_network**=""
The default CNI network name to be selected. If not set or "", then CRI-O will pick-up the first one found in network_dir.

**network_dir**="/etc/cni/net.d/"
Path to the directory where CNI configuration files are located.

**plugin_dirs**=["/opt/cni/bin/",]
List of paths to directories where CNI plugin binaries are located.

## CRIO.METRICS TABLE

The `crio.metrics` table containers settings pertaining to the Prometheus based metrics retrieval.

**enable_metrics**=false
Globally enable or disable metrics support.

**metrics_collectors**=["image_pulls_layer_size", "containers_events_dropped_total", "containers_oom_total", "processes_defunct", "operations_total", "operations_latency_seconds", "operations_latency_seconds_total", "operations_errors_total", "image_pulls_bytes_total", "image_pulls_skipped_bytes_total", "image_pulls_failure_total", "image_pulls_success_total", "image_layer_reuse_total", "containers_oom_count_total", "containers_seccomp_notifier_count_total", "resources_stalled_at_stage"]
Specify enabled metrics collectors. Per default all metrics are enabled.

**metrics_host**="127.0.0.1"
The IP address or hostname on which the metrics server will listen.

**metrics_port**=9090
The port on which the metrics server will listen.

**metrics_socket**=""
The socket on which the metrics server will listen.

**metrics_cert**=""
The certificate for the secure metrics server.

**metrics_key**=""
The certificate key for the secure metrics server.

## CRIO.TRACING TABLE

[EXPERIMENTAL] The `crio.tracing` table containers settings pertaining to the export of OpenTelemetry trace data.

**enable_tracing**=false
Globally enable or disable OpenTelemetry trace data exporting.

**tracing_endpoint**="127.0.0.1:4317"
Address on which the gRPC trace collector will listen.

**tracing_sampling_rate_per_million**=""
Number of samples to collect per million OpenTelemetry spans. Set to 1000000 to always sample.

## CRIO.STATS TABLE

The `crio.stats` table specifies all necessary configuration for reporting container and pod stats.

**stats_collection_period**=0
The number of seconds between collecting pod and container stats. If set to 0, the stats are collected on-demand instead. **DEPRECATED:** This option will be removed in the future. Please use `collection_period` instead.

**collection_period**=0
The number of seconds between collecting pod/container stats and pod sandbox metrics. If set to 0, the metrics/stats are collected on-demand instead.

**included_pod_metrics**=[]
A list of pod metrics to include. Specify the names of the metrics to include in this list.

## CRIO.NRI TABLE

The `crio.nri` table contains settings for controlling NRI (Node Resource Interface) support in CRI-O.
**enable_nri**=true
Enable CRI-O NRI support.

**nri_plugin_dir**="/opt/nri/plugins"
Directory to scan for pre-installed plugins to automatically start.

**nri_plugin_config_dir**="/etc/nri/conf.d"
Directory to scan for configuration of pre-installed plugins.

**nri_listen**="/var/run/nri/nri.sock"
Socket to listen on for externally started NRI plugins to connect to.

**nri_disable_connections**=false
Disable connections from externally started NRI plugins.

**nri_plugin_registration_timeout**="5s"
Timeout for a plugin to register itself with NRI.

**nri_plugin_request_timeout**="2s"
Timeout for a plugin to handle an NRI request.

# SEE ALSO

crio.conf.d(5), containers-storage.conf(5), containers-policy.json(5), containers-registries.conf(5), crio(8)

# HISTORY

Aug 2018, Update to the latest state by Valentin Rothberg <vrothberg@suse.com>

Oct 2016, Originally compiled by Aleksa Sarai <asarai@suse.de>

[toml]: https://github.com/toml-lang/toml
[crio]: ./crio.8.md
